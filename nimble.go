package nimble

import (
	"context"
	"sync/atomic"

	"github.com/meshenka/nimble/handler"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/transport"
	"golang.org/x/sync/errgroup"
)

// Serve starts the HTTP server.
func Serve(parent context.Context, options ...Option) error {
	// return errors.New("not implemented")
	var cfg config
	for _, opt := range options {
		if err := opt(&cfg); err != nil {
			log.Ctx(parent).Error("invalid configuration", log.Err(err))
			return err
		}
	}

	httpReady := new(atomic.Bool)

	group, ctx := errgroup.WithContext(parent)

	group.Go(func() error {
		mux := transport.NewServeMux()
		mux.Handle("GET /api/heros", handler.Hero())
		mux.Handle("GET /api/classes", handler.Classes())
		mux.Handle("GET /api/classes/{name}", handler.GetClass())
		mux.Handle("GET /api/ancestries", handler.Races())
		mux.Handle("GET /api/ancestries/{name}", handler.GetRace())
		mw := transport.Use(log.HTTPMiddleware())
		srv := transport.NewServer(mw.Wrap(mux))
		return transport.Serve(ctx, cfg.ApplicationHTTPServerAddr, srv, nil, httpReady)
	})

	if err := group.Wait(); err != nil {
		log.Ctx(parent).Error("shutdown", log.Err(err))
		return err
	}

	return nil
}
