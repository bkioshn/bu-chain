package keeper

import (
	"bu-chain/x/exchange/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetExchangeRate set a specific exchangeRate in the store from its index
func (k Keeper) SetExchangeRate(ctx sdk.Context, exchangeRate types.ExchangeRate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKeyPrefix))
	b := k.cdc.MustMarshal(&exchangeRate)
	store.Set(types.ExchangeRateKey(
		exchangeRate.Index,
	), b)
}

// GetExchangeRate returns a exchangeRate from its index
func (k Keeper) GetExchangeRate(
	ctx sdk.Context,
	index string,

) (val types.ExchangeRate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKeyPrefix))

	b := store.Get(types.ExchangeRateKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExchangeRate removes a exchangeRate from the store
func (k Keeper) RemoveExchangeRate(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKeyPrefix))
	store.Delete(types.ExchangeRateKey(
		index,
	))
}

// GetAllExchangeRate returns all exchangeRate
func (k Keeper) GetAllExchangeRate(ctx sdk.Context) (list []types.ExchangeRate) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExchangeRateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExchangeRate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
