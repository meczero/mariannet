package keeper_test

import (
	"testing"

	testkeeper "github.com/meczero/mariannet/testutil/keeper"
	"github.com/meczero/mariannet/x/mariannet/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MariannetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
