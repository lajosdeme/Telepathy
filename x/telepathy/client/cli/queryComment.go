package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
	"github.com/spf13/cobra"
)

func GetCmdListComment(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-comment",
		Short: "list all comment",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListComment, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Comment\n%s\n", err.Error())
				return nil
			}
			var out []types.Comment
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetComment(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-comment [key]",
		Short: "Query a comment by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetComment, key), nil)
			if err != nil {
				fmt.Printf("could not resolve comment %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Comment
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

//Get comments for a thought
func GetCmdListCommentsForThought(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-comments [id]",
		Short: "list comments for a thought",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			thoughtId := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListComment, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Comment\n%s\n", err.Error())
				return nil
			}
			var comments []types.Comment
			cdc.MustUnmarshalJSON(res, &comments)

			var out []types.Comment
			for _, c := range comments {
				if c.ThoughtId == thoughtId {
					out = append(out, c)
				}
			}
			return cliCtx.PrintOutput(out)
		},
	}
}
