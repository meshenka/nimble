package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/meshenka/nimble/internal/hero"
	"github.com/meshenka/nimble/internal/hero/ancestry"
	"github.com/meshenka/nimble/internal/hero/background"
	"github.com/meshenka/nimble/internal/hero/class"
)

type Store struct {
	db *sql.DB
	q  *Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		q:  New(db),
	}
}

func (s *Store) SaveHero(ctx context.Context, h hero.Hero) (hero.Hero, error) {
	if h.ID == uuid.Nil {
		h.ID = uuid.Must(uuid.NewV7())
	}

	quirksJSON, err := json.Marshal(h.Quirks)
	if err != nil {
		return hero.Hero{}, fmt.Errorf("marshal quirks: %w", err)
	}

	_, err = s.q.CreateHero(ctx, CreateHeroParams{
		ID:             h.ID[:],
		AncestryName:   h.Ancestry.Name,
		ClassName:      h.Class.Name,
		Motivation:     h.Motivation,
		Origin:         h.Origin,
		BackgroundName: h.Background.Name,
		Quirks:         string(quirksJSON),
	})
	if err != nil {
		return hero.Hero{}, fmt.Errorf("insert hero: %w", err)
	}

	return h, nil
}

func (s *Store) GetHero(ctx context.Context, id uuid.UUID) (hero.Hero, error) {
	row, err := s.q.GetHero(ctx, id[:])
	if err != nil {
		return hero.Hero{}, fmt.Errorf("get hero: %w", err)
	}

	return s.toDomain(row)
}

func (s *Store) ListHeroes(ctx context.Context) ([]hero.Hero, error) {
	rows, err := s.q.ListHeroes(ctx)
	if err != nil {
		return nil, fmt.Errorf("list heroes: %w", err)
	}

	heroes := make([]hero.Hero, 0, len(rows))
	for _, row := range rows {
		h, err := s.toDomain(row)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, h)
	}
	return heroes, nil
}

func (s *Store) toDomain(row Hero) (hero.Hero, error) {
	id, err := uuid.FromBytes(row.ID)
	if err != nil {
		return hero.Hero{}, fmt.Errorf("parse uuid: %w", err)
	}

	a, err := ancestry.Get(row.AncestryName)
	if err != nil {
		return hero.Hero{}, err
	}
	b, err := background.Get(row.BackgroundName)
	if err != nil {
		return hero.Hero{}, err
	}
	c, err := class.Get(row.ClassName)
	if err != nil {
		return hero.Hero{}, err
	}

	var quirks []string
	if err := json.Unmarshal([]byte(row.Quirks), &quirks); err != nil {
		return hero.Hero{}, fmt.Errorf("unmarshal quirks: %w", err)
	}

	return hero.Hero{
		ID:         id,
		Ancestry:   a,
		Class:      c,
		Motivation: row.Motivation,
		Origin:     row.Origin,
		Background: b,
		Quirks:     quirks,
	}, nil
}
