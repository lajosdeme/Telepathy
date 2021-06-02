package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* -------------------------------------------------------------------------- */
/* ------------------------ Create thought message -------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* -------------------------- Set thought message --------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* ------------------------ Delete thought message -------------------------- */
/* -------------------------------------------------------------------------- */

type MsgDeleteThought struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteThought(id string, creator sdk.AccAddress) MsgDeleteThought {
	return MsgDeleteThought{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgDeleteThought) Route() string {
	return RouterKey
}

func (msg MsgDeleteThought) Type() string {
	return "DeleteThought"
}

func (msg MsgDeleteThought) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteThought) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteThought) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

/* -------------------------------------------------------------------------- */
/* ------------------------- Like thought message --------------------------- */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/* ------------------------ Dislike thought message ------------------------- */
/* -------------------------------------------------------------------------- */

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
