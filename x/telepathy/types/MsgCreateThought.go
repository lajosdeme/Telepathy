package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateThought{}

type MsgCreateThought struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Message string         `json:"message" yaml:"message"`
}

func NewMsgCreateThought(creator sdk.AccAddress, message string) MsgCreateThought {
	return MsgCreateThought{
		Creator: creator,
		Message: message,
	}
}

func (msg MsgCreateThought) Route() string {
	return RouterKey
}

func (msg MsgCreateThought) Type() string {
	return "CreateThought"
}

func (msg MsgCreateThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	//checking message length, cant be more than 250 chars
	if len(msg.Message) > 250 {
		return sdkerrors.Wrap(ErrMsgTooLong, "Message too long.")
	}

	return nil
}
