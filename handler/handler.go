package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/store"
)

// Handler handles HTTP requests.
type Handler struct {
	store *store.Store
}

// New creates a new Handler with the given store.
func New(s *store.Store) *Handler {
	return &Handler{store: s}
}

func writeJSON(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Ctx(ctx).Error("could not write response body", log.Err(err))
		panic(http.ErrAbortHandler)
	}
}
