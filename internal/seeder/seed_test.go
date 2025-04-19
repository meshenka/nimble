package seeder_test

import (
	"testing"
	"time"

	"github.com/meshenka/nimble/internal/seeder"
	"github.com/stretchr/testify/assert"
)

func TestConfigure(t *testing.T) {
	assert.NotPanics(t, func() {
		seed := uint64(time.Now().UnixNano()) //nolint:gosec // G115 int64->uint64 overflow
		seeder.Configure(seed)
	})
}

func TestContext(t *testing.T) {
	defaultSeed := seeder.Ctx(t.Context()).Seed

	id := uint64(time.Now().UnixNano()) //nolint:gosec // G115 int64->uint64 overflow
	s := seeder.Configure(id)
	ctx := seeder.WithContext(t.Context(), s)
	assert.NotEqual(t, defaultSeed, seeder.Ctx(ctx).Seed)
	assert.Equal(t, id, s.Seed)
}
