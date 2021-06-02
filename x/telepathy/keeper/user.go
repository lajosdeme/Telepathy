package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
	"github.com/lajosdeme/telepathy/x/telepathy/utils"
)

// GetUserCount get the total number of user
func (k Keeper) GetUserCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.UserCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetUserCount set the total number of user
func (k Keeper) SetUserCount(ctx sdk.Context, count int64) {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.UserCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateUser creates a user
func (k Keeper) CreateUser(ctx sdk.Context, msg types.MsgCreateUser) {
	// Create the user if it doesnt already exist for creator address
	if !k.UserExists(ctx, msg.Creator.String()) {
		count := k.GetUserCount(ctx)
		var user = types.User{
			Creator:  msg.Creator,
			ID:       msg.Creator.String(),
			Username: msg.Username,
			Bio:      msg.Bio,
		}

		store := ctx.KVStore(k.storeKey)
		key := []byte(types.UserPrefix + user.ID)
		value := k.cdc.MustMarshalBinaryLengthPrefixed(user)
		store.Set(key, value)

		// Update user count
		k.SetUserCount(ctx, count+1)
	}
}

// GetUser returns the user information
func (k Keeper) GetUser(ctx sdk.Context, key string) (types.User, error) {
	store := ctx.KVStore(k.storeKey)
	var user types.User
	byteKey := []byte(types.UserPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

// SetUsername sets a new username for the user
func (k Keeper) SetUser(ctx sdk.Context, id string, username string, bio string) (*sdk.Result, error) {

	user, err := k.GetUser(ctx, id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User with address can not be found.")
	}

	if user.Username != username && username != "" {
		user.Username = username
	}
	if user.Bio != bio {
		user.Bio = bio
	}

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(user)
	key := []byte(types.UserPrefix + id)
	store.Set(key, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

// DeleteUser deletes a user
func (k Keeper) DeleteUser(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.UserPrefix + key))
}

//Follow a user
func (k Keeper) FollowUser(ctx sdk.Context, userId string, creator sdk.AccAddress) (*sdk.Result, error) {
	user, err := k.GetUser(ctx, userId)
	creatorUser, err2 := k.GetUser(ctx, creator.String())

	if err != nil || err2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User with address can not be found.")
	}

	if creator.Equals(user.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User cant follow him/herself.")
	}

	if !utils.Contains(user.Followers, creator) {
		user.Followers = append(user.Followers, creator)
		creatorUser.Following = append(creatorUser.Following, user.Creator)

		store := ctx.KVStore(k.storeKey)
		bz := k.cdc.MustMarshalBinaryLengthPrefixed(user)
		key := []byte(types.UserPrefix + user.ID)
		store.Set(key, bz)

		bz2 := k.cdc.MustMarshalBinaryLengthPrefixed(creatorUser)
		key2 := []byte(types.UserPrefix + creatorUser.ID)
		store.Set(key2, bz2)

		return &sdk.Result{Events: ctx.EventManager().Events()}, nil
	}
	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "User already followed.")
}

//Unfollow a user
func (k Keeper) UnfollowUser(ctx sdk.Context, userId string, creator sdk.AccAddress) (*sdk.Result, error) {
	user, err := k.GetUser(ctx, userId)
	creatorUser, err2 := k.GetUser(ctx, creator.String())

	if err != nil || err2 != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User with address can not be found.")
	}

	if creator.Equals(user.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User cant follow him/herself.")
	}

	store := ctx.KVStore(k.storeKey)

	for i, f := range user.Followers {
		if f.Equals(creator) {
			user.Followers = utils.RemoveAddr(user.Followers, i)

			bz := k.cdc.MustMarshalBinaryLengthPrefixed(user)
			key := []byte(types.UserPrefix + user.ID)
			store.Set(key, bz)
		}
	}

	for i, f := range creatorUser.Following {
		if f.Equals(user.Creator) {
			creatorUser.Following = utils.RemoveAddr(creatorUser.Following, i)

			bz2 := k.cdc.MustMarshalBinaryLengthPrefixed(creatorUser)
			key2 := []byte(types.UserPrefix + creatorUser.ID)
			store.Set(key2, bz2)
		}
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func (k Keeper) SetAvatar(ctx sdk.Context, id string, avatar string) (*sdk.Result, error) {
	user, err := k.GetUser(ctx, id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "User with address can not be found.")
	}

	user.Avatar = avatar

	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(user)
	key := []byte(types.UserPrefix + id)
	store.Set(key, bz)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

//
// Functions used by querier
//
//List all users
func listUser(ctx sdk.Context, k Keeper) ([]byte, error) {
	var userList []types.User
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.UserPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var user types.User
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &user)
		userList = append(userList, user)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, userList)
	return res, nil
}

//Get user by id
func getUser(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	user, err := k.GetUser(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, user)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

//Get the complete profile of a user
func getCompleteProfile(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	id := path[0]
	user, err := k.GetUser(ctx, id)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "user with address not found")
	}

	thoughts := k.ListThoughtByCreator(ctx, user.Creator)

	var followers []types.User
	var followings []types.User

	for _, addr := range user.Followers {
		follower, err := k.GetUser(ctx, addr.String())
		if err == nil {
			followers = append(followers, follower)
		}
	}

	for _, addr := range user.Following {
		following, err := k.GetUser(ctx, addr.String())
		if err == nil {
			followings = append(followings, following)
		}
	}

	var completeProfile = types.CompleteProfile{
		Creator:   user.Creator,
		ID:        user.ID,
		Username:  user.Username,
		Bio:       user.Bio,
		Following: followings,
		Followers: followers,
		Thoughts:  thoughts,
	}

	res, err = codec.MarshalJSONIndent(k.cdc, completeProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetUserOwner(ctx sdk.Context, key string) sdk.AccAddress {
	user, err := k.GetUser(ctx, key)
	if err != nil {
		return nil
	}
	return user.Creator
}

// Check if the key exists in the store
func (k Keeper) UserExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.UserPrefix + key))
}
