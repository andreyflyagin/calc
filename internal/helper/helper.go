package helper

import (
	"calc/internal/model"
	"context"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"sync"
)

// PairsToStrings converts pairs to strings
func PairsToStrings(pairs []model.Pair) []string {
	var ret []string
	for _, pair := range pairs {
		ret = append(ret, fmt.Sprintf("%s/%s", pair.Base, pair.Quote))
	}
	return ret
}

// CreateClient creates a new regrettable client
func CreateClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 3
	retryClient.Logger = &retryableLogger{logger: &log.Logger}
	return retryClient
}

type retryableLogger struct {
	logger *zerolog.Logger
}

func (rl *retryableLogger) Printf(format string, v ...interface{}) {
	rl.logger.Info().Msgf(format, v...)
}

// DoRequest makes a request
func DoRequest(ctx context.Context, retryClient *retryablehttp.Client, url string, path ...interface{}) (jsoniter.Any, error) {
	req, err := retryablehttp.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")

	res, err := retryClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer SafeClose(ctx, res.Body, "response body")

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return jsoniter.Get(body, path...), nil
}

// RunHTTPSenderPool runs a pool of http senders
func RunHTTPSenderPool(workers int, ch <-chan func(*retryablehttp.Client)) {
	var wg sync.WaitGroup
	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			var workerClient = CreateClient()
			for f := range ch {
				f(workerClient)
			}
		}()
	}
	wg.Wait()
}

// SafeClose closes a closer
func SafeClose(ctx context.Context, closer io.Closer, closerName string) {
	if err := closer.Close(); err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("closer error: %s", closerName))
	}
}
