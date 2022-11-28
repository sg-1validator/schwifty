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

func createNCollection(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Collection {
	items := make([]types.Collection, n)
	for i := range items {
		items[i].Id = keeper.AppendCollection(ctx, items[i])
	}
	return items
}

func TestCollectionGet(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNCollection(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCollection(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCollectionRemove(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNCollection(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCollection(ctx, item.Id)
		_, found := keeper.GetCollection(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCollectionGetAll(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNCollection(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCollection(ctx)),
	)
}

func TestCollectionCount(t *testing.T) {
	keeper, ctx := keepertest.SchwiftyKeeper(t)
	items := createNCollection(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCollectionCount(ctx))
}
