package store_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/meshenka/nimble"
	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/store"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func setupStore(t *testing.T) *store.Store {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	t.Cleanup(func() { _ = db.Close() })

	goose.SetBaseFS(nimble.Migrations)
	require.NoError(t, goose.SetDialect("sqlite3"))
	require.NoError(t, goose.Up(db, "migrations"))

	return store.NewStore(db)
}

func TestStore_SaveAndGet(t *testing.T) {
	s := setupStore(t)
	ctx := context.Background()

	// Generate a hero to save
	h := hero.New(ctx)
	h.Quirks = []string{"Loud", "Fast"} // Ensure predictable quirks for test

	// Save
	saved, err := s.SaveHero(ctx, h)
	require.NoError(t, err)
	assert.NotEqual(t, [16]byte{}, saved.ID)

	// Get
	got, err := s.GetHero(ctx, saved.ID)
	require.NoError(t, err)
	assert.Equal(t, saved.ID, got.ID)
	assert.Equal(t, h.Ancestry.Name, got.Ancestry.Name)
	assert.Equal(t, h.Class.Name, got.Class.Name)
	assert.Equal(t, h.Quirks, got.Quirks)
}

func TestStore_ListHeroes(t *testing.T) {
	s := setupStore(t)
	ctx := context.Background()

	// Save two heroes
	h1 := hero.New(ctx)
	_, err := s.SaveHero(ctx, h1)
	require.NoError(t, err)

	h2 := hero.New(ctx)
	_, err = s.SaveHero(ctx, h2)
	require.NoError(t, err)

	// List
	list, err := s.ListHeroes(ctx)
	require.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestStore_GetNotFound(t *testing.T) {
	s := setupStore(t)
	ctx := context.Background()

	h := hero.New(ctx) // Not saved
	_, err := s.GetHero(ctx, h.ID)
	assert.Error(t, err)
}
