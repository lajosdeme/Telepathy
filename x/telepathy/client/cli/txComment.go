package cli

import (
	"bufio"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/lajosdeme/telepathy/x/telepathy/types"
)

func GetCmdCreateComment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-comment [message] [thoughtId] [commentId]",
		Short: "Creates a new comment",
		Long:  "Can add a comment either to a thought or another comment. Pass an empty string for 'commentId' if adding to a thought and vice versa.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMessage := string(args[0])
			argsThoughtId := string(args[1])
			argsCommentId := string(args[2])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateComment(cliCtx.GetFromAddress(), string(argsMessage), string(argsThoughtId), argsCommentId)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetComment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-comment [id]  [message] [thoughtId] [commentId]",
		Short: "Set a new comment",
		Long:  "Can edit a comment either to a thought or another comment. Pass an empty string for 'commentId' if adding to a thought and vice versa.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsMessage := string(args[1])
			argsThoughtId := string(args[2])
			argsCommentId := string(args[3])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetComment(cliCtx.GetFromAddress(), id, string(argsMessage), string(argsThoughtId), argsCommentId)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteComment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-comment [id]",
		Short: "Delete a new comment by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteComment(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

//Like a comment cmd
func GetCmdLikeComment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "like-comment [id]",
		Short: "Like a comment by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgLikeComment(cliCtx.GetFromAddress(), args[0])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

//Dislike a comment cmd
func GetCmdDislikeComment(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "dislike-comment [id]",
		Short: "Dislikes a comment by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDislikeComment(cliCtx.GetFromAddress(), args[0])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
