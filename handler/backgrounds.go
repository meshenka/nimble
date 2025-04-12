package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/log"
)

type BackgroundsList struct {
	Backgrounds []background.Background `json:"backgrounds"`
}

// @Summary      All backgrounds
// @Description  List all available backgrounds
// @Tags         background
// @Produce      json
// @Success      200  {object}  BackgroundsList
// @Router       /backgrounds [get]func Backgrounds() http.Handler {
func Backgrounds() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, BackgroundsList{
			Backgrounds: background.All(),
		})
	})
}

// @Summary      One Background
// @Description  Get a background by name
// @Tags         background
// @Produce      json
// @Param        name   path      string  true  "Background name"
// @Success      200  {object}  background.Background
// @Failure      404
// @Router       /backgrounds/{name} [get]func GetBackround() http.Handler {
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
