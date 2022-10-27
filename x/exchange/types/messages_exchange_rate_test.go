package types

import (
	"testing"

	"bu-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateExchangeRate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateExchangeRate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateExchangeRate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateExchangeRate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateExchangeRate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateExchangeRate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateExchangeRate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateExchangeRate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteExchangeRate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteExchangeRate
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteExchangeRate{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteExchangeRate{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
