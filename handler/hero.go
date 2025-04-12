package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero"
)

type HeroResponse struct {
	Hero     hero.Hero `json:"hero"`
	Sentence string    `json:"sentence"`
}

// Hero godoc
// @Summary      Generate a new random hero
// @Description  Generate a new hero character concept
// @Tags         hero
// @Produce      json
// @Success      200  {object}  HeroResponse
// @Router       /heros [get]func Hero() http.Handler {
func Hero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := hero.New()
		writeJSON(r.Context(), w, HeroResponse{
			Hero:     h,
			Sentence: hero.String(h),
		})
	})
}
