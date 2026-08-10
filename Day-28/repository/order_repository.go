package repository

import (
	"sync"

	"day-28/domain"
)

// memoryOrderRepository provides a thread-safe in-memory store for orders.
type memoryOrderRepository struct {
	mu     sync.RWMutex
	orders map[string]*domain.Order
}

// NewMemoryOrderRepository initializes and returns a domain.OrderRepository interface implementation.
func NewMemoryOrderRepository() domain.OrderRepository {
	return &memoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (r *memoryOrderRepository) Save(order *domain.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.orders[order.ID] = order
	return nil
}

func (r *memoryOrderRepository) FindByID(id string) (*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, exists := r.orders[id]
	if !exists {
		return nil, domain.ErrOrderNotFound
	}
	return order, nil
}

func (r *memoryOrderRepository) FindAll() ([]*domain.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]*domain.Order, 0, len(r.orders))
	for _, order := range r.orders {
		result = append(result, order)
	}
	return result, nil
}
