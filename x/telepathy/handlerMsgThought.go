package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateThought) (*sdk.Result, error) {
	k.CreateThought(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSetThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetThought) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetThoughtOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Thought can only be edited by its owner") // If not, throw an error
	}

	k.SetThought(ctx, msg.ID, msg.Message)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// Handle a message to delete name
func handleMsgDeleteThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteThought) (*sdk.Result, error) {
	if !k.ThoughtExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetThoughtOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteThought(ctx, msg.ID)
	return &sdk.Result{}, nil
}

func handleMsgLikeThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgLikeThought) (*sdk.Result, error) {
	return k.LikeThought(ctx, msg.ThoughtId, msg.Creator)
}

func handleMsgDislikeThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgDislikeThought) (*sdk.Result, error) {
	return k.DislikeThought(ctx, msg.ThoughtId, msg.Creator)
}
