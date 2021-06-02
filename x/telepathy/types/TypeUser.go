package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type User struct {
	Creator   sdk.AccAddress   `json:"creator" yaml:"creator"`
	ID        string           `json:"id" yaml:"id"`
	Username  string           `json:"username" yaml:"username"`
	Bio       string           `json:"bio" yaml:"bio"`
	Avatar    string           `json:"avatar" yaml:"avatar"`
	Following []sdk.AccAddress `json:"following" yaml:"following"`
	Followers []sdk.AccAddress `json:"followers" yaml:"followers"`
}

type CompleteProfile struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ID        string         `json:"id" yaml:"id"`
	Username  string         `json:"username" yaml:"username"`
	Bio       string         `json:"bio" yaml:"bio"`
	Avatar    string         `json:"avatar" yaml:"avatar"`
	Following []User         `json:"following" yaml:"following"`
	Followers []User         `json:"followers" yaml:"followers"`
	Thoughts  []Thought      `json:"thoughts" yaml:"thoughts"`
}
