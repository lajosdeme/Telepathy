package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgSetUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetUser) (*sdk.Result, error) {
	if !msg.Creator.Equals(k.GetUserOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetUser(ctx, msg.ID, msg.Username, msg.Bio)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
