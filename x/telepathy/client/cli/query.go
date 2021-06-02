package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group telepathy queries under a subcommand
	telepathyQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	telepathyQueryCmd.AddCommand(
		flags.GetCommands(
			// this line is used by starport scaffolding # 1
			GetCmdListUser(queryRoute, cdc),
			GetCmdGetUser(queryRoute, cdc),
			GetCmdListComment(queryRoute, cdc),
			GetCmdGetComment(queryRoute, cdc),
			GetCmdListThought(queryRoute, cdc),
			GetCmdGetThought(queryRoute, cdc),
			GetCmdListCommentsForThought(queryRoute, cdc),
			GetCmdListCommentsForComment(queryRoute, cdc),
			GetCmdListThoughtByCreator(queryRoute, cdc),
			GetCmdGetCompleteProfile(queryRoute, cdc),
			GetCmdGetAvatar(queryRoute, cdc),
		)...,
	)

	return telepathyQueryCmd
}
