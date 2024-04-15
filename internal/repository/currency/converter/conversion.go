package converter

import (
	"calc/internal/model"
	repoModel "calc/internal/repository/currency/model"
)

// ToCurrencyFromRepo converts repoModel.Currency to model.Currency
func ToCurrencyFromRepo(pair *repoModel.Currency) *model.Currency {
	return &model.Currency{
		Symbol: pair.Symbol,
	}
}
