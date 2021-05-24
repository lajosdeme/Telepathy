package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgDislikeThought struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ThoughtId string         `json:"thoughtId" yaml:"thoughtId"`
}

func NewMsgDislikeThought(creator sdk.AccAddress, thoughtId string) MsgDislikeThought {
	return MsgDislikeThought{
		Creator:   creator,
		ThoughtId: thoughtId,
	}
}

func (msg MsgDislikeThought) Route() string {
	return RouterKey
}

func (msg MsgDislikeThought) Type() string {
	return "DislikeThought"
}

func (msg MsgDislikeThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDislikeThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDislikeThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
