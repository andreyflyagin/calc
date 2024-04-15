package main

import (
	"calc/internal/app"
	"calc/internal/logger"
	"context"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.SetupZeroLog()
	log.Info().Msg("initialize app")
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize app")
	}

	err = a.Run(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run app")
	}

}
