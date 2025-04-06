package origin_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/origin"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := origin.Select()
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, origin.All())
}
