package class

import (
	_ "embed"
	"encoding/json"

	"github.com/meshenka/nimble/internal"
)

type Class struct {
	Name         string   `json:"name"`
	KeyStats     []string `json:"key_stats"`
	HitDie       string   `json:"hit_die"`
	StartingHP   int      `json:"starting_hp"`
	Saves        []string `json:"saves"`
	Armor        []string `json:"armor"`
	Weapons      []string `json:"weapons"`
	StartingGear []string `json:"starting_gear"`
}

//go:embed class.json
var classesJSON []byte

var classes []Class

func init() {
	err := json.Unmarshal(classesJSON, &classes)
	if err != nil {
		panic(err)
	}
}

func Select() Class {
	return internal.Choose(classes)
}
