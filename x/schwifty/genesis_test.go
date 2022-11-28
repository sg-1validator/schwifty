package schwifty_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/toschdev/schwifty/testutil/keeper"
	"github.com/toschdev/schwifty/testutil/nullify"
	"github.com/toschdev/schwifty/x/schwifty"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CollectionList: []types.Collection{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CollectionCount: 2,
		NftList: []types.Nft{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		NftCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SchwiftyKeeper(t)
	schwifty.InitGenesis(ctx, *k, genesisState)
	got := schwifty.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CollectionList, got.CollectionList)
	require.Equal(t, genesisState.CollectionCount, got.CollectionCount)
	require.ElementsMatch(t, genesisState.NftList, got.NftList)
	require.Equal(t, genesisState.NftCount, got.NftCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
