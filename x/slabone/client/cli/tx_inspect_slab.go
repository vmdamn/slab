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

func CmdInspectSlab() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inspect-slab [vetter-social-id] [vetting-note] [uri-vetter] [id]",
		Short: "Broadcast message inspect-slab",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVetterSocialId := args[0]
			argVettingNote := args[1]
			argUriVetter := args[2]
			argId, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInspectSlab(
				clientCtx.GetFromAddress().String(),
				argVetterSocialId,
				argVettingNote,
				argUriVetter,
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
