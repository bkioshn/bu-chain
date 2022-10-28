package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/exchange module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrTokenPairNotFound = sdkerrors.Register(ModuleName, 1, "Token pair rate not found")
	ErrTokenShouldBePair = sdkerrors.Register(ModuleName, 2, "Input tokens need to be in pair")
)
