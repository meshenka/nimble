package internal

import (
	"math/rand/v2"
	"time"
)

var (
	seed uint64
	rn   *rand.Rand
)

func init() {
	seed = uint64(time.Now().UnixNano())   //nolint:gosec // G115 int64->uint64 overflow
	rn = rand.New(rand.NewPCG(seed, 3999)) //nolint:gosec
}

func Configure(s uint64) {
	seed = s
	rn = rand.New(rand.NewPCG(seed, 2999)) //nolint:gosec
}

func Choose(options []string) string {
	idx := rn.IntN(len(options))
	return options[idx]
}
