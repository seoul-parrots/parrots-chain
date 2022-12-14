package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"parrots/x/parrots/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group parrots queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdGetProfiles())
	cmd.AddCommand(CmdProfileCount())
	cmd.AddCommand(CmdGetProfile())
	cmd.AddCommand(CmdGetProfileByUsername())

	cmd.AddCommand(CmdGetBeaksCount())
	cmd.AddCommand(CmdGetAllBeaks())
	cmd.AddCommand(CmdGetBeakById())

	cmd.AddCommand(CmdGetBeaksByNameSubstring())

	cmd.AddCommand(CmdGetBeaksByTag())

	cmd.AddCommand(CmdGetRespectedBeaks())

	cmd.AddCommand(CmdGetProfileByCreator())

	cmd.AddCommand(CmdGetCommentsByBeakId())

	cmd.AddCommand(CmdGetFeeds())

	cmd.AddCommand(CmdGetCommentById())

	// this line is used by starport scaffolding # 1

	return cmd
}
