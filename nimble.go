package nimble

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"sync/atomic"

	"github.com/meshenka/nimble/handler"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/store"
	"github.com/meshenka/nimble/internal/transport"
	"golang.org/x/sync/errgroup"
	_ "modernc.org/sqlite"
)

//go:embed internal/store/schema.sql
var schema string

// Serve run t service.
func Serve(parent context.Context, options ...Option) error {
	cfg := config{
		DatabasePath: "nimble.db",
	}
	for _, opt := range options {
		if err := opt(&cfg); err != nil {
			log.Ctx(parent).Error("invalid configuration", log.Err(err))
			return err
		}
	}

	db, err := sql.Open("sqlite", cfg.DatabasePath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()

	if _, err := db.ExecContext(parent, schema); err != nil {
		return fmt.Errorf("init schema: %w", err)
	}

	s := store.NewStore(db)
	h := handler.New(s)

	httpReady := new(atomic.Bool)

	group, ctx := errgroup.WithContext(parent)

	group.Go(func() error {
		mux := transport.NewRouter(h)
		mw := transport.Use(
			log.HTTPMiddleware(),
		)
		srv := transport.NewServer(mw.Wrap(mux))
		return transport.Serve(ctx, cfg.ApplicationHTTPServerAddr, srv, nil, httpReady)
	})

	if err := group.Wait(); err != nil {
		log.Ctx(parent).Error("shutdown", log.Err(err))
		return err
	}

	return nil
}
