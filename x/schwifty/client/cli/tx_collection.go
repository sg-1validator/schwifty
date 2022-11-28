package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func CmdCreateCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-collection [name] [description] [ticker] [uri] [uri-hash] [data]",
		Short: "Create a new collection",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argName := args[0]
			argDescription := args[1]
			argTicker := args[2]
			argUri := args[3]
			argUriHash := args[4]
			argData := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCollection(clientCtx.GetFromAddress().String(), clientCtx.GetFromAddress().String(), argName, argDescription, argTicker, argUri, argUriHash, argData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-collection [id] [owner] [name] [description] [ticker] [uri] [uri-hash] [data]",
		Short: "Update a collection",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argOwner := args[1]

			argName := args[2]

			argDescription := args[3]

			argTicker := args[4]

			argUri := args[5]

			argUriHash := args[6]

			argData := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCollection(clientCtx.GetFromAddress().String(), id, argOwner, argName, argDescription, argTicker, argUri, argUriHash, argData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteCollection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-collection [id]",
		Short: "Delete a collection by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteCollection(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
