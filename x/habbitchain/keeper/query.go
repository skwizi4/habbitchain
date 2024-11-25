package keeper

import (
	"habbitchain/x/habbitchain/types"
)

var _ types.QueryServer = Keeper{}
