package log_test

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/meshenka/nimble/internal/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	t.Skip("TODO")
}

func TestConfigure(t *testing.T) {
	assert.NotPanics(t, func() {
		log.Configure(slog.LevelDebug)
	})
}

func TestContext(t *testing.T) {
	buf := new(bytes.Buffer)
	handler := slog.NewJSONHandler(buf, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler).With(
		slog.Time("test_received_at", time.Now()),
	)

	ctx := log.WithContext(t.Context(), logger)
	log.Ctx(ctx).Info("test", log.Err(errors.New("sentinel_test")))

	have, err := io.ReadAll(buf)
	require.NoError(t, err)
	assert.Contains(t, string(have), "test_received_at")
	assert.Contains(t, string(have), "error")
}
