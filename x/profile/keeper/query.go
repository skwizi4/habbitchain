package keeper

import (
	"habbitchain/x/profile/types"
)

var _ types.QueryServer = Keeper{}
