package background

import (
	_ "embed"
	"encoding/json"
	"errors"
	"slices"

	"github.com/meshenka/nimble/internal"
)

type Background struct {
	Name         string   `json:"name"`
	Requirements string   `json:"requirements"`
	Abilities    []string `json:"abilities"`
}

//go:embed backgrounds.json
var backgroundsJSON []byte
var bgs []Background

func init() {
	err := json.Unmarshal(backgroundsJSON, &bgs)
	if err != nil {
		panic(err)
	}
}

// Select generate a random Background.
func Select() Background {
	return internal.Choose(bgs)
}

// All returns all available backgrounds.
func All() []Background {
	return bgs
}

// Get finds a single background matching by name.
func Get(name string) (Background, error) {
	index := slices.IndexFunc(bgs, func(b Background) bool {
		return b.Name == name
	})
	if index == -1 {
		return Background{}, errors.New("not found")
	}

	return bgs[index], nil
}
