package config

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// ConfigChangeEvent notifies watchers of real-time configuration updates.
type ConfigChangeEvent struct {
	Key       string
	OldValue  string
	NewValue  string
	Timestamp time.Time
}

// DynamicConfigServer manages centralized application configuration and watch listeners.
type DynamicConfigServer struct {
	mu        sync.RWMutex
	store     map[string]string
	watchers  map[string][]chan ConfigChangeEvent
}

// NewDynamicConfigServer initializes config server.
func NewDynamicConfigServer() *DynamicConfigServer {
	return &DynamicConfigServer{
		store:    make(map[string]string),
		watchers: make(map[string][]chan ConfigChangeEvent),
	}
}

// Get returns the value of a configuration key.
func (s *DynamicConfigServer) Get(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, exists := s.store[key]
	if !exists {
		return "", fmt.Errorf("config key %s not found", key)
	}
	return val, nil
}

// Set updates or inserts a key-value pair and notifies active watchers.
func (s *DynamicConfigServer) Set(key string, value string) {
	s.mu.Lock()
	oldVal := s.store[key]
	s.store[key] = value

	event := ConfigChangeEvent{
		Key:       key,
		OldValue:  oldVal,
		NewValue:  value,
		Timestamp: time.Now(),
	}

	listeners := s.watchers[key]
	s.mu.Unlock()

	// Notify all active watchers
	for _, ch := range listeners {
		select {
		case ch <- event:
		default:
			// Non-blocking write if watcher channel buffer full
		}
	}
}

// Watch subscribes to real-time change events for a specific config key.
func (s *DynamicConfigServer) Watch(key string) (<-chan ConfigChangeEvent, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if key == "" {
		return nil, errors.New("key cannot be empty")
	}

	ch := make(chan ConfigChangeEvent, 5)
	s.watchers[key] = append(s.watchers[key], ch)
	return ch, nil
}
