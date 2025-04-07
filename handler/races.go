package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/race"
	"github.com/meshenka/nimble/internal/log"
)

func Races() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, map[string]any{
			"races": race.All(),
		})
	})
}

func GetRace() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		c, err := race.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("race not found", log.Err(err), "race_name", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, c)
	})
}
