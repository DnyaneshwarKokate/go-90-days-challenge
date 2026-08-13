package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/dnyaneshwarkokate/go-90-days-challenge/Day-48/domain"
)

// Store defines thread-safe storage with disk persistence capabilities for PVCs.
type Store struct {
	mu       sync.RWMutex
	filePath string
	orders   map[string]domain.Order
}

// NewStore initializes a Store and loads existing data from disk if present.
func NewStore(filePath string) (*Store, error) {
	s := &Store{
		filePath: filePath,
		orders:   make(map[string]domain.Order),
	}

	// Load existing orders from disk volume if present
	if err := s.loadFromDisk(); err != nil {
		return nil, fmt.Errorf("failed to load orders from storage [%s]: %w", filePath, err)
	}

	return s, nil
}

// Save inserts an order into memory and syncs to disk storage.
func (s *Store) Save(order domain.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[order.ID] = order
	return s.syncToDisk()
}

// GetAll returns a list of all stored orders.
func (s *Store) GetAll() []domain.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]domain.Order, 0, len(s.orders))
	for _, o := range s.orders {
		list = append(list, o)
	}
	return list
}

// GetByID returns an order by ID.
func (s *Store) GetByID(id string) (domain.Order, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, exists := s.orders[id]
	return order, exists
}

// Count returns the total number of orders.
func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.orders)
}

func (s *Store) loadFromDisk() error {
	if s.filePath == "" {
		return nil
	}

	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		// File does not exist yet; ensure directory exists
		dir := filepath.Dir(s.filePath)
		if dir != "." && dir != "/" {
			_ = os.MkdirAll(dir, 0755)
		}
		return nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	if len(data) == 0 {
		return nil
	}

	var list []domain.Order
	if err := json.Unmarshal(data, &list); err != nil {
		return err
	}

	for _, o := range list {
		s.orders[o.ID] = o
	}
	return nil
}

func (s *Store) syncToDisk() error {
	if s.filePath == "" {
		return nil
	}

	list := make([]domain.Order, 0, len(s.orders))
	for _, o := range s.orders {
		list = append(list, o)
	}

	data, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}

	dir := filepath.Dir(s.filePath)
	if dir != "." && dir != "/" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Write atomically or write to file
	tmpFile := fmt.Sprintf("%s.tmp-%d", s.filePath, time.Now().UnixNano())
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		return err
	}
	return os.Rename(tmpFile, s.filePath)
}
