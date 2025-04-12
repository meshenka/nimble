package handler

import (
	"net/http"

	"github.com/meshenka/nimble/internal/hero/class"
	"github.com/meshenka/nimble/internal/log"
)

type ClassesList struct {
	Classes []class.Class `json:"classes"`
}

// @Summary      All classes
// @Description  Get all classes
// @Tags         class
// @Produce      json
// @Success      200  {object}  ClassesList
// @Router       /classes [get]func Classes() http.Handler {
func Classes() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=5184000, s-maxage=5184000") // 24h
		writeJSON(r.Context(), w, ClassesList{
			Classes: class.All(),
		})
	})
}

// @Summary      Get class
// @Description  Get one class by name
// @Tags         class
// @Param        name   path      string  true  "Class name"
// @Produce      json
// @Success      200  {object}  class.Class
// @Success      404
// @Router       /classes/{name} [get]func Classes() http.Handler {
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
