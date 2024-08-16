package main

import (
	"os"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := newRootCmd()
	if err := rootCmd.Execute(); err != nil {
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "bnbtool",
	}

	rootCmd.AddCommand(
		cmdGenHeaders(),
		cmdGenProof(),
	)

	rootCmd.PersistentFlags().String("node", "https://bsc-testnet-dataseed.bnbchain.org", "bsc node url")
	return rootCmd
}

func cmdGenHeaders() *cobra.Command {
	return &cobra.Command{
		Use:   "gen-headers [from-height] [to-height] [output-file]",
		Short: "generate bnb light client header",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			url, err := cmd.Flags().GetString("node")
			if err != nil {
				return err
			}

			c, err := newClient(cmd.Context(), url)
			if err != nil {
				return err
			}
			return c.genHeaders(cmd.Context(), cast.ToInt64(args[0]), cast.ToInt64(args[1]), args[2])
		},
	}
}

func cmdGenProof() *cobra.Command {
	return &cobra.Command{
		Use:   "gen-proof [height] [txHash] [receipt-file] [proof-file]",
		Short: "generate receipt proof",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			url, err := cmd.Flags().GetString("node")
			if err != nil {
				return err
			}

			c, err := newClient(cmd.Context(), url)
			if err != nil {
				return err
			}
			return c.genReceiptProof(cmd.Context(), cast.ToInt64(args[0]), args[1], args[2], args[3])
		},
	}
}
