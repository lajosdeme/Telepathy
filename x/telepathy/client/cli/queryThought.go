package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
	"github.com/spf13/cobra"
)

func GetCmdListThought(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-thought",
		Short: "list all thought",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListThought, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Thought\n%s\n", err.Error())
				return nil
			}
			var out []types.Thought
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetThought(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-thought [key]",
		Short: "Query a thought by key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetThought, key), nil)
			if err != nil {
				fmt.Printf("could not resolve thought %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Thought
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

//List thoughts by creator cmd
func GetCmdListThoughtByCreator(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-thoughts [id]",
		Short: "List all thoughts by creator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListThoughtByCreator, id), nil)
			if err != nil {
				fmt.Printf("could not list Thought\n%s\n", err.Error())
				return nil
			}
			var out []types.Thought
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
