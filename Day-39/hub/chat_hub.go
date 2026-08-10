package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"day-39/domain"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	ID       string
	RoomID   string
	Username string
	Conn     *websocket.Conn
	Send     chan domain.ChatEvent
	Hub      *ChatHub
}

type ChatHub struct {
	rooms      map[string]map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan domain.ChatEvent
	mu         sync.RWMutex
	logger     domain.Logger
}

func NewChatHub(logger domain.Logger) *ChatHub {
	h := &ChatHub{
		rooms:      make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan domain.ChatEvent, 256),
		logger:     logger,
	}

	go h.run()
	return h
}

func (h *ChatHub) run() {
	ctx := context.Background()
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if _, ok := h.rooms[client.RoomID]; !ok {
				h.rooms[client.RoomID] = make(map[*Client]bool)
			}
			h.rooms[client.RoomID][client] = true
			h.mu.Unlock()

			h.logger.Info(ctx, "Client joined chat room", "username", client.Username, "room", client.RoomID)
			h.broadcastToRoom(client.RoomID, domain.ChatEvent{
				Type:      "JOIN",
				RoomID:    client.RoomID,
				Sender:    "SYSTEM",
				Content:   fmt.Sprintf("%s joined the room", client.Username),
				Timestamp: time.Now(),
			})

		case client := <-h.unregister:
			h.mu.Lock()
			if clients, ok := h.rooms[client.RoomID]; ok {
				if _, exists := clients[client]; exists {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.rooms, client.RoomID)
					}
				}
			}
			h.mu.Unlock()

			h.logger.Info(ctx, "Client left chat room", "username", client.Username, "room", client.RoomID)
			h.broadcastToRoom(client.RoomID, domain.ChatEvent{
				Type:      "LEAVE",
				RoomID:    client.RoomID,
				Sender:    "SYSTEM",
				Content:   fmt.Sprintf("%s left the room", client.Username),
				Timestamp: time.Now(),
			})

		case event := <-h.broadcast:
			h.broadcastToRoom(event.RoomID, event)
		}
	}
}

func (h *ChatHub) broadcastToRoom(roomID string, event domain.ChatEvent) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if clients, ok := h.rooms[roomID]; ok {
		for client := range clients {
			select {
			case client.Send <- event:
			default:
				close(client.Send)
				delete(clients, client)
			}
		}
	}
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()

	for {
		_, messageBytes, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		var event domain.ChatEvent
		if err := json.Unmarshal(messageBytes, &event); err != nil {
			event = domain.ChatEvent{
				Type:      "CHAT",
				RoomID:    c.RoomID,
				Sender:    c.Username,
				Content:   string(messageBytes),
				Timestamp: time.Now(),
			}
		} else {
			event.RoomID = c.RoomID
			event.Sender = c.Username
			event.Timestamp = time.Now()
		}

		c.Hub.broadcast <- event
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for event := range c.Send {
		c.Conn.WriteJSON(event)
	}
}

func (h *ChatHub) HandleWS(w http.ResponseWriter, r *http.Request, roomID, username string) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error(r.Context(), "Chat Hub upgrade error", "error", err)
		return
	}

	client := &Client{
		ID:       uuid.New().String()[:8],
		RoomID:   roomID,
		Username: username,
		Conn:     conn,
		Send:     make(chan domain.ChatEvent, 256),
		Hub:      h,
	}

	client.Hub.register <- client

	go client.WritePump()
	go client.ReadPump()
}
