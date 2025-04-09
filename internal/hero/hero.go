package hero

import (
	"fmt"

	"github.com/meshenka/nimble/internal/hero/ancestry"
	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/meshenka/nimble/internal/hero/motivation"
	"github.com/meshenka/nimble/internal/hero/origin"
	"github.com/meshenka/nimble/internal/hero/quirk"
)

type Hero struct {
	Ancestry   ancestry.Ancestry     `json:"ancestry"`
	Class      class.Class           `json:"class"`
	Motivation string                `json:"motivation"`
	Origin     string                `json:"origin"`
	Background background.Background `json:"background"`
	Quirks     []string              `json:"quirks"`
}

func New() Hero {
	return Hero{
		Ancestry:   ancestry.Select(),
		Class:      class.Select(),
		Motivation: motivation.Select(),
		Background: background.Select(),
		Origin:     origin.Select(),
		Quirks:     quirk.Select(),
	}
}

func String(h Hero) string {
	return fmt.Sprintf("I am a %s %s %s from %s who was %s that end up adventuring cause %s",
		h.Quirks[0],
		h.Class.Name,
		h.Ancestry.Name,
		h.Origin,
		h.Background.Name,
		h.Motivation,
	)
}
