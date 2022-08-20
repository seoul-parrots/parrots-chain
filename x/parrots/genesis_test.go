package parrots_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "parrots/testutil/keeper"
	"parrots/testutil/nullify"
	"parrots/x/parrots"
	"parrots/x/parrots/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ParrotsKeeper(t)
	parrots.InitGenesis(ctx, *k, genesisState)
	got := parrots.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
