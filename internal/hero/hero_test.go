package hero_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	h := hero.New()
	require.NotZero(t, h)

	t.Run("string", func(t *testing.T) {
		resp := hero.String(h)
		assert.NotEqual(t, "", resp)
	})
}
