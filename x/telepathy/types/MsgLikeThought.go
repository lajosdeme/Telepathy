package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgLikeThought struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ThoughtId string         `json:"thoughtId" yaml:"thoughtId"`
}

func NewMsgLikeThought(creator sdk.AccAddress, thoughtId string) MsgLikeThought {
	return MsgLikeThought{
		Creator:   creator,
		ThoughtId: thoughtId,
	}
}

func (msg MsgLikeThought) Route() string {
	return RouterKey
}

func (msg MsgLikeThought) Type() string {
	return "LikeThought"
}

func (msg MsgLikeThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgLikeThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgLikeThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
