// Package seeder provides a seeded random number generator.
package seeder

import (
	"context"
	"math/rand/v2"
	"time"
)

var (
	// Seed is the seed used for the random number generator.
	Seed          uint64
	rn            *rand.Rand
	defaultSeeder Rand
)

func init() {
	Seed = uint64(time.Now().UnixNano())   //nolint:gosec // G115 int64->uint64 overflow
	rn = rand.New(rand.NewPCG(Seed, 3999)) //nolint:gosec
	defaultSeeder = Rand{
		Seed: Seed,
		rnd:  rn,
	}
}

// Configure configures the seeder with a specific seed.
func Configure(s uint64) Rand {
	Seed = s
	rn = rand.New(rand.NewPCG(Seed, 2999)) //nolint:gosec
	return Rand{
		Seed: s,
		rnd:  rn,
	}
}

// Rand is a seeded random number generator.
type Rand struct {
	rnd  *rand.Rand
	Seed uint64
}

// IntN returns a random integer in [0, n).
func (r Rand) IntN(n int) int {
	return r.rnd.IntN(n)
}

type key int

const rndKey key = iota

// WithContext returns a new context with the seeder.
func WithContext(parent context.Context, seeder Rand) context.Context {
	return context.WithValue(parent, rndKey, seeder)
}

// Ctx returns the Rand seeder from the context, or the default seeder if not found.
func Ctx(ctx context.Context) Rand {
	if seeder, ok := ctx.Value(rndKey).(Rand); ok {
		return seeder
	}
	return defaultSeeder
}
