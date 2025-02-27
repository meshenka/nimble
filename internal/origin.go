package internal

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

func Origin() string {
	return choose(origins)
}
