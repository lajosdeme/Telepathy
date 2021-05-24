package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetThought{}

type MsgSetThought struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Message string         `json:"message" yaml:"message"`
}

func NewMsgSetThought(creator sdk.AccAddress, id string, message string) MsgSetThought {
	return MsgSetThought{
		ID:      id,
		Creator: creator,
		Message: message,
	}
}

func (msg MsgSetThought) Route() string {
	return RouterKey
}

func (msg MsgSetThought) Type() string {
	return "SetThought"
}

func (msg MsgSetThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	//checking message length, cant be more than 250 chars
	if len(msg.Message) > 250 {
		return sdkerrors.Wrap(ErrMsgTooLong, "Message too long.")
	}

	return nil
}
