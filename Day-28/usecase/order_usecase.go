package usecase

import (
	"fmt"
	"time"

	"day-28/domain"
)

// OrderUseCase handles core order business rules.
type OrderUseCase struct {
	repo     domain.OrderRepository
	notifier domain.NotificationService
	logger   domain.Logger
}

// Option defines a functional option for configuring OrderUseCase dependencies dynamically.
type Option func(*OrderUseCase)

// WithLogger allows injecting a custom logger via functional options.
func WithLogger(logger domain.Logger) Option {
	return func(u *OrderUseCase) {
		if logger != nil {
			u.logger = logger
		}
	}
}

// WithNotifier allows injecting a custom notification service via functional options.
func WithNotifier(notifier domain.NotificationService) Option {
	return func(u *OrderUseCase) {
		if notifier != nil {
			u.notifier = notifier
		}
	}
}

// NewOrderUseCase constructs OrderUseCase using explicit Constructor Injection with optional Functional Options.
func NewOrderUseCase(
	repo domain.OrderRepository,
	notifier domain.NotificationService,
	logger domain.Logger,
	opts ...Option,
) *OrderUseCase {
	uc := &OrderUseCase{
		repo:     repo,
		notifier: notifier,
		logger:   logger,
	}

	for _, opt := range opts {
		opt(uc)
	}

	return uc
}

// CreateOrder validates input, constructs an order entity, persists it, and sends notification.
func (u *OrderUseCase) CreateOrder(input domain.CreateOrderInput) (*domain.Order, error) {
	if input.Amount <= 0 {
		u.logger.Error("Validation failed: order amount must be positive (received %.2f)", input.Amount)
		return nil, domain.ErrInvalidOrderAmount
	}
	if input.CustomerEmail == "" {
		u.logger.Error("Validation failed: customer email is empty")
		return nil, domain.ErrEmptyCustomerEmail
	}

	order := &domain.Order{
		ID:            fmt.Sprintf("ORD-%d", time.Now().UnixNano()),
		CustomerEmail: input.CustomerEmail,
		Amount:        input.Amount,
		Status:        "COMPLETED",
		CreatedAt:     time.Now(),
	}

	if err := u.repo.Save(order); err != nil {
		u.logger.Error("Failed to save order %s: %v", order.ID, err)
		return nil, fmt.Errorf("repository save error: %w", err)
	}

	u.logger.Info("Order %s saved successfully", order.ID)

	if u.notifier != nil {
		if err := u.notifier.SendOrderConfirmation(order); err != nil {
			u.logger.Error("Failed to send order confirmation for %s: %v", order.ID, err)
			// Business decision: non-fatal, order is already created
		}
	}

	return order, nil
}

// GetOrder retrieves an order by its ID.
func (u *OrderUseCase) GetOrder(id string) (*domain.Order, error) {
	order, err := u.repo.FindByID(id)
	if err != nil {
		u.logger.Error("Failed to find order %s: %v", id, err)
		return nil, err
	}
	return order, nil
}

// ListOrders retrieves all orders.
func (u *OrderUseCase) ListOrders() ([]*domain.Order, error) {
	return u.repo.FindAll()
}
