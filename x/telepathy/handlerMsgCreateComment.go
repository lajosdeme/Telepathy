package telepathy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lajosdeme/telepathy/x/telepathy/keeper"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func handleMsgCreateComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateComment) (*sdk.Result, error) {
	return k.CreateComment(ctx, msg)
}
