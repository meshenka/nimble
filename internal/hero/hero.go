package hero

import (
	"fmt"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/meshenka/nimble/internal/hero/motivation"
	"github.com/meshenka/nimble/internal/hero/origin"
	"github.com/meshenka/nimble/internal/hero/quirk"
	"github.com/meshenka/nimble/internal/hero/race"
)

type Hero struct {
	Race       string                `json:"race"`
	Class      class.Class           `json:"class"`
	Motivation string                `json:"motivation"`
	Origin     string                `json:"origin"`
	Background background.Background `json:"background"`
	Quirks     []string              `json:"quirks"`
}

func New() Hero {
	return Hero{
		Race:       race.Select(),
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
		h.Race,
		h.Origin,
		h.Background.Name,
		h.Motivation,
	)
}
