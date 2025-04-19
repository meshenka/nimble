package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/seeder"
)

type HeroResponse struct {
	Hero     hero.Hero `json:"hero"`
	Sentence string    `json:"sentence"`
	ID       string    `json:"id"`
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
		id := uint64(time.Now().UnixNano()) //nolint:gosec // G115 int64->uint64 overflow
		s := seeder.Configure(id)
		ctx := seeder.WithContext(r.Context(), s)
		h := hero.New(ctx)
		writeJSON(ctx, w, response(h, s))
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

		s := seeder.Configure(id)
		ctx := seeder.WithContext(r.Context(), s)
		h := hero.New(ctx)
		writeJSON(ctx, w, response(h, s))
	})
}

func response(h hero.Hero, s seeder.Rand) HeroResponse {
	return HeroResponse{
		Hero:     h,
		Sentence: hero.String(h),
		ID:       strconv.FormatUint(s.Seed, 10),
	}
}
