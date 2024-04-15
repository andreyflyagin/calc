package currency

import (
	"calc/internal/helper"
	"calc/internal/model"
	"calc/internal/repository/currency/converter"
	repoModel "calc/internal/repository/currency/model"
	"context"
	"github.com/jmoiron/sqlx"
)
import def "calc/internal/repository"

var _ def.CurrencyRepository = (*repository)(nil)

type repository struct {
	db *sqlx.DB
}

// NewRepository creates a new instance of currency repository
func NewRepository(db *sqlx.DB) def.CurrencyRepository {
	return &repository{db: db}
}

// GetCurrencies returns all active currencies
func (r repository) GetCurrencies(ctx context.Context) ([]model.Currency, error) {

	rows, err := r.db.Query("SELECT symbol FROM currency WHERE is_active = true")
	if err != nil {

		return nil, err
	}
	defer helper.SafeClose(ctx, rows, "GetCurrencies")

	var result []model.Currency

	for rows.Next() {
		var conversion repoModel.Currency
		if err := rows.Scan(&conversion.Symbol); err != nil {
			return nil, err
		}
		result = append(result, *converter.ToCurrencyFromRepo(&conversion))
	}
	return result, nil
}
