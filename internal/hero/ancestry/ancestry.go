package ancestry

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"slices"

	"github.com/meshenka/nimble/internal"
)

type Capacity struct {
	Name    string   `json:"name"`
	Effects []string `json:"effects"`
}

type Ancestry struct {
	Name        string   `json:"name"`
	Size        string   `json:"size"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Capacity    Capacity `json:"capacity"`
}

//go:embed ancestries.json
var ancestriesJSON []byte
var ancestries []Ancestry

func init() {
	err := json.Unmarshal(ancestriesJSON, &ancestries)
	if err != nil {
		panic(err)
	}
}

// Select choose randomly an ancestry.
func Select(ctx context.Context) Ancestry {
	return internal.Choose(ctx, ancestries)
}

// All returns all available ancestries.
func All() []Ancestry {
	return ancestries
}

// Get finds an ancestry by name.
func Get(name string) (Ancestry, error) {
	index := slices.IndexFunc(ancestries, func(r Ancestry) bool {
		return r.Name == name
	})
	if index == -1 {
		return Ancestry{}, errors.New("not found")
	}

	return ancestries[index], nil
}
