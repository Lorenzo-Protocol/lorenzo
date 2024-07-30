package cli

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/Lorenzo-Protocol/lorenzo/v2/x/plan/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for plan module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(GetCmdQueryParams())
	cmd.AddCommand(GetCmdQueryPlan())
	cmd.AddCommand(GetCmdQueryPlans())
	cmd.AddCommand(GetCmdQueryClaimLeafNode())
	return cmd
}

// GetCmdQueryParams returns a new Cobra command for showing the parameters of the plan module.
//
// No parameters.
// Returns *cobra.Command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "query the parameters of the plan module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryPlan returns a new Cobra command for querying a plan by id.
//
// Args: 1
//
//	0: plan id
//
// Returns *cobra.Command.
func GetCmdQueryPlan() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plan [id]",
		Short: "query a plan by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			planId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid plan ID: %s, error: %s", args[0], err.Error())
			}
			res, err := queryClient.Plan(context.Background(), &types.QueryPlanRequest{Id: planId})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// GetCmdQueryPlans returns a new Cobra command for querying all plans.
//
// No parameters.
// Returns *cobra.Command.
func GetCmdQueryPlans() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plans",
		Short: "query a plan",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			res, err := queryClient.Plans(context.Background(), &types.QueryPlansRequest{
				Pagination: pageReq,
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "plans")
	return cmd
}

func GetCmdQueryClaimLeafNode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-leaf-node [id] [round-id] [leaf-node]",
		Short: "claim a leaf node",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			queryClient := types.NewQueryClient(clientCtx)
			planId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid plan ID: %s, error: %s", args[0], err.Error())
			}

			roundId, ok := sdkmath.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid round ID: %s", args[1])
			}

			leafNode := common.HexToHash(args[2])
			if len(leafNode.Bytes()) != 32 {
				return fmt.Errorf("invalid merkle leaf node")
			}

			req := &types.QueryClaimLeafNodeRequest{
				LeafNode: args[2],
				Id:       planId,
				RoundId:  roundId,
			}
			res, err := queryClient.ClaimLeafNode(context.Background(), req)
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
