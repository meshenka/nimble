package internal

import (
	"math/rand/v2"
	"time"
)

var (
	Seed uint64
	R    *rand.Rand
)

func init() {
	Seed = uint64(time.Now().UnixNano())  //nolint:gosec // G115 int64->uint64 overflow
	R = rand.New(rand.NewPCG(Seed, 3999)) //nolint:gosec
}

func Configure(seed uint64) {
	Seed = seed
	R = rand.New(rand.NewPCG(Seed, 2999)) //nolint:gosec
}

func Choose(options []string) string {
	idx := R.IntN(len(options))
	return options[idx]
}
