package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math/rand/v2"

	"github.com/meshenka/nimble/internal"
	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
)

func main() {
	// accept a seed for Deterministic setting
	var dst uint64
	flag.Uint64Var(&dst, "seed", rand.Uint64(), "set a seed to have deterministic output")
	flag.Parse()
	internal.Configure(dst)
	log.Configure(slog.LevelDebug)
	h := hero.New()
	fmt.Println(hero.String(h))
	slog.Info("seeded with", "seed", dst)
	slog.Debug("hero details", "hero", h)
}
