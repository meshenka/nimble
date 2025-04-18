package nimble

import (
	"context"
	"net/http"
	"sync/atomic"

	"github.com/meshenka/nimble/handler"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/transport"
	"golang.org/x/sync/errgroup"
)

// @title           Who my f*cking Nimble 5e character is?
// @version         1.0
// @description     Instant random character generator
// @termsOfService  http://swagger.io/terms/
// @contact.name   Meshenka
// @contact.email  meshee.knight@gmail.com
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
		mux.Handle("GET /api/heros", handler.RandomHero())
		mux.Handle("GET /api/heros/{id}", handler.GetHero())
		mux.Handle("GET /api/classes", handler.Classes())
		mux.Handle("GET /api/classes/{name}", handler.GetClass())
		mux.Handle("GET /api/ancestries", handler.Ancestries())
		mux.Handle("GET /api/ancestries/{name}", handler.GetAncestry())
		mux.Handle("GET /api/backgrounds", handler.Backgrounds())
		mux.Handle("GET /api/backgrounds/{name}", handler.GetBackround())
		// Create a file server handler for the static assets directory
		fs := http.FileServer(http.Dir("./public"))
		mux.Handle("GET /", http.StripPrefix("/", fs))
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
