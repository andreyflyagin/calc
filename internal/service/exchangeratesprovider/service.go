package exchangeratesprovider

import (
	"calc/internal/config"
	"calc/internal/helper"
	"calc/internal/model"
	"calc/internal/service"
	"context"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
	"math"
	"strings"
)

const (
	fiatCurrenciesReq   = "https://api.fastforex.io/currencies?api_key=%s"
	cryptoCurrenciesReq = "https://api.fastforex.io/crypto/currencies?api_key=%s"
	cryptoPairsReq      = "https://api.fastforex.io/crypto/pairs?api_key=%s"

	cryptoFetchPricesReq = "https://api.fastforex.io/crypto/fetch-prices?pairs=%s&api_key=%s"
	cryptoFetchOneReq    = "https://api.fastforex.io/fetch-one?from=%s&to=%s&api_key=%s"
)

type pair struct {
	Base  string
	Quote string
}

var _ service.ExchangeRatesProviderService = (*provider)(nil)

type provider struct {
	cfg config.DataProvider
}

// NewProvider creates a new exchange rates provider
func NewProvider(cfg config.DataProvider) service.ExchangeRatesProviderService {
	return &provider{cfg}
}

// GetData gets data
func (srv *provider) GetData(ctx context.Context, currencies []model.Currency) ([]model.PairRates, error) {
	var client = helper.CreateClient()

	fiatCurrencies, err := srv.getCurrencies(ctx, client, fiatCurrenciesReq)
	if err != nil {
		return nil, err
	}

	cryptoCurrencies, err := srv.getCurrencies(ctx, client, cryptoCurrenciesReq)
	if err != nil {
		return nil, err
	}

	cryptoPairs, err := srv.getCryptoPairs(ctx, client)
	if err != nil {
		return nil, err
	}

	var crypto, fiat []string

	for _, currency := range currencies {
		if _, ok := cryptoCurrencies[currency.Symbol]; ok {
			crypto = append(crypto, currency.Symbol)
		} else if _, ok := fiatCurrencies[currency.Symbol]; ok {
			fiat = append(fiat, currency.Symbol)
		}
	}
	log.Debug().Msgf("symbols: %v", currencies)
	log.Debug().Msgf("crypto: %v", crypto)
	log.Debug().Msgf("fiat: %v", fiat)

	ch := make(chan func(*retryablehttp.Client))
	results := make(chan model.PairRates)

	go func() {
		defer close(ch)
		srv.submitRequests(ctx, crypto, fiat, cryptoPairs, ch, results)
	}()

	go func() {
		defer close(results)
		helper.RunHTTPSenderPool(10, ch)
	}()

	var result []model.PairRates
	for r := range results {
		result = append(result, r)
	}

	log.Debug().Msgf("availaible pairs: %v", result)
	return result, nil
}

func (srv *provider) submitRequests(ctx context.Context,
	crypto, fiat []string,
	cryptoPairs map[string]pair,
	ch chan<- func(*retryablehttp.Client),
	results chan<- model.PairRates,
) {
	var (
		pairReq []model.Pair
		oneReq  []model.Pair
	)

	for _, c := range crypto {
		for _, f := range fiat {
			if _, ok := cryptoPairs[fmt.Sprintf("%s/%s", c, f)]; ok {
				pairReq = append(pairReq, model.Pair{Base: c, Quote: f})
			} else {
				oneReq = append(oneReq, model.Pair{Base: c, Quote: f})
			}
		}
	}

	for _, f := range fiat {
		for _, c := range crypto {
			oneReq = append(oneReq, model.Pair{Base: f, Quote: c})
		}
	}

	pairReqCount := int(math.Ceil(float64(len(pairReq)) / float64(10)))

	for i := range pairReqCount {
		m := i * 10
		n := 10 * (i + 1)
		if n > len(pairReq) {
			n = len(pairReq)
		}

		block := pairReq[m:n]

		ch <- func(client *retryablehttp.Client) {
			pairRates, resultErr := srv.fetchPrices(ctx, client, block)
			if resultErr != nil {
				log.Error().Err(resultErr).Msgf("failed to fetch price for block")
			}

			log.Debug().Msgf("add pairs: %v", pairRates)

			for _, p := range pairRates {
				results <- p
			}
		}
	}

	for _, p := range oneReq {
		p := p
		ch <- func(client *retryablehttp.Client) {
			rate, resultErr := srv.fetchOne(ctx, client, p.Base, p.Quote)
			if resultErr != nil {
				log.Error().Err(resultErr).Msgf("failed to fetch price for %s/%s", p.Base, p.Quote)
			}

			log.Debug().Msgf("add pair: %v %v", p, rate)

			results <- model.PairRates{Pair: p, Rate: rate}
		}
	}
}

func (srv *provider) getCurrencies(ctx context.Context, client *retryablehttp.Client, req string) (map[string]string, error) {
	dat := make(map[string]string)

	v, err := helper.DoRequest(ctx, client, fmt.Sprintf(req, srv.cfg.Secret), "currencies")
	if err != nil {
		return nil, err
	}

	v.ToVal(&dat)

	return dat, nil
}

func (srv *provider) getCryptoPairs(ctx context.Context, client *retryablehttp.Client) (map[string]pair, error) {
	dat := make(map[string]pair)

	v, err := helper.DoRequest(ctx, client, fmt.Sprintf(cryptoPairsReq, srv.cfg.Secret), "pairs")
	if err != nil {
		return nil, err
	}

	v.ToVal(&dat)

	return dat, nil
}

func (srv *provider) fetchPrices(ctx context.Context, client *retryablehttp.Client, pairs []model.Pair) ([]model.PairRates, error) {
	pairsConcatenated := strings.Join(helper.PairsToStrings(pairs), ",")

	dat := make(map[string]decimal.Decimal)

	v, err := helper.DoRequest(ctx, client, fmt.Sprintf(cryptoFetchPricesReq, pairsConcatenated, srv.cfg.Secret), "prices")
	if err != nil {
		return nil, err
	}

	v.ToVal(&dat)

	ret := make([]model.PairRates, 0, len(pairs))

	for _, pair := range pairs {
		if v, ok := dat[fmt.Sprintf("%s/%s", pair.Base, pair.Quote)]; ok {
			ret = append(ret, model.PairRates{Pair: pair, Rate: v})
		} else {
			log.Error().Msgf("no such currency: %s/%s", pair.Base, pair.Quote)
		}
	}

	return ret, nil
}

func (srv *provider) fetchOne(ctx context.Context, client *retryablehttp.Client, from, to string) (decimal.Decimal, error) {
	dat := make(map[string]decimal.Decimal)

	v, err := helper.DoRequest(ctx, client, fmt.Sprintf(cryptoFetchOneReq, from, to, srv.cfg.Secret), "result")
	if err != nil {
		return decimal.Decimal{}, err
	}

	v.ToVal(&dat)

	ret, ok := dat[to]
	if !ok {
		return decimal.NewFromInt(0), fmt.Errorf("no such currency: %s", to)
	}

	return ret, nil
}
