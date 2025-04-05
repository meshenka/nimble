package background_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := background.Select()
		t.Log(bg)
	})
}
