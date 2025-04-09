package ancestry_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/ancestry"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := ancestry.Select()
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, ancestry.All())
}
