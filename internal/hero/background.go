package hero

import "github.com/meshenka/nimble/internal"

// who was ..
var backgrounds = []string{
	"Backout of retirement",
	"a Devoted protector",
	"an Academic Dropout",
	"Made a BAD choice",
	"Haunted by my past",
	"Ear to the Ground",
	"suffering Amnesia", // "What? I have been around?"
	"an Acrobat",
	"a Wild one",
	"Fey Touched",
	"a Survivalist",
	"at Home at Sea",
	"at Home Underground",
	"Raised by Goblins",
	"a History Buff",
	"a Former Con Artist",
	"Secretly Undead",
	"apprecating a Taste for Finer things",
	"Fearless",
	"So Dumb I'm smart sometimes",
	"a Wily Underdog",
	"Bumblewise",
	"an Accidental Acrobat",
	"a Tradesman",
	"an Artisan",
}

func Background() string {
	return internal.Choose(backgrounds)
}
