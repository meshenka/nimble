package background_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := background.Select(t.Context())
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, background.All())
}

func TestGet(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		have, err := background.Get("Raised by Goblins")
		require.NoError(t, err)
		assert.NotNil(t, have.Requirements)
	})
	t.Run("ko", func(t *testing.T) {
		_, err := background.Get("NO_MATCH")
		assert.Error(t, err)
	})
}
