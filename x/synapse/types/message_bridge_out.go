package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBridgeOut = "bridge_out"

var _ sdk.Msg = &MsgBridgeOut{}

func NewMsgBridgeOut(creator string, data []*OutBridgeData) *MsgBridgeOut {
	return &MsgBridgeOut{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgBridgeOut) Route() string {
	return RouterKey
}

func (msg *MsgBridgeOut) Type() string {
	return TypeMsgBridgeOut
}

func (msg *MsgBridgeOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBridgeOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBridgeOut) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
