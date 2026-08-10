package usecase_test

import (
	"testing"

	"day-28/domain"
	"day-28/logger"
	"day-28/notification"
	"day-28/repository"
	"day-28/usecase"
)

func TestCreateOrder_Success(t *testing.T) {
	// 1. Instantiating injected mock dependencies
	mockRepo := repository.NewMemoryOrderRepository()
	mockNotifier := notification.NewMockNotificationService()
	mockLogger := logger.NewMockLogger()

	// 2. Inject dependencies into UseCase via Constructor Injection
	orderUC := usecase.NewOrderUseCase(mockRepo, mockNotifier, mockLogger)

	// 3. Execute method
	input := domain.CreateOrderInput{
		CustomerEmail: "alice@example.com",
		Amount:        150.75,
	}

	order, err := orderUC.CreateOrder(input)

	// 4. Assertions
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if order.ID == "" {
		t.Errorf("expected generated Order ID, got empty string")
	}

	if order.CustomerEmail != input.CustomerEmail {
		t.Errorf("expected email %s, got %s", input.CustomerEmail, order.CustomerEmail)
	}

	// Verify notification dependency behavior
	if len(mockNotifier.SentOrders) != 1 {
		t.Errorf("expected 1 sent notification, got %d", len(mockNotifier.SentOrders))
	}

	// Verify logger dependency behavior
	if len(mockLogger.Logs) == 0 {
		t.Errorf("expected logged entries, got 0")
	}
}

func TestCreateOrder_InvalidAmount(t *testing.T) {
	mockRepo := repository.NewMemoryOrderRepository()
	mockNotifier := notification.NewMockNotificationService()
	mockLogger := logger.NewMockLogger()

	orderUC := usecase.NewOrderUseCase(mockRepo, mockNotifier, mockLogger)

	input := domain.CreateOrderInput{
		CustomerEmail: "bob@example.com",
		Amount:        -50.00,
	}

	order, err := orderUC.CreateOrder(input)

	if err != domain.ErrInvalidOrderAmount {
		t.Fatalf("expected ErrInvalidOrderAmount, got %v", err)
	}

	if order != nil {
		t.Errorf("expected nil order on error, got %+v", order)
	}

	// Verify notification service was NOT invoked when validation fails
	if len(mockNotifier.SentOrders) != 0 {
		t.Errorf("expected 0 sent notifications on failure, got %d", len(mockNotifier.SentOrders))
	}
}
