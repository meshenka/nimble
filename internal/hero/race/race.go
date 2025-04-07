package race

import (
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

type Race struct {
	Name        string   `json:"name"`
	Size        string   `json:"size"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Capacity    Capacity `json:"capacity"`
}

//go:embed race.json
var racesJSON []byte
var races []Race

func init() {
	err := json.Unmarshal(racesJSON, &races)
	if err != nil {
		panic(err)
	}
}

func Select() Race {
	return internal.Choose(races)
}

func All() []Race {
	return races
}

func Get(name string) (Race, error) {
	index := slices.IndexFunc(races, func(r Race) bool {
		return r.Name == name
	})
	if index == -1 {
		return Race{}, errors.New("not found")
	}

	return races[index], nil
}
