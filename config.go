package nimble

import (
	"fmt"
	"log/slog"

	"github.com/meshenka/nimble/internal/log"
)

// Option is an application option.
type Option func(*config) error

type config struct {
	ApplicationHTTPServerAddr string
}

// WithApplicationServer configures the application server addresses.
func WithApplicationServer(http string) Option {
	return func(cfg *config) error {
		cfg.ApplicationHTTPServerAddr = http
		return nil
	}
}

// WithLogLevel configures the log level.
func WithLogLevel(lvl string) Option {
	return func(_ *config) error {
		level := new(slog.Level)
		if err := level.UnmarshalText([]byte(lvl)); err != nil {
			return fmt.Errorf("could not parse log level: %w", err)
		}
		log.Configure(*level)
		return nil
	}
}
