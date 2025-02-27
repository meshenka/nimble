package internal

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
