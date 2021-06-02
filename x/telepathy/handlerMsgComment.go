package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateComment) (*sdk.Result, error) {
	return k.CreateComment(ctx, msg)
}

func handleMsgSetComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetComment) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetCommentOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetComment(ctx, msg.ID, msg.Message)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// Handle a message to delete comment
func handleMsgDeleteComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteComment) (*sdk.Result, error) {
	if !k.CommentExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetCommentOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteComment(ctx, msg.ID)
	return &sdk.Result{}, nil
}

func handleMsgLikeComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgLikeComment) (*sdk.Result, error) {
	return k.LikeComment(ctx, msg.CommentId, msg.Creator)
}

func handleMsgDislikeComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgDislikeComment) (*sdk.Result, error) {
	return k.DislikeComment(ctx, msg.ThoughtId, msg.Creator)
}
