package handler_test

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/meshenka/nimble"
	"github.com/meshenka/nimble/handler"
	"github.com/meshenka/nimble/internal/store"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func setup(t *testing.T) *handler.Handler {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	goose.SetBaseFS(nimble.Migrations)
	require.NoError(t, goose.SetDialect("sqlite3"))
	require.NoError(t, goose.Up(db, "migrations"))

	s := store.NewStore(db)
	return handler.New(s)
}

func TestRandomHeroHandler(t *testing.T) {
	h := setup(t)
	sut := h.RandomHero()

	t.Run("random", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/heros", http.NoBody)
		w := httptest.NewRecorder()

		sut.ServeHTTP(w, req)
		require.Equal(t, http.StatusOK, w.Code)

		resp := new(handler.HeroResponse)
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), resp))
		assert.NotEqual(t, [16]byte{}, resp.ID)
	})

	t.Run("invalid seed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/heros?seed=not-a-number", http.NoBody)
		w := httptest.NewRecorder()

		sut.ServeHTTP(w, req)
		require.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetHeroHandler(t *testing.T) {
	h := setup(t)
	sut := h.GetHero()

	// 1. Create a hero first
	reqCreate := httptest.NewRequest(http.MethodGet, "/api/heros", http.NoBody)
	wCreate := httptest.NewRecorder()
	h.RandomHero().ServeHTTP(wCreate, reqCreate)
	require.Equal(t, http.StatusOK, wCreate.Code)

	var respCreate handler.HeroResponse
	require.NoError(t, json.Unmarshal(wCreate.Body.Bytes(), &respCreate))

	// 2. Try to get it
	router := http.NewServeMux()
	router.Handle("GET /api/heros/{id}", sut)

	req := httptest.NewRequest(http.MethodGet, "/api/heros/"+respCreate.ID.String(), http.NoBody)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	resp := new(handler.HeroResponse)
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), resp))
	assert.Equal(t, respCreate.ID, resp.ID)
	assert.Equal(t, respCreate.Hero.Ancestry.Name, resp.Hero.Ancestry.Name)
}
