package cli

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	lrz "github.com/Lorenzo-Protocol/lorenzo/v3/types"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/keeper"
	"github.com/Lorenzo-Protocol/lorenzo/v3/x/btcstaking/types"
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
		NewCreateBTCBStaking(),
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
				AgentId: agentId,
				Signer:  clientCtx.GetFromAddress().String(),
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

// NewCreateBTCBStaking creates a new cobra.Command for creating a new BTCB staking request with proof and receipt from the BNB chain.
//
// The command expects exactly 3 arguments: the BNB number, the receipt RLP file, and the proof RLP file.
//
// The function returns an error if there is an issue parsing the arguments or reading the files.
//
// The function sets up the command with the appropriate usage, short description, and arguments.
// It also sets up the command to run a function that generates and broadcasts a transaction using the provided arguments.
//
// The function adds transaction flags to the command and returns the command.
func NewCreateBTCBStaking() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xbtcstaking [chain-id] [number] [receipt-rlp-file] [proof-rlp-file]",
		Short: "Create a new xbtc staking request with proof and receipt from evm chain",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			chainID, err := strconv.ParseUint(args[0], 10, 32)
			if err != nil {
				return fmt.Errorf("failed to parse chainID(%s): %w", args[0], err)
			}

			number, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse number(%s): %w", args[0], err)
			}

			receiptRLP, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read receipt from file: %w", err)
			}

			proofRLP, err := os.ReadFile(args[3])
			if err != nil {
				return fmt.Errorf("failed to read proof from file: %w", err)
			}

			msg := types.MsgCreatexBTCStaking{
				ChainId: uint32(chainID),
				Signer:  clientCtx.GetFromAddress().String(),
				Number:  number,
				Receipt: receiptRLP,
				Proof:   proofRLP,
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
