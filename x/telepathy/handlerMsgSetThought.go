package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgSetThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetThought) (*sdk.Result, error) {
	var thought = types.Thought{
		Creator: msg.Creator,
		ID:      msg.ID,
		Message: msg.Message,
	}
	if !msg.Creator.Equals(k.GetThoughtOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Thought can only be edited by its owner") // If not, throw an error
	}

	k.SetThought(ctx, thought)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
