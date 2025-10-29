package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/log"
)

// BackgroundsList is a list of backgrounds.
type BackgroundsList struct {
	Backgrounds []background.Background `json:"backgrounds"`
}

// Backgrounds returns a http.Handler that serves a list of all backgrounds.
// @Summary      All backgrounds.
// @Description  List all available backgrounds.
// @Tags         background
// @Produce      json
// @Success      200  {object}  BackgroundsList
// @Router       /backgrounds [get]
// .
func Backgrounds() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, BackgroundsList{
			Backgrounds: background.All(),
		})
	})
}

// GetBackround returns a http.Handler that serves a single background by name.
// @Summary      One Background.
// @Description  Get a background by name.
// @Tags         background
// @Produce      json
// @Param        name   path      string  true  "Background name"
// @Success      200  {object}  background.Background
// @Failure      404
// @Router       /backgrounds/{name} [get]
// .
func GetBackround() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		b, err := background.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("ancestry not found", "error", err, "ancestry", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, b)
	})
}
