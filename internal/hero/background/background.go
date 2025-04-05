package background

import (
	_ "embed"
	"encoding/json"

	"github.com/meshenka/nimble/internal"
)

type Background struct {
	Name         string `json:"name"`
	Requirements string `json:"requirements"`
	Abilities    string `json:"abilites"`
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

func Select() Background {
	return internal.Choose(bgs)
}
