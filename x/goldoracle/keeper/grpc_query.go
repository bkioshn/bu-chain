package keeper

import (
	"bu-chain/x/goldoracle/types"
)

var _ types.QueryServer = Keeper{}
