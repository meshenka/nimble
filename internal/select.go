// Package internal provides internal helper functions.
package internal

import (
	"context"

	"github.com/meshenka/nimble/internal/seeder"
)

// Choose returns a random element from the given options using the seeder from the context.
func Choose[T any](ctx context.Context, options []T) T {
	idx := seeder.Ctx(ctx).IntN(len(options))
	return options[idx]
}
