package main

import (
	"fmt"

	"github.com/meshenka/nimble/internal/hero"
)

func main() {
	h := hero.NewHero()
	fmt.Print(hero.String(h))
}
