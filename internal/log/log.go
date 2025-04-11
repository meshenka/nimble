// Package log is a simple wrapper around log/slog.
package log

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// Configure configures the default logger.
func Configure(level slog.Level) {
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	l := slog.New(h)
	slog.SetDefault(l)
}

type contextKey int

const loggerContextKey contextKey = iota

// WithContext returns a child context holding a reference to the given logger.
func WithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// Ctx returns the logger associated to the current context.
// It returns the default logger if none is found.
func Ctx(ctx context.Context) *slog.Logger {
	logger := ctx.Value(loggerContextKey)
	if impl, ok := logger.(*slog.Logger); ok {
		return impl
	}
	return slog.Default()
}

// Err embeds an error in a slog.Attr for easier structured log.
func Err(err error) slog.Attr {
	return slog.String("error", err.Error())
}

// Writer is an implementation of io.Writer that logs the given data.
func Writer(logger *slog.Logger, level slog.Level) io.Writer {
	return &writer{
		Logger: logger,
		level:  level,
	}
}

type writer struct {
	*slog.Logger
	level slog.Level
}

// See: zerolog.Logger.Write().
func (w *writer) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 && p[n-1] == '\n' {
		p = p[0 : n-1]
	}
	w.Log(context.Background(), w.level, string(p))
	return
}

// HTTPMiddleware is a middleware that adds context to the current application logger
// such as an unique request ID.
func HTTPMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rid := r.Header.Get("X-Correlation-Id")
			if rid == "" {
				rid = uuid.Must(uuid.NewV7()).String()
			}

			ctx := r.Context()
			logger := slog.With(
				slog.String("ctx_request_id", rid),
				slog.String("ctx_method", r.Method),
				slog.String("ctx_path", r.URL.Path),
				slog.String("ctx_pattern", r.Pattern),
				slog.String("ctx_user_agent", r.UserAgent()),
				slog.Time("ctx_received_at", time.Now()),
			)

			logger.Debug("HTTP request received")

			h.ServeHTTP(w, r.WithContext(WithContext(ctx, logger)))
		})
	}
}
