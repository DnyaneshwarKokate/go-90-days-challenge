package mocks

import (
	"context"

	"day-42/domain"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(ctx context.Context, order *domain.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockOrderRepository) FindByID(ctx context.Context, id string) (*domain.Order, error) {
	args := m.Called(ctx, id)
	if order, ok := args.Get(0).(*domain.Order); ok {
		return order, args.Error(1)
	}
	return nil, args.Error(1)
}
