package keeper

import (
	"bu-chain/x/exchange/types"
)

var _ types.QueryServer = Keeper{}
