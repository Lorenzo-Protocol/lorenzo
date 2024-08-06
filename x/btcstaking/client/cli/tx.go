package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v2/types"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v2/x/btcstaking/types"
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
		Use:   "btcstaking [btc_tx_bytes] [proof] [agent_id]",
		Short: "Create a new btc staking request with proof from bitcoin-cli getrawtransaction&gettxoutproof output",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txBytes, err := hex.DecodeString(args[0])
			if err != nil {
				return fmt.Errorf("failed to decode tx bytes: %w", err)
			}
			agentId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse agent id(%s): %w", args[2], err)
			}
			proofRaw, err := hex.DecodeString(args[1])
			if err != nil {
				return fmt.Errorf("failed to decode proof bytes: %w", err)
			}
			merkleBlk, err := keeper.ParseMerkleBlock(proofRaw)
			if err != nil {
				return err
			}
			txIndex, proofBytes, err := keeper.ParseBTCProof(merkleBlk)
			if err != nil {
				return fmt.Errorf("failed to parse btc proof: %w", err)
			}

			blkHdr := &merkleBlk.Header

			var blkHdrHashBytes lrz.BTCHeaderHashBytes
			tmp := blkHdr.BlockHash()
			blkHdrHashBytes.FromChainhash(&tmp)

			msg := types.MsgCreateBTCStaking{
				AgentId:  agentId,
				Receiver: " ",
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
		Use:   "burn [btc_address] [amount]",
		Short: "burn stBTC tokens, accepting two parameters: the btc address as the recipient address for BTC and the amount to be burned",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			amount, ok := math.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("amount must be a valid integer")
			}
			msg := types.NewMsgBurnRequest(clientCtx.GetFromAddress().String(), args[0], amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
