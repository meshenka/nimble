// Package handler provides http handlers for the application.
package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/ancestry"
	"github.com/meshenka/nimble/internal/log"
)

// AncestriesList is a list of ancestries.
type AncestriesList struct {
	Ancestries []ancestry.Ancestry `json:"ancestries"`
}

// Ancestries returns a http.Handler that serves a list of all ancestries.
// @Summary      All ancestries.
// @Description  Get all ancestries.
// @Tags         ancestry
// @Produce      json
// @Success      200  {object}  AncestriesList
// @Router       /ancestries [get]
// .
func Ancestries() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, AncestriesList{
			Ancestries: ancestry.All(),
		})
	})
}

// GetAncestry returns a http.Handler that serves a single ancestry by name.
// @Summary      Get ancestry.
// @Description  Get one ancestry by name.
// @Tags         ancestry
// @Param        name   path      string  true  "Ancestry name"
// @Produce      json
// @Success      200  {object}  ancestry.Ancestry
// @Success      404
// @Router       /ancestries/{name} [get]
// .
func GetAncestry() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		c, err := ancestry.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("ancestry not found", "error", err, "ancestry", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, c)
	})
}
