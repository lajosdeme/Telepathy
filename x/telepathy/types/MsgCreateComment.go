package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateComment{}

type MsgCreateComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Message   string         `json:"message" yaml:"message"`
	ThoughtId string         `json:"thoughtId" yaml:"thoughtId"`
}

func NewMsgCreateComment(creator sdk.AccAddress, message string, thoughtId string) MsgCreateComment {
	return MsgCreateComment{
		Creator:   creator,
		Message:   message,
		ThoughtId: thoughtId,
	}
}

func (msg MsgCreateComment) Route() string {
	return RouterKey
}

func (msg MsgCreateComment) Type() string {
	return "CreateComment"
}

func (msg MsgCreateComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	//checking message length, cant be more than 250 chars
	if len(msg.Message) > 250 {
		return sdkerrors.Wrap(ErrMsgTooLong, "Message too long.")
	}

	return nil
}
