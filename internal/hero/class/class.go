// Package class provides the class for a hero.
package class

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"slices"

	"github.com/meshenka/nimble/internal"
)

// Class represents a hero's class.
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

// Select returns a random class.
func Select(ctx context.Context) Class {
	return internal.Choose(ctx, classes)
}

// All returns all available classes.
func All() []Class {
	return classes
}

// Get returns a class by name.
func Get(name string) (Class, error) {
	index := slices.IndexFunc(classes, func(c Class) bool {
		return c.Name == name
	})
	if index == -1 {
		return Class{}, errors.New("not found")
	}

	return classes[index], nil
}
