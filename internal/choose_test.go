package internal_test

import (
	"testing"

	"github.com/meshenka/nimble/internal"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {

	options := []int{4, 5, 6}
	have := internal.Choose(t.Context(), options)
	assert.Greater(t, have, 3)
	assert.Less(t, have, 7)
}
