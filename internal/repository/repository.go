package repository

import (
	"calc/internal/model"
	"context"
)

// CurrencyRepository is the interface that wraps the basic GetConversions method.
type CurrencyRepository interface {
	GetCurrencies(ctx context.Context) ([]model.Currency, error)
}
