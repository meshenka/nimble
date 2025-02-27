package internal

import "math/rand/v2"

var (
	Seed uint64
	R    *rand.Rand
)

func Configure(seed uint64) {
	Seed = seed
	R = rand.New(rand.NewPCG(Seed, 2999))
}

func choose(options []string) string {
	idx := R.IntN(len(options))
	return options[idx]
}
