package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateUser) (*sdk.Result, error) {
	k.CreateUser(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
