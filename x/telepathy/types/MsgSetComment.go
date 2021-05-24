package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetComment{}

type MsgSetComment struct {
	ID        string         `json:"id" yaml:"id"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Message   string         `json:"message" yaml:"message"`
	ThoughtId string         `json:"thoughtId" yaml:"thoughtId"`
}

func NewMsgSetComment(creator sdk.AccAddress, id string, message string, thoughtId string) MsgSetComment {
	return MsgSetComment{
		ID:        id,
		Creator:   creator,
		Message:   message,
		ThoughtId: thoughtId,
	}
}

func (msg MsgSetComment) Route() string {
	return RouterKey
}

func (msg MsgSetComment) Type() string {
	return "SetComment"
}

func (msg MsgSetComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgSetComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}

	//checking message length, cant be more than 250 chars
	if len(msg.Message) > 250 {
		return sdkerrors.Wrap(ErrMsgTooLong, "Message too long.")
	}

	return nil
}
