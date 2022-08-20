package cli

import (
	"strconv"

	"parrots/x/parrots/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetBeaksByTag() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-beaks-by-tag [tag]",
		Short: "Query getBeaksByTag",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			tag := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetBeaksByTagRequest{Tag: tag}

			res, err := queryClient.GetBeaksByTag(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
