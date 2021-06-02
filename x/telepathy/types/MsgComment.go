package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

/* -------------------------------------------------------------------------- */
/* ----------------------- Create comment message --------------------------- */
/* -------------------------------------------------------------------------- */

type MsgCreateComment struct {
	Creator        sdk.AccAddress `json:"creator" yaml:"creator"`
	Message        string         `json:"message" yaml:"message"`
	ThoughtId      string         `json:"thoughtId" yaml:"thoughtId"`
	OwnerCommentId string         `json:"ownerCommentId" yaml:"ownerCommentId"`
}

func NewMsgCreateComment(creator sdk.AccAddress, message string, thoughtId string, ownerCommentId string) MsgCreateComment {
	return MsgCreateComment{
		Creator:        creator,
		Message:        message,
		ThoughtId:      thoughtId,
		OwnerCommentId: ownerCommentId,
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

/* -------------------------------------------------------------------------- */
/* ------------------------- Set comment message ---------------------------- */
/* -------------------------------------------------------------------------- */

type MsgSetComment struct {
	ID             string         `json:"id" yaml:"id"`
	Creator        sdk.AccAddress `json:"creator" yaml:"creator"`
	Message        string         `json:"message" yaml:"message"`
	ThoughtId      string         `json:"thoughtId" yaml:"thoughtId"`
	OwnerCommentId string         `json:"ownerCommentId" yaml:"ownerCommentId"`
}

func NewMsgSetComment(creator sdk.AccAddress, id string, message string, thoughtId string, ownerCommentId string) MsgSetComment {
	return MsgSetComment{
		ID:             id,
		Creator:        creator,
		Message:        message,
		ThoughtId:      thoughtId,
		OwnerCommentId: ownerCommentId,
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

/* -------------------------------------------------------------------------- */
/* ----------------------- Delete comment message --------------------------- */
/* -------------------------------------------------------------------------- */

type MsgDeleteComment struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

func NewMsgDeleteComment(id string, creator sdk.AccAddress) MsgDeleteComment {
	return MsgDeleteComment{
		ID:      id,
		Creator: creator,
	}
}

func (msg MsgDeleteComment) Route() string {
	return RouterKey
}

func (msg MsgDeleteComment) Type() string {
	return "DeleteComment"
}

func (msg MsgDeleteComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

/* -------------------------------------------------------------------------- */
/* ------------------------- Like comment message --------------------------- */
/* -------------------------------------------------------------------------- */

type MsgLikeComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	CommentId string         `json:"commentId" yaml:"commentId"`
}

func NewMsgLikeComment(creator sdk.AccAddress, commentId string) MsgLikeComment {
	return MsgLikeComment{
		Creator:   creator,
		CommentId: commentId,
	}
}

func (msg MsgLikeComment) Route() string {
	return RouterKey
}

func (msg MsgLikeComment) Type() string {
	return "LikeComment"
}

func (msg MsgLikeComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgLikeComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgLikeComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}

/* -------------------------------------------------------------------------- */
/* ------------------------ Dislike comment message ------------------------- */
/* -------------------------------------------------------------------------- */

type MsgDislikeComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ThoughtId string         `json:"commentId" yaml:"commentId"`
}

func NewMsgDislikeComment(creator sdk.AccAddress, thoughtId string) MsgDislikeComment {
	return MsgDislikeComment{
		Creator:   creator,
		ThoughtId: thoughtId,
	}
}

func (msg MsgDislikeComment) Route() string {
	return RouterKey
}

func (msg MsgDislikeComment) Type() string {
	return "DislikeComment"
}

func (msg MsgDislikeComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDislikeComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgDislikeComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
