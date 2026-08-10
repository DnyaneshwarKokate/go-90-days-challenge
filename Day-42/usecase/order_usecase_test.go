package usecase_test

import (
	"context"
	"testing"

	"day-42/domain"
	"day-42/mocks"
	"day-42/usecase"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPlaceOrder_SuccessWithMock(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)

	// Setup mock expectation: Save should be called and return nil
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*domain.Order")).Return(nil)

	orderUC := usecase.NewOrderUseCase(mockRepo)
	ctx := context.Background()

	order, err := orderUC.PlaceOrder(ctx, "Dnyaneshwar", 200.0, "SAVE10")

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "Dnyaneshwar", order.Customer)
	assert.Equal(t, 20.0, order.Discount)
	assert.Equal(t, 32.4, order.Tax)       // (200 - 20) * 0.18 = 32.4
	assert.Equal(t, 212.4, order.Total)    // 180 + 32.4 = 212.4

	mockRepo.AssertExpectations(t)
}

func TestPlaceOrder_InvalidAmount(t *testing.T) {
	mockRepo := new(mocks.MockOrderRepository)
	orderUC := usecase.NewOrderUseCase(mockRepo)

	order, err := orderUC.PlaceOrder(context.Background(), "Alice", 0.0, "SAVE10")

	assert.Error(t, err)
	assert.Nil(t, order)
	assert.Equal(t, domain.ErrInvalidOrderAmount, err)

	mockRepo.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
}
