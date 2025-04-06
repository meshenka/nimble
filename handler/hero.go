package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero"
)

func Hero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := hero.New()
		writeJSON(r.Context(), w, map[string]any{
			"hero":     h,
			"sentence": hero.String(h),
		})
	})
}
