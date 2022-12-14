package cli

import (
	"strconv"
	"strings"

	"parrots/x/parrots/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUploadBeak() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload-beak [file-index] [name] [creator-username] [creator-displayname] [description] [license] [linked-beaks] [tags]",
		Short: "Broadcast message uploadBeak",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFileIndex := args[0]
			argName := args[1]
			argCreatorUsername := args[2]
			argCreatorDisplayName := args[3]
			argDescription := args[4]
			argLicense := args[5]
			argLinkedBeaks := strings.Split(args[6], ",")
			argTags := strings.Split(args[7], ",")

			var argLinkedBeaksSlice []uint64
			for _, linkedBeak := range argLinkedBeaks {
				value, err := strconv.ParseUint(linkedBeak, 10, 64)
				if err != nil {
					return err
				}
				argLinkedBeaksSlice = append(argLinkedBeaksSlice, value)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUploadBeak(
				clientCtx.GetFromAddress().String(),
				argFileIndex,
				argName,
				argCreatorUsername,
				argCreatorDisplayName,
				argDescription,
				argLicense,
				argLinkedBeaksSlice,
				argTags,
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
