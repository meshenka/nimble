package internal

import (
	"context"

	"github.com/meshenka/nimble/internal/seeder"
)

func Choose[T any](ctx context.Context, options []T) T {
	idx := seeder.Ctx(ctx).IntN(len(options))
	return options[idx]
}
