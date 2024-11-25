package habbitchain_test

import (
	"testing"

	keepertest "habbitchain/testutil/keeper"
	"habbitchain/testutil/nullify"
	habbitchain "habbitchain/x/habbitchain/module"
	"habbitchain/x/habbitchain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HabbitchainKeeper(t)
	habbitchain.InitGenesis(ctx, k, genesisState)
	got := habbitchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
