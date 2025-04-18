package handler

import (
	"net/http"
	"strconv"

	"github.com/meshenka/nimble/internal"
	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
)

type HeroResponse struct {
	Hero     hero.Hero `json:"hero"`
	Sentence string    `json:"sentence"`
	ID       uint64    `json:"id"`
}

// RandomHero godoc
// @Summary      Generate a new random hero
// @Description  Generate a new hero character concept
// @Tags         hero
// @Produce      json
// @Success      200  {object}  HeroResponse
// @Router       /heros [get]func RandomHero() http.Handler {
func RandomHero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := hero.New()
		writeJSON(r.Context(), w, HeroResponse{
			Hero:     h,
			Sentence: hero.String(h),
			ID:       internal.Seed,
		})
	})
}

// GetHero godoc
// @Summary      Get a specific hero from it's id
// @Description  Every random hero is generated from a seed. Once seed is set, the generation is deterministic.
// @Tags         hero
// @Produce      json
// @Success      200  {object}  HeroResponse
// @Router       /heros/{id} [get]func GetHero() http.Handler {
func GetHero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.PathValue("id")
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Ctx(r.Context()).Error("invalid id", log.Err(err), "hero_id", v)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		internal.Configure(id)
		h := hero.New()
		writeJSON(r.Context(), w, HeroResponse{
			Hero:     h,
			Sentence: hero.String(h),
			ID:       internal.Seed,
		})
	})
}
