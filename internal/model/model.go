package model

import "github.com/shopspring/decimal"

// Conversion represents a conversion between two currencies
type Conversion struct {
	Base   string
	Amount decimal.Decimal
	Quote  string
	Value  decimal.Decimal
	Rate   decimal.Decimal
}

// Currency represents a currency
type Currency struct {
	Symbol string
}

// Pair represents a pair of currencies
type Pair struct {
	Base  string
	Quote string
}

// PairRates represents a pair of currencies with a rate
type PairRates struct {
	Pair
	Rate decimal.Decimal
}
