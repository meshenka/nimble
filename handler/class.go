package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/meshenka/nimble/internal/log"
)

// ClassesList is a list of classes.
type ClassesList struct {
	Classes []class.Class `json:"classes"`
}

// Classes returns a http.Handler that serves a list of all classes.
//
// @Summary      All classes.
// @Description  Get all classes.
// @Tags         class
// @Produce      json
// @Success      200  {object}  ClassesList
// @Router       /classes [get]
// .
func Classes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, ClassesList{
			Classes: class.All(),
		})
	})
}

// GetClass returns a http.Handler that serves a single class by name.
//
// @Summary      Get class.
// @Description  Get one class by name.
// @Tags         class
// @Param        name   path      string  true  "Class name"
// @Produce      json
// @Success      200  {object}  class.Class
// @Success      404
// @Router       /classes/{name} [get]
// .
func GetClass() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		name := r.PathValue("name")

		c, err := class.Get(name)
		if err != nil {
			log.Ctx(r.Context()).Error("class not found", "error", err, "class_name", name)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		writeJSON(r.Context(), w, c)
	})
}
