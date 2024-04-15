package service

import (
	"calc/internal/model"
	"context"
	"errors"
	"github.com/shopspring/decimal"
)

var (
	// ErrorConversionNotFound is the error returned when a conversion is not found
	ErrorConversionNotFound = errors.New("conversion not found")
)

// ConversionService is the interface for the conversion service
type ConversionService interface {
	Run(ctx context.Context)
	GetConversion(ctx context.Context, from, to string, amount decimal.Decimal) (*model.Conversion, error)
}

// ExchangeRatesProviderService is the interface for the exchange rates provider service
type ExchangeRatesProviderService interface {
	GetData(ctx context.Context, pairs []model.Currency) ([]model.PairRates, error)
}
