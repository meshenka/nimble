package main

import (
	"context"
	"os"

	"github.com/meshenka/nimble"
	"github.com/meshenka/nimble/cmd"
)

func main() {
	if err := cmd.Run(run); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	return nimble.Serve(ctx,
		nimble.WithApplicationServer(
			cmd.Env("APPLICATION_HTTP_ADDR", "127.0.0.1:3000"),
		),
		nimble.WithLogLevel(
			cmd.Env("LOG_LEVEL", "debug"),
		),
	)
}
