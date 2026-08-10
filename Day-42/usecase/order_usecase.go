package usecase

import (
	"context"
	"fmt"
	"time"

	"day-42/calculator"
	"day-42/domain"

	"github.com/google/uuid"
)

type OrderUseCase struct {
	repo domain.OrderRepository
}

func NewOrderUseCase(repo domain.OrderRepository) *OrderUseCase {
	return &OrderUseCase{repo: repo}
}

func (u *OrderUseCase) PlaceOrder(ctx context.Context, customer string, subtotal float64, promoCode string) (*domain.Order, error) {
	if subtotal <= 0 {
		return nil, domain.ErrInvalidOrderAmount
	}

	discount, err := calculator.CalculateDiscount(subtotal, promoCode)
	if err != nil {
		return nil, err
	}

	taxableAmount := subtotal - discount
	tax := calculator.CalculateTax(taxableAmount, 0.18)
	total := taxableAmount + tax

	order := &domain.Order{
		ID:        fmt.Sprintf("ord_%s", uuid.New().String()[:8]),
		Customer:  customer,
		Subtotal:  subtotal,
		Discount:  discount,
		Tax:       tax,
		Total:     total,
		Status:    "CONFIRMED",
		CreatedAt: time.Now(),
	}

	if err := u.repo.Save(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, nil
}
