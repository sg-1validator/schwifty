package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/toschdev/schwifty/testutil/keeper"
	"github.com/toschdev/schwifty/x/schwifty/keeper"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SchwiftyKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
