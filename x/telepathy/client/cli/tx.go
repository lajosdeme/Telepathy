package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	telepathyTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	telepathyTxCmd.AddCommand(flags.PostCommands(
		// this line is used by starport scaffolding # 1
		GetCmdCreateUser(cdc),
		GetCmdSetUser(cdc),
		GetCmdDeleteUser(cdc),
		GetCmdCreateComment(cdc),
		GetCmdSetComment(cdc),
		GetCmdDeleteComment(cdc),
		GetCmdCreateThought(cdc),
		GetCmdSetThought(cdc),
		GetCmdDeleteThought(cdc),
		GetCmdLikeThought(cdc),
		GetCmdLikeComment(cdc),
		GetCmdDislikeThought(cdc),
		GetCmdDislikeComment(cdc),
		GetCmdFollowUser(cdc),
		GetCmdUnfollowUser(cdc),
		GetCmdSetAvatar(cdc),
	)...)

	return telepathyTxCmd
}
