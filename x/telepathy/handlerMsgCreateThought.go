package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateThought(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateThought) (*sdk.Result, error) {
	k.CreateThought(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
