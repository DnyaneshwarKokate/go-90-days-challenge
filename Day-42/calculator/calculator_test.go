package calculator_test

import (
	"testing"

	"day-42/calculator"

	"github.com/stretchr/testify/assert"
)

func assertCloseFloat(t *testing.T, expected, actual float64) {
	t.Helper()
	assert.InDelta(t, expected, actual, 0.01)
}

func TestCalculateDiscount_TableDriven(t *testing.T) {
	tests := []struct {
		name          string
		subtotal      float64
		promoCode     string
		expectedDisc  float64
		expectedErr   bool
		errType       error
	}{
		{
			name:         "Valid 10% Discount",
			subtotal:     100.0,
			promoCode:    "SAVE10",
			expectedDisc: 10.0,
			expectedErr:  false,
		},
		{
			name:         "Valid 20% Discount",
			subtotal:     250.0,
			promoCode:    "SAVE20",
			expectedDisc: 50.0,
			expectedErr:  false,
		},
		{
			name:         "Empty Promo Code",
			subtotal:     150.0,
			promoCode:    "",
			expectedDisc: 0.0,
			expectedErr:  false,
		},
		{
			name:         "Invalid Promo Code",
			subtotal:     100.0,
			promoCode:    "INVALID_CODE",
			expectedDisc: 0.0,
			expectedErr:  true,
			errType:      calculator.ErrInvalidDiscountCode,
		},
		{
			name:         "Zero Subtotal",
			subtotal:     0.0,
			promoCode:    "SAVE10",
			expectedDisc: 0.0,
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			discount, err := calculator.CalculateDiscount(tt.subtotal, tt.promoCode)
			if tt.expectedErr {
				assert.Error(t, err)
				if tt.errType != nil {
					assert.Equal(t, tt.errType, err)
				}
			} else {
				assert.NoError(t, err)
				assertCloseFloat(t, tt.expectedDisc, discount)
			}
		})
	}
}

func TestCalculateTax(t *testing.T) {
	tax := calculator.CalculateTax(100.0, 0.18)
	assert.Equal(t, 18.0, tax)
}
