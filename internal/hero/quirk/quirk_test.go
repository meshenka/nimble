package quirk_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/quirk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelect(t *testing.T) {
	have := quirk.Select()
	require.Len(t, have, 3)
	assert.NotZero(t, have[0])
}
