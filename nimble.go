package nimble

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"sync/atomic"

	"github.com/meshenka/nimble/handler"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/store"
	"github.com/meshenka/nimble/internal/transport"
	"github.com/pressly/goose/v3"
	"golang.org/x/sync/errgroup"
	_ "modernc.org/sqlite" // sqlite driver
)

// Migrations is the embedded filesystem containing database migrations.
//
//go:embed migrations/*.sql
var Migrations embed.FS

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
	defer db.Close() //nolint:errcheck

	goose.SetBaseFS(Migrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("run migrations: %w", err)
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
