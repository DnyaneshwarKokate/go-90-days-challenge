package calculator

import (
	"errors"
	"fmt"
	"math"
)

var ErrInvalidDiscountCode = errors.New("invalid or expired discount code")

func CalculateDiscount(subtotal float64, promoCode string) (float64, error) {
	if subtotal <= 0 {
		return 0, fmt.Errorf("subtotal must be greater than zero")
	}

	switch promoCode {
	case "SAVE10":
		return math.Round(subtotal*0.10*100) / 100, nil
	case "SAVE20":
		return math.Round(subtotal*0.20*100) / 100, nil
	case "VIP50":
		return math.Round(subtotal*0.50*100) / 100, nil
	case "":
		return 0, nil
	default:
		return 0, ErrInvalidDiscountCode
	}
}

func CalculateTax(amount float64, taxRate float64) float64 {
	if amount <= 0 || taxRate <= 0 {
		return 0
	}
	return math.Round(amount*taxRate*100) / 100
}
