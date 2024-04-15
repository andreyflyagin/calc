package conversion

import (
	"calc/internal/model"
	"calc/internal/repository"
	"context"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"time"
)
import def "calc/internal/service"

var _ def.ConversionService = (*service)(nil)

type service struct {
	currencyRepository           repository.CurrencyRepository
	exchangeRatesProviderService def.ExchangeRatesProviderService
	m                            *map[model.Pair]model.PairRates
}

// NewService creates a new conversion service
func NewService(currencyRepository repository.CurrencyRepository, exchangeRatesProviderService def.ExchangeRatesProviderService) def.ConversionService {
	m := make(map[model.Pair]model.PairRates)
	return &service{
		currencyRepository:           currencyRepository,
		exchangeRatesProviderService: exchangeRatesProviderService,
		m:                            &m,
	}
}

// Run starts the service
func (s *service) Run(ctx context.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		err := s.updateRates(ctx)
		if err != nil {
			log.Error().Err(err).Msg("error updating rates")
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func (s *service) updateRates(ctx context.Context) error {
	currencies, err := s.currencyRepository.GetCurrencies(ctx)
	if err != nil {
		return err
	}

	pairRates, err := s.exchangeRatesProviderService.GetData(ctx, currencies)
	if err != nil {
		return err
	}

	m := make(map[model.Pair]model.PairRates, len(pairRates))

	for _, v := range pairRates {
		m[v.Pair] = v
	}

	s.m = &m
	return nil
}

// GetConversion returns the conversion between two currencies
func (s *service) GetConversion(_ context.Context, from, to string, amount decimal.Decimal) (*model.Conversion, error) {
	pair, ok := (*s.m)[model.Pair{
		Base:  from,
		Quote: to,
	}]

	if !ok {
		return nil, def.ErrorConversionNotFound
	}
	value := amount.Mul(pair.Rate)

	ret := model.Conversion{
		Base:   from,
		Amount: amount,
		Quote:  to,
		Value:  value,
		Rate:   pair.Rate,
	}

	return &ret, nil
}
