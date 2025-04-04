// Package cmd regroups reusable components to build CLIs.
package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Env returns the value of the given environment variable or uses the provided
// fallback value instead.
func Env(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return fallback
}

// Run runs the given function with a context that is closed as soon as an OS
// signal is caught.
func Run(f func(context.Context) error) error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer cancel()

	return f(ctx)
}
