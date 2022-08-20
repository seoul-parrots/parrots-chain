package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "parrots/testutil/keeper"
	"parrots/x/parrots/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ParrotsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
