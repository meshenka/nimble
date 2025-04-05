package class_test

import (
	"testing"

	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert.NotPanics(t, func() {
		c := class.Select()
		t.Log(c)
	})
}
