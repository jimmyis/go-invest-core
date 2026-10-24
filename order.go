package main

import (
	"errors"
	"fmt"
)

type Order struct {
	AmountMinor int64  `json:"amount_minor"`
	Currency    string `json:"currency"`
}

var (
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrUnsupportedCurrency = errors.New("unsupported currency")
)

func (o Order) Validate() error {

	if o.AmountMinor <= 0 {
		return fmt.Errorf("amount %d: %w", o.AmountMinor, ErrInvalidAmount)
	}

	if o.Currency != "THB" {
		return fmt.Errorf("currency = %s: %w", o.Currency, ErrUnsupportedCurrency)
	}

	return nil
}
