package hero

import (
	"fmt"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/hero/quirk"
)

type Hero struct {
	Race       string                `json:"race"`
	Class      string                `json:"class"`
	Motivation string                `json:"motivation"`
	Origin     string                `json:"origin"`
	Background background.Background `json:"background"`
	Quirks     []string              `json:"quirks"`
}

func NewHero() Hero {
	return Hero{
		Race:       Race(),
		Class:      Class(),
		Motivation: Motivation(),
		Background: background.Select(),
		Origin:     Origin(),
		Quirks:     quirk.Select(),
	}
}

func String(h Hero) string {
	return fmt.Sprintf("I am a %s %s %s from %s who was %s that end up adventuring cause %s",
		h.Quirks[0],
		h.Class,
		h.Race,
		h.Origin,
		h.Background.Name,
		h.Motivation,
	)
}
