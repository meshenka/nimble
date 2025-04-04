package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
)

func Hero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := hero.NewHero()
		writeJSON(r.Context(), w, map[string]any{
			"hero":     h,
			"sentence": hero.String(h),
		})
	})
}

func writeJSON(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Ctx(ctx).Error("could not write response body", log.Err(err))
		panic(http.ErrAbortHandler)
	}
}
