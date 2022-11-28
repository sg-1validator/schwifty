package schwifty

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/toschdev/schwifty/x/schwifty/keeper"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the collection
	for _, elem := range genState.CollectionList {
		k.SetCollection(ctx, elem)
	}

	// Set collection count
	k.SetCollectionCount(ctx, genState.CollectionCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CollectionList = k.GetAllCollection(ctx)
	genesis.CollectionCount = k.GetCollectionCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
