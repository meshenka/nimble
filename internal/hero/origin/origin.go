package origin

import (
	"context"

	"github.com/meshenka/nimble/internal"
)

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

	// some exta for good measure
	"a Ragged Outpost",
	"an Isolated Hamlet",
	"a Bustling Port",
	"a Remote Village",
	"a Trading Hub",
	"a Sacred Site",
	"a Forgotten Ruin",
	"a Nomadic Camp",
	"a Frozen Wastes",
	"The Mountain Hall",
	"a Coastal Settlement",
	"a Forest Refuge",
	"a Borderland Fort",
	"a Fallen Stronghold",
	"a River Town",
	"a Hidden Sanctum",
	"an Ancient Grove",
	"a Cursed Lands",
	"a Windswept Cliff",
	"a Sunken Hold",
}

func Select(ctx context.Context) string {
	return internal.Choose(ctx, origins)
}

func All() []string {
	return origins
}
