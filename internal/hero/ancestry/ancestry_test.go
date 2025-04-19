package ancestry_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/ancestry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := ancestry.Select(t.Context())
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, ancestry.All())
}

func TestGet(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		have, err := ancestry.Get("Elf")
		require.NoError(t, err)
		assert.NotNil(t, have.Size)
	})
	t.Run("ko", func(t *testing.T) {
		_, err := ancestry.Get("NO_MATCH")
		assert.Error(t, err)
	})
}
