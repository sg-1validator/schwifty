package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/toschdev/schwifty/testutil/keeper"
	"github.com/toschdev/schwifty/testutil/nullify"
	"github.com/toschdev/schwifty/x/schwifty/keeper"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func createNNft(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Nft {
	items := make([]types.Nft, n)
	for i := range items {
		items[i].Id = keeper.AppendNft(ctx, items[i])
	}
	return items
}

func TestNftGet(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetNft(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestNftRemove(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNft(ctx, item.Id)
		_, found := keeper.GetNft(ctx, item.Id)
		require.False(t, found)
	}
}

func TestNftGetAll(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNNft(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNft(ctx)),
	)
}

func TestNftCount(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNNft(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetNftCount(ctx))
}
