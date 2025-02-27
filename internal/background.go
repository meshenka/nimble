package internal

var backgrounds = []string{
	"Backout of retirement",
	"Devoted protector",
	"Academic Dropout",
	"Made a BAD choice",
	"Haunted past",
	"Ear to the Ground",
	"What? I have been around?",
	"Acrobat",
	"Wild one",
	"Fey Touched",
	"Survivalist",
	"Home at Sea",
	"At Home Underground",
	"Raised by Goblins",
	"History Buff",
	"Former Con Artist",
	"Secretly Undead",
	"Taste for Finer things",
	"Fearless",
	"So Dumb I'm smart sometimes",
	"Wily Underdog",
	"Bumblewise",
	"Accidental Acrobat",
	"Tradesman/Artisan",
}

func Background() string {
	return choose(backgrounds)
}
