package transport

import (
	"net/http"

	"github.com/meshenka/nimble/handler"
)

// NewRouter creates and configures a new router instance.
func NewRouter() *http.ServeMux {
	mux := NewServeMux()
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

	return mux
}
