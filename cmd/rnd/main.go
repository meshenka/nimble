package main

import (
	"flag"
	"fmt"
	"math/rand/v2"

	"github.com/meshenka/nimble/internal"
	"github.com/meshenka/nimble/internal/hero"
)

func main() {
	// accept a seed for Deterministic setting
	var dst uint64
	flag.Uint64Var(&dst, "seed", rand.Uint64(), "set a seed to have deterministic output")
	flag.Parse()
	internal.Configure(dst)
	h := hero.NewHero()
	fmt.Print(hero.String(h))
}
