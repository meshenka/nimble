package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/meshenka/nimble/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomHeroHandler(t *testing.T) {
	sut := handler.RandomHero()

	t.Run("random", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/heros", http.NoBody)
		w := httptest.NewRecorder()

		sut.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)

		resp := new(handler.HeroResponse)
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), resp))
		assert.NotEmpty(t, resp.ID)
	})

	t.Run("with seed", func(t *testing.T) {
		var seed uint64 = 9527120102352189307
		want := strconv.FormatUint(seed, 10)
		req := httptest.NewRequest(http.MethodGet, "/api/heros?seed="+want, http.NoBody)
		w := httptest.NewRecorder()

		sut.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)

		resp := new(handler.HeroResponse)
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), resp))
		assert.Equal(t, want, resp.ID)
		assert.Equal(t, "Turtlefolk", resp.Hero.Ancestry.Name)
	})

	t.Run("invalid seed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/heros?seed=not-a-number", http.NoBody)
		w := httptest.NewRecorder()

		sut.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetHeroHandler(t *testing.T) {
	sut := handler.GetHero()

	router := http.NewServeMux()
	router.Handle("GET /api/heros/{id}", sut)
	var seed uint64 = 9527120102352189307
	want := strconv.FormatUint(seed, 10)
	req := httptest.NewRequest(http.MethodGet, "/api/heros/"+want, http.NoBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	resp := new(handler.HeroResponse)
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), resp))
	assert.Equal(t, want, resp.ID)
	assert.Equal(t, "Turtlefolk", resp.Hero.Ancestry.Name)
	assert.Equal(t, "Mage", resp.Hero.Class.Name)
	assert.Equal(t, "Wild One", resp.Hero.Background.Name)
	assert.Contains(t, resp.Hero.Quirks, "Anxious")
}
