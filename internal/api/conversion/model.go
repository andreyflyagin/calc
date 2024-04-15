package conversion

import "github.com/shopspring/decimal"

// Conversion model
type Conversion struct {
	Base   string          `json:"base"`
	Amount decimal.Decimal `json:"amount"`
	Quote  string          `json:"quote"`
	Value  decimal.Decimal `json:"value"`
	Rate   decimal.Decimal `json:"rate"`
}
