package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/log"
	"github.com/meshenka/nimble/internal/seeder"
)

// HeroResponse is the response for a hero request.
type HeroResponse struct {
	Hero     hero.Hero `json:"hero"`
	Sentence string    `json:"sentence"`
	ID       uuid.UUID `json:"id"`
}

// RandomHero godoc.
//
// @Summary      Generate a new hero.
// @Description  Generate a new hero character concept. If 'seed' query parameter is provided, it returns a deterministic hero for that seed.
// @Tags         hero
// @Produce      json
// @Param        seed  query     string  false  "Optional seed for deterministic generation"
// @Success      200  {object}  HeroResponse
// @Router       /heros [get]
// .
func (h *Handler) RandomHero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id uint64
		if s := r.URL.Query().Get("seed"); s != "" {
			var err error
			id, err = strconv.ParseUint(s, 10, 64)
			if err != nil {
				log.Ctx(r.Context()).Error("invalid seed", "error", err, "seed", s)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		} else {
			id = uint64(time.Now().UnixNano())
		}
		s := seeder.Configure(id)
		ctx := seeder.WithContext(r.Context(), s)
		heroData := hero.New(ctx)

		savedHero, err := h.store.SaveHero(ctx, heroData)
		if err != nil {
			log.Ctx(ctx).Error("could not save hero", log.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeJSON(ctx, w, response(savedHero))
	})
}

// GetHero godoc.
//
// @Summary      Get a specific hero from it's id.
// @Description  Retrieve a hero by its UUID from the database.
// @Tags         hero
// @Produce      json
// @Success      200  {object}  HeroResponse
// @Router       /heros/{id} [get]
// .
func (h *Handler) GetHero() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := r.PathValue("id")
		id, err := uuid.Parse(v)
		if err != nil {
			log.Ctx(r.Context()).Error("invalid id", "error", err, "hero_id", v)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		heroData, err := h.store.GetHero(r.Context(), id)
		if err != nil {
			log.Ctx(r.Context()).Error("could not find hero", "error", err, "hero_id", id)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		writeJSON(r.Context(), w, response(heroData))
	})
}

func response(h hero.Hero) HeroResponse {
	return HeroResponse{
		Hero:     h,
		Sentence: hero.String(h),
		ID:       h.ID,
	}
}
