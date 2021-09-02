package cmd

import (
	"fmt"

	"github.com/polarbit/protoactor-go-sample/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "client",
	Short: "Actor client ",
	Long:  `Starts a client node in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Client ...")

		cluster.StartClient()
	},
}
