package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/meshenka/nimble/internal/log"
)

func Classes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, map[string]any{
			"classes": class.All(),
		})
	})
}

func GetClass() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		c, err := class.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("class not found", log.Err(err), "class_name", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, c)
	})
}
