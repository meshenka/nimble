package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/log"
)

func Backgrounds() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, map[string]any{
			"backgrounds": background.All(),
		})
	})
}

func GetBackround() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		b, err := background.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("ancestry not found", log.Err(err), "ancestry", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, b)
	})
}
