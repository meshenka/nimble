package hero

import "github.com/meshenka/nimble/internal"

var races = []string{
	"Human",
	"Human",
	"Human",
	"Elf",
	"Elf",
	"Elf",
	"Dwarf",
	"Dwarf",
	"Dwarf",
	"Halfling",
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
