package keeper

import (
	"context"
	"errors"
	"math"
	"strconv"
	"strings"

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

func (k Keeper) ExchangeAmount(goCtx context.Context, req *types.QueryExchangeAmountRequest) (*types.QueryExchangeAmountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	exchangePair := req.Denom + "-" + req.ExchangeToken
	rate, isFound := k.GetExchangeRate(ctx, exchangePair)
	exRate, err := strconv.ParseFloat(rate.Rate, 64)

	if strings.Compare(req.Denom, req.ExchangeToken) == 1 {
		exchangePair = req.ExchangeToken + "-" + req.Denom
		rate, isFound = k.GetExchangeRate(ctx, exchangePair)
		if !isFound {
			return nil, errors.New("Token pair rate not found")
		}
		exRate, err = strconv.ParseFloat(rate.Rate, 64)
		if err != nil {
			return nil, err
		}
		exRate = 1 / exRate
	}
	if !isFound {
		return nil, errors.New("Token pair rate not found")
	}
	if err != nil {
		return nil, err
	}
	exRate64 := uint64(exRate * math.Pow10(3))
	tokenReceivedAmount, err := strconv.ParseUint(req.Amount, 10, 64)
	tokenReceivedAmount = tokenReceivedAmount * exRate64 / uint64(math.Pow10(3))
	return &types.QueryExchangeAmountResponse{Amount: tokenReceivedAmount}, nil
}
