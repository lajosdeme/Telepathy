package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/lajosdeme/telepathy/x/telepathy/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for telepathy clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListUser:
			return listUser(ctx, k)
		case types.QueryGetUser:
			return getUser(ctx, path[1:], k)
		case types.QueryListComment:
			return listComment(ctx, k)
		case types.QueryGetComment:
			return getComment(ctx, path[1:], k)
		case types.QueryListCommentForThought:
			return listCommentForThought(ctx, path[1:], k)
		case types.QueryListCommentForComment:
			return listCommentForComment(ctx, path[1:], k)
		case types.QueryListThought:
			return listThought(ctx, k)
		case types.QueryGetThought:
			return getThought(ctx, path[1:], k)
		case types.QueryListThoughtByCreator:
			return listThoughtByCreator(ctx, path[1:], k)
		case types.QueryGetCompleteProfile:
			return getCompleteProfile(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown telepathy query endpoint")
		}
	}
}
