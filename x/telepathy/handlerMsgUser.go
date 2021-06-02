package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUser) (*sdk.Result, error) {
	k.CreateUser(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSetUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetUser) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetUserOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetUser(ctx, msg.ID, msg.Username, msg.Bio)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// Handle a message to delete name
func handleMsgDeleteUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteUser) (*sdk.Result, error) {
	if !k.UserExists(ctx, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetUserOwner(ctx, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteUser(ctx, msg.ID)
	return &sdk.Result{}, nil
}

func handleMsgFollowUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgFollowUser) (*sdk.Result, error) {
	return k.FollowUser(ctx, msg.UserId, msg.Creator)
}

func handleMsgUnfollowUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgUnfollowUser) (*sdk.Result, error) {
	return k.UnfollowUser(ctx, msg.UserId, msg.Creator)
}

func handleMsgSetAvatar(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetAvatar) (*sdk.Result, error) {
	return k.SetAvatar(ctx, msg.ID, msg.Avatar)
}
