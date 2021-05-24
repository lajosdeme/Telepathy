package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Comment struct {
	Creator   sdk.AccAddress   `json:"creator" yaml:"creator"`
	ID        string           `json:"id" yaml:"id"`
	Message   string           `json:"message" yaml:"message"`
	ThoughtId string           `json:"thoughtId" yaml:"thoughtId"`
	Comments  []Comment        `json:"comments" yaml:"comments"`
	Likes     []sdk.AccAddress `json:"likes" yaml:"likes"`
	Shares    []sdk.AccAddress `json:"shares" yaml:"shares"`
}
