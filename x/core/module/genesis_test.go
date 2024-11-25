package core_test

import (
	"testing"

	keepertest "habbitchain/testutil/keeper"
	"habbitchain/testutil/nullify"
	core "habbitchain/x/core/module"
	"habbitchain/x/core/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	core.InitGenesis(ctx, k, genesisState)
	got := core.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
