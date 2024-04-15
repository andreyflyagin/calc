package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetupZeroLog sets up zero log logger
func SetupZeroLog() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339Nano,
	})
	zerolog.TimeFieldFormat = time.RFC3339Nano
}
