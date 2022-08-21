package cli

import (
	"strconv"

	"parrots/x/parrots/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-comment [name] [display-name] [comment] [beak-id]",
		Short: "Broadcast message createComment",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDisplayName := args[1]
			argComment := args[2]
			argBeakId, err := strconv.ParseUint(args[3], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateComment(
				clientCtx.GetFromAddress().String(),
				argName,
				argDisplayName,
				argComment,
				argBeakId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
