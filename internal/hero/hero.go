package hero

import "fmt"

type Hero struct {
	Race       string
	Class      string
	Motivation string
	Origin     string
	Background string
	Quirk      string
}

func NewHero() Hero {
	return Hero{
		Race:       Race(),
		Class:      Class(),
		Motivation: Motivation(),
		Background: Background(),
		Origin:     Origin(),
		Quirk:      Quirk(),
	}
}

func String(h Hero) string {
	return fmt.Sprintf("I am a %s %s %s from %s who was %s that end up adventuring cause %s\n",
		h.Quirk,
		h.Class,
		h.Race,
		h.Origin,
		h.Background,
		h.Motivation,
	)
}
