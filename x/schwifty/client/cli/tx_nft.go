package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/toschdev/schwifty/x/schwifty/types"
)

func CmdCreateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft [collection-id] [owner] [uri] [uri-hash] [data]",
		Short: "Create a new nft",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCollectionId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argOwner := args[1]
			argUri := args[2]
			argUriHash := args[3]
			argData := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateNft(clientCtx.GetFromAddress().String(), argCollectionId, argOwner, argUri, argUriHash, argData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-nft [id] [collection-id] [owner] [uri] [uri-hash] [data]",
		Short: "Update a nft",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argCollectionId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			argOwner := args[2]

			argUri := args[3]

			argUriHash := args[4]

			argData := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateNft(clientCtx.GetFromAddress().String(), id, argCollectionId, argOwner, argUri, argUriHash, argData)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-nft [id]",
		Short: "Delete a nft by id",
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

			msg := types.NewMsgDeleteNft(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
