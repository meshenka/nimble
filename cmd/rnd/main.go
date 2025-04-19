package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand/v2"

	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/seeder"
)

func main() {
	// accept a seed for Deterministic setting
	var dst uint64
	flag.Uint64Var(&dst, "seed", rand.Uint64(), "set a seed to have deterministic output")
	flag.Parse()
	s := seeder.Configure(dst)
	ctx := seeder.WithContext(context.Background(), s)
	log.Configure(slog.LevelDebug)
	h := hero.New(ctx)
	fmt.Println(hero.String(h))
	slog.Info("seeded with", "seed", s.Seed)
	slog.Debug("hero details", "hero", h)
}
