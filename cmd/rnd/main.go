package main

import (
	"fmt"
	"time"

	"github.com/meshenka/nimble/internal"
)

func main() {
	internal.Configure(uint64(time.Now().UnixNano()))
	hero := NewHero()
	fmt.Printf("I am a %s %s from %s who %s that end up adventuring cause %s\n",
		hero.Class,
		hero.Race,
		hero.Origin,
		hero.Background,
		hero.Motivation,
	)
}

type Hero struct {
	Race       string
	Class      string
	Motivation string
	Origin     string
	Background string
}

func NewHero() Hero {
	return Hero{
		Race:       internal.Race(),
		Class:      internal.Class(),
		Motivation: internal.Motivation(),
		Background: internal.Background(),
		Origin:     internal.Origin(),
	}
}
