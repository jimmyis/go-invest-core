package main

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		order   Order
		wantErr error
	}{
		{"Expect Valid Input", Order{AmountMinor: 100, Currency: "THB"}, nil},
		{"Expect Invalid Amount (0)", Order{AmountMinor: 0, Currency: "THB"}, ErrInvalidAmount},
		{"Expect Invalid Amount (-1)", Order{AmountMinor: -1, Currency: "THB"}, ErrInvalidAmount},
		{"Expect Unsupported Currency (USD)", Order{AmountMinor: 100, Currency: "USD"}, ErrUnsupportedCurrency},
		{"Expect Invalid Currency (blank)", Order{AmountMinor: 100, Currency: ""}, ErrUnsupportedCurrency},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.order.Validate()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
