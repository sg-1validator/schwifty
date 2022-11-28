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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SchwiftyKeeper(t)
	schwifty.InitGenesis(ctx, *k, genesisState)
	got := schwifty.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
