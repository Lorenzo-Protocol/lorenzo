package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	lrz "github.com/Lorenzo-Protocol/lorenzo/types"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/x/btcstaking/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateBTCStakingWithBTCProofCmd(),
		NewBurnCmd(),
	)

	return cmd
}

func NewCreateBTCStakingWithBTCProofCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-btcstaking-with-btc-proof [btc_tx_bytes] [proof] [receiver_name]",
		Short: "Create a new btc staking request with proof from bitcoin-cli getrawtransaction&gettxoutproof output",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txBytes, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("failed to decode tx bytes: %s", err)
			}
			proofRaw, err := hex.DecodeString(args[1])
			if err != nil {
				return fmt.Errorf("failed to decode proof bytes: %s", err)
			}
			merkleBlk, err := keeper.ParseMerkleBlock(proofRaw)
			if err != nil {
				return err
			}
			txIndex, proofBytes, err := keeper.ParseBTCProof(merkleBlk)

			queryClient := types.NewQueryClient(clientCtx)
			resp, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			if _, receiverExists := resp.Params.Receivers[args[2]]; !receiverExists {
				return fmt.Errorf("receiver(%s) not found", args[2])
			}

			blkHdr := &merkleBlk.Header

			var blkHdrHashBytes lrz.BTCHeaderHashBytes
			tmp := blkHdr.BlockHash()
			blkHdrHashBytes.FromChainhash(&tmp)

			msg := types.MsgCreateBTCStaking{
				Receiver: args[2],
				Signer:   clientCtx.GetFromAddress().String(),
				StakingTx: &types.TransactionInfo{
					Key: &types.TransactionKey{
						Index: txIndex,
						Hash:  &blkHdrHashBytes,
					},
					Transaction: txBytes,
					Proof:       proofBytes,
				},
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBurnCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [btc_target_address] [amouont]",
		Short: "burn tokens with btc target address and amount",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amount, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			msg := types.NewMsgBurnRequest(clientCtx.GetFromAddress().String(), args[0], uint64(amount))
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
