package hero

import "github.com/meshenka/nimble/internal"

var races = []string{
	"Human",
	"Human",
	"Elf",
	"Elf",
	"Dwarf",
	"Dwarf",
	"Halfling",
	"Halfling",
	"Gnome",
	"Gnome",
	"Bunbun",
	"Orc",
	"Minotaur",
	"Celestial",
	"Fiendkin",
	"Half-Giant",
	"Ooze",
	"Planarbeing",
	"Goblin",
	"Kobold",
	"Birdfolk",
	"Changeling",
	"Crystalborn",
	"Dryad",
	"Ratfolk",
	"Turtlefolk",
	"Stoatling",
}

func Race() string {
	return internal.Choose(races)
}
