package keeper

import (
	exchangetypes "bu-chain/x/exchange/types"
	"bu-chain/x/goldoracle/types"
	"errors"

	bandpacket "github.com/bandprotocol/bandchain-packet/packet"
	"github.com/bandprotocol/chain/v2/pkg/obi"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

func (k Keeper) TransmitOraclePacket(
	ctx sdk.Context,
	packetData bandpacket.OracleRequestPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {
	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes := packetData.GetBytes()

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)
	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}
	return nil
}

type Result struct {
	Rates []uint64
}

func (k Keeper) OnRecvOraclePacket(ctx sdk.Context, packet channeltypes.Packet, data bandpacket.OracleResponsePacketData) (packetAck bandpacket.OracleRequestPacketAcknowledgement, err error) {

	var result Result

	obi.MustDecode(data.Result, &result)
	exchangePair := "XAU-BUBU"
	curTime := uint64(ctx.BlockTime().Unix())
	//TODO - Creator should be something else
	var exchangeRate = exchangetypes.ExchangeRate{
		Creator:    data.ClientID,
		Index:      exchangePair,
		Rate:       result.Rates[0],
		Time:       curTime,
		Multiplier: 1e6,
	}

	rate, isFound := k.exchangeKeeper.GetExchangeRate(ctx, exchangePair)
	if !isFound {
		return packetAck, errors.New("not found")
	}

	if rate.Time < curTime {
		k.exchangeKeeper.SetExchangeRate(
			ctx,
			exchangeRate,
		)
	}
	return packetAck, nil
}

func (k Keeper) OnAcknowledgemenOraclePacket(ctx sdk.Context, packet channeltypes.Packet, data bandpacket.OracleRequestPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck bandpacket.OracleRequestPacketAcknowledgement

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

func (k Keeper) OnTimeoutOraclePacket(ctx sdk.Context, packet channeltypes.Packet, data bandpacket.OracleRequestPacketData) error {
	// TODO: packet timeout logic

	return nil
}
