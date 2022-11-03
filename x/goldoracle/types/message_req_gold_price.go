package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendReqGoldPrice = "send_req_gold_price"

var _ sdk.Msg = &MsgSendReqGoldPriceRequest{}

func NewMsgSendReqGoldPrice(
	creator string,
	port string,
	channelID string,
	timeout uint64,
) *MsgSendReqGoldPriceRequest {
	return &MsgSendReqGoldPriceRequest{
		Creator:   creator,
		Port:      port,
		ChannelID: channelID,
		Timeout:   timeout,
	}
}

func (msg *MsgSendReqGoldPriceRequest) Route() string {
	return RouterKey
}

func (msg *MsgSendReqGoldPriceRequest) Type() string {
	return TypeMsgSendReqGoldPrice
}

func (msg *MsgSendReqGoldPriceRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendReqGoldPriceRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendReqGoldPriceRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.Timeout == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
