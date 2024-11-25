package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "habbitchain/testutil/keeper"
	"habbitchain/x/profile/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ProfileKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
