// Package http contains generic primitives to create an HTTP server.
package transport

import (
	"context"
	"crypto/tls"
	stdlog "log"
	"log/slog"
	"net"
	"net/http"
	"net/http/pprof"
	"sync/atomic"
	"time"

	"github.com/meshenka/nimble/internal/log"
)

const shutdownTimeout = 20 * time.Second

// Heartbeat returns HTTP 200 if all of  the given values are true or nil. It
// returns HTTP 503 otherwise.
func Heartbeat(ready ...*atomic.Bool) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		status := http.StatusOK
		for i := range ready {
			if !ready[i].Load() {
				status = http.StatusServiceUnavailable
			}
		}
		w.WriteHeader(status)
	}
}

// Expose inner library type for convenience.
var NewServeMux = http.NewServeMux

// NewServer returns a new HTTP server.
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Handler:           handler,
		ErrorLog:          stdlog.New(log.Writer(slog.Default(), slog.LevelDebug), "", 0),
		ReadHeaderTimeout: 10 * time.Second,
	}
}

// Middleware is an HTTP middleware.
type Middleware func(http.Handler) http.Handler

// With is a convenience function to wire middlewares into a handler.
func (f Middleware) Wrap(h http.Handler) http.Handler {
	return f(h)
}

// Use composes multiple middlewares into a single middleware.
func Use(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares); i > 0; i-- {
			h = middlewares[i-1](h)
		}
		return h
	}
}

// Serve routes HTTP requests to handler.
func Serve(ctx context.Context, addr string, srv *http.Server, cfg *tls.Config, ready *atomic.Bool) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func() { _ = l.Close() }()

	log.Ctx(ctx).With(slog.String("address", l.Addr().String())).Debug("starting HTTP server")

	if cfg != nil {
		l = tls.NewListener(l, cfg)
	}

	if ready != nil {
		ready.Store(true)
	}

	sink := make(chan error, 1)

	go func() {
		defer close(sink)
		sink <- srv.Serve(l)
	}()

	select {
	case <-ctx.Done():
		return shutdown(srv)
	case err := <-sink:
		return err
	}
}

func shutdown(srv *http.Server) error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	return srv.Shutdown(ctx)
}

// RegisterProfiler maps HTTP pprof handlers to the given router, using standard URLs.
func RegisterProfiler(mux *http.ServeMux) {
	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
}
