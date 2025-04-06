package origin

import "github.com/meshenka/nimble/internal"

var origins = []string{
	"The Shadow Blight",
	"White Gate",
	"Merivale",
	"King's Reach",
	"Bramblekrag foothills",
	"The Sea of Tears",
	"Highway of Flame",
	"The Stone of Seasons",
	"Farhope, The last harbor",
	"The Skyreach Islands",
	"Elderwild",
	"The Grawling marches",
	"The Withering Dunes",
	"Frostforge mountains",
}

func Select() string {
	return internal.Choose(origins)
}

func All() []string {
	return origins
}
