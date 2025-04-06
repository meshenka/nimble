package motivation_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/motivation"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		bg := motivation.Select()
		t.Log(bg)
	})
}

func TestAll(t *testing.T) {
	assert.NotNil(t, motivation.All())
}
