package seeder

import (
	"context"
	"math/rand/v2"
	"time"
)

var (
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

func Configure(s uint64) Rand {
	Seed = s
	rn = rand.New(rand.NewPCG(Seed, 2999)) //nolint:gosec
	return Rand{
		Seed: s,
		rnd:  rn,
	}
}

type Rand struct {
	rnd  *rand.Rand
	Seed uint64
}

func (r Rand) IntN(n int) int {
	return r.rnd.IntN(n)
}

type key int

const rndKey key = iota

func WithContext(parent context.Context, seeder Rand) context.Context {
	return context.WithValue(parent, rndKey, seeder)
}

func Ctx(ctx context.Context) Rand {
	seeder := ctx.Value(rndKey)
	if seeder != nil {
		return seeder.(Rand)
	}
	return defaultSeeder
}
