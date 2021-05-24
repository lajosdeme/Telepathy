package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgSetComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetComment) (*sdk.Result, error) {
	var comment = types.Comment{
		Creator:   msg.Creator,
		ID:        msg.ID,
		Message:   msg.Message,
		ThoughtId: msg.ThoughtId,
	}
	if !msg.Creator.Equals(k.GetCommentOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetComment(ctx, comment)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
