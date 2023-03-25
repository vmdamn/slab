package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"slabone/x/slabone/types"
)

var _ = strconv.Itoa(0)

func CmdCreateSlab() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-slab [originator-social-id] [directed-towards] [assertion] [uri-originator]",
		Short: "Broadcast message create-slab",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOriginatorSocialId := args[0]
			argDirectedTowards := args[1]
			argAssertion := args[2]
			argUriOriginator := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSlab(
				clientCtx.GetFromAddress().String(),
				argOriginatorSocialId,
				argDirectedTowards,
				argAssertion,
				argUriOriginator,
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
