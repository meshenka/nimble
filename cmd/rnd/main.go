package main

import (
	"fmt"
	"time"

	"github.com/meshenka/nimble/internal"
)

func main() {
	internal.Configure(uint64(time.Now().UnixNano()))
	hero := internal.NewHero()
	fmt.Printf("I am a %s %s %s from %s who was %s that end up adventuring cause %s\n",
		hero.Quirk,
		hero.Class,
		hero.Race,
		hero.Origin,
		hero.Background,
		hero.Motivation,
	)
}
