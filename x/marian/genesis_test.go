package marian_test

import (
	"testing"

	keepertest "github.com/meczero/mariannet/testutil/keeper"
	"github.com/meczero/mariannet/testutil/nullify"
	"github.com/meczero/mariannet/x/marian"
	"github.com/meczero/mariannet/x/marian/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MarianKeeper(t)
	marian.InitGenesis(ctx, *k, genesisState)
	got := marian.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
