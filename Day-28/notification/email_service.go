package notification

import (
	"fmt"
	"sync"

	"day-28/domain"
)

// EmailNotificationService simulates real email delivery.
type EmailNotificationService struct {
	logger domain.Logger
}

// NewEmailNotificationService constructs an EmailNotificationService using constructor injection.
func NewEmailNotificationService(logger domain.Logger) domain.NotificationService {
	return &EmailNotificationService{logger: logger}
}

func (s *EmailNotificationService) SendOrderConfirmation(order *domain.Order) error {
	s.logger.Info("📧 Sending email confirmation for Order ID %s to %s ($%.2f)",
		order.ID, order.CustomerEmail, order.Amount)
	return nil
}

// MockNotificationService captures sent notifications in-memory for testing.
type MockNotificationService struct {
	mu           sync.Mutex
	SentOrders   []*domain.Order
	ShouldFail   bool
}

// NewMockNotificationService creates a configurable mock notification service.
func NewMockNotificationService() *MockNotificationService {
	return &MockNotificationService{
		SentOrders: make([]*domain.Order, 0),
	}
}

func (m *MockNotificationService) SendOrderConfirmation(order *domain.Order) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.ShouldFail {
		return fmt.Errorf("simulated notification failure")
	}
	m.SentOrders = append(m.SentOrders, order)
	return nil
}
