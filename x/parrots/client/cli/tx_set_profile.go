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

func CmdSetProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-profile [username] [display-name] [description]",
		Short: "Broadcast message setProfile",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUsername := args[0]
			argDisplayName := args[1]
			argDescription := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetProfile(
				clientCtx.GetFromAddress().String(),
				argUsername,
				argDisplayName,
				argDescription,
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
