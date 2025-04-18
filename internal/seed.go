package internal

import (
	"math/rand/v2"
	"time"
)

var (
	Seed uint64
	rn   *rand.Rand
)

func init() {
	Seed = uint64(time.Now().UnixNano())   //nolint:gosec // G115 int64->uint64 overflow
	rn = rand.New(rand.NewPCG(Seed, 3999)) //nolint:gosec
}

func Configure(s uint64) {
	Seed = s
	rn = rand.New(rand.NewPCG(Seed, 2999)) //nolint:gosec
}

func Choose[T any](options []T) T {
	idx := rn.IntN(len(options))
	return options[idx]
}
