package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"slabone/x/slabone/types"
)

var _ = strconv.Itoa(0)

func CmdRevokeSlab() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-slab [revoker-social-id] [revoking-note] [uri-revoker] [id]",
		Short: "Broadcast message revoke-slab",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argRevokerSocialId := args[0]
			argRevokingNote := args[1]
			argUriRevoker := args[2]
			argId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeSlab(
				clientCtx.GetFromAddress().String(),
				argRevokerSocialId,
				argRevokingNote,
				argUriRevoker,
				argId,
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
