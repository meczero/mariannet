package mariannet_test

import (
	"testing"

	keepertest "github.com/meczero/mariannet/testutil/keeper"
	"github.com/meczero/mariannet/testutil/nullify"
	"github.com/meczero/mariannet/x/mariannet"
	"github.com/meczero/mariannet/x/mariannet/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MariannetKeeper(t)
	mariannet.InitGenesis(ctx, *k, genesisState)
	got := mariannet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
