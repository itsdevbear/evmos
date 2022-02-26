package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetKmsAddress = "set_kms_address"

var _ sdk.Msg = &MsgSetKmsAddress{}

func NewMsgSetKmsAddress(creator string) *MsgSetKmsAddress {
	return &MsgSetKmsAddress{
		Creator: creator,
	}
}

func (msg *MsgSetKmsAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetKmsAddress) Type() string {
	return TypeMsgSetKmsAddress
}

func (msg *MsgSetKmsAddress) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetKmsAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetKmsAddress) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
