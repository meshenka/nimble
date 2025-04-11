package class_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		c := class.Select()
		t.Log(c)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, class.All())
}

func TestGet(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		have, err := class.Get("Commander")
		require.NoError(t, err)
		assert.NotNil(t, have.HitDie)
	})
	t.Run("ko", func(t *testing.T) {
		_, err := class.Get("NO_MATCH")
		assert.Error(t, err)
	})
}
