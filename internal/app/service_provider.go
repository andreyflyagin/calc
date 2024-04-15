package app

import (
	"calc/internal/config"
	"calc/internal/repository"
	"calc/internal/repository/currency"
	"calc/internal/service"
	conversionService "calc/internal/service/conversion"
	"calc/internal/service/exchangeratesprovider"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	convertRepository    repository.CurrencyRepository
	conversionService    service.ConversionService
	exchangeRatesService service.ExchangeRatesProviderService
	db                   *sqlx.DB
	cfg                  *config.Config
}

func newServiceProvider(db *sqlx.DB, cfg *config.Config) *serviceProvider {
	return &serviceProvider{db: db, cfg: cfg}
}

func (s *serviceProvider) ConvertRepository() repository.CurrencyRepository {
	if s.convertRepository == nil {

		s.convertRepository = currency.NewRepository(s.db)
	}

	return s.convertRepository
}

func (s *serviceProvider) ConvertService() service.ConversionService {
	if s.conversionService == nil {
		s.conversionService = conversionService.NewService(s.ConvertRepository(), s.ExchangeRatesService())
	}

	return s.conversionService
}

func (s *serviceProvider) ExchangeRatesService() service.ExchangeRatesProviderService {
	if s.exchangeRatesService == nil {
		s.exchangeRatesService = exchangeratesprovider.NewProvider(s.cfg.DataProvider)
	}

	return s.exchangeRatesService
}
