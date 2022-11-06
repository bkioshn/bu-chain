package keeper

import (
	appparams "bu-chain/app/params"
	"context"
	"fmt"
	"strconv"

	"bu-chain/x/exchange/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExchangeRateAll(c context.Context, req *types.QueryAllExchangeRateRequest) (*types.QueryAllExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exchangeRates []types.ExchangeRate
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	exchangeRateStore := prefix.NewStore(store, types.KeyPrefix(types.ExchangeRateKeyPrefix))

	pageRes, err := query.Paginate(exchangeRateStore, req.Pagination, func(key []byte, value []byte) error {
		var exchangeRate types.ExchangeRate
		if err := k.cdc.Unmarshal(value, &exchangeRate); err != nil {
			return err
		}

		exchangeRates = append(exchangeRates, exchangeRate)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExchangeRateResponse{ExchangeRate: exchangeRates, Pagination: pageRes}, nil
}

func (k Keeper) ExchangeRate(c context.Context, req *types.QueryGetExchangeRateRequest) (*types.QueryGetExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExchangeRate(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExchangeRateResponse{ExchangeRate: val}, nil
}

func (k Keeper) ExchangeAmount(c context.Context, req *types.QueryExchangeAmountRequest) (*types.QueryExchangeAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	// TODO: Process the query

	exchangePair := req.Denom + "-" + req.ExchangeToken

	tokenReceivedAmount, err := strconv.ParseFloat(req.Amount, 64) // ParseUint(req.Amount, 10, 64)
	if err != nil {
		return nil, err
	}
	if req.Denom == appparams.DisplayDenom {
		exchangePair = req.ExchangeToken + "-" + req.Denom
		rate, isFound := k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, types.ErrTokenPairNotFound
		}
		tokenReceivedAmount = tokenReceivedAmount * float64(rate.Multiplier) / float64(rate.Rate)

	} else {
		rate, isFound := k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, types.ErrTokenPairNotFound
		}
		tokenReceivedAmount = tokenReceivedAmount * float64(rate.Rate) / float64(rate.Multiplier)
	}
	return &types.QueryExchangeAmountResponse{Amount: fmt.Sprint(tokenReceivedAmount)}, nil
}

func (k Keeper) ExchangePairs(c context.Context, req *types.QueryExchangePairsRequest) (*types.QueryExchangePairsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var exchangePair []string
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)
	exchangeRateStore := prefix.NewStore(store, types.KeyPrefix(types.ExchangeRateKeyPrefix))
	pageRes, err := query.Paginate(exchangeRateStore, req.Pagination, func(key []byte, value []byte) error {
		var exchangeRate types.ExchangeRate
		if err := k.cdc.Unmarshal(value, &exchangeRate); err != nil {
			return err
		}

		exchangePair = append(exchangePair, exchangeRate.Index)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryExchangePairsResponse{ExchangePair: exchangePair, Pagination: pageRes}, nil
}
