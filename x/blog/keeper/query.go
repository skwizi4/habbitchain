package keeper

import (
	"habbitchain/x/blog/types"
)

var _ types.QueryServer = Keeper{}
