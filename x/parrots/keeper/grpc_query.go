package keeper

import (
	"parrots/x/parrots/types"
)

var _ types.QueryServer = Keeper{}
