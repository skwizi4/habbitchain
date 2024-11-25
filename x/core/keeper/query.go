package keeper

import (
	"habbitchain/x/core/types"
)

var _ types.QueryServer = Keeper{}
