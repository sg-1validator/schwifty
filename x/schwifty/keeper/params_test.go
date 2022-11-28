package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/toschdev/schwifty/testutil/keeper"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SchwiftyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
