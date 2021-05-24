package utils

import sdk "github.com/cosmos/cosmos-sdk/types"

func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func RemoveAddr(s []sdk.AccAddress, i int) []sdk.AccAddress {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Contains(l []sdk.AccAddress, like sdk.AccAddress) bool {
	for _, v := range l {
		if like.Equals(v) {
			return true
		}
	}
	return false
}
