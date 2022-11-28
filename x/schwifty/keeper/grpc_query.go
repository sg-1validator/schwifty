package keeper

import (
	"github.com/toschdev/schwifty/x/schwifty/types"
)

var _ types.QueryServer = Keeper{}
