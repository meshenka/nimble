package internal

var classes = []string{
	"Berserker",
	"Cheat",
	"Commander",
	"Hunter",
	"Mage",
	"Oathsworn",
	"Shadowmancer",
	"Shepherd",
	"Songweaver",
	"Stormshifter",
	"Zephyr",
}

func Class() string {
	return choose(classes)
}
