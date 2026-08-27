package queue

import (
	"errors"
	"sync"
	"time"
)

var ErrQueueClosed = errors.New("message queue is closed")

type Message struct {
	ID        string
	Topic     string
	Payload   string
	Timestamp time.Time
}

type MessageBroker struct {
	mu       sync.RWMutex
	topics   map[string]chan Message
	dlq      chan Message
	isClosed bool
}

func NewMessageBroker(bufferSize int) *MessageBroker {
	return &MessageBroker{
		topics: make(map[string]chan Message),
		dlq:    make(chan Message, bufferSize),
	}
}

func (mb *MessageBroker) Publish(topic string, msg Message) error {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	if mb.isClosed {
		return ErrQueueClosed
	}

	ch, ok := mb.topics[topic]
	if !ok {
		ch = make(chan Message, 100)
		mb.topics[topic] = ch
	}

	select {
	case ch <- msg:
		return nil
	default:
		// Queue full, route to Dead Letter Queue (DLQ)
		select {
		case mb.dlq <- msg:
			return nil
		default:
			return errors.New("topic and DLQ buffers full")
		}
	}
}

func (mb *MessageBroker) Subscribe(topic string, handler func(Message) bool) {
	mb.mu.RLock()
	ch, ok := mb.topics[topic]
	mb.mu.RUnlock()

	if !ok {
		mb.mu.Lock()
		ch = make(chan Message, 100)
		mb.topics[topic] = ch
		mb.mu.Unlock()
	}

	go func() {
		for msg := range ch {
			ack := handler(msg)
			if !ack {
				// Failed message sent to DLQ
				mb.dlq <- msg
			}
		}
	}()
}

func (mb *MessageBroker) ReadDLQ() []Message {
	var messages []Message
	for {
		select {
		case msg := <-mb.dlq:
			messages = append(messages, msg)
		default:
			return messages
		}
	}
}
