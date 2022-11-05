package types

import (
	exchangeTypes "bu-chain/x/exchange/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type ExchangeKeeper interface {
	SetExchangeRate(sdk.Context, exchangeTypes.ExchangeRate)
	GetExchangeRate(sdk.Context, string) (val exchangeTypes.ExchangeRate, found bool)
}
