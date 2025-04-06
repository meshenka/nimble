package race_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/race"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := race.Select()
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, race.All())
}
