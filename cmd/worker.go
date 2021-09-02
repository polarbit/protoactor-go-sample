package cmd

import (
	"errors"
	"fmt"

	"github.com/polarbit/protoactor-go-sample/cluster"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(func() *cobra.Command {
		var cPort int
		var sPort int
		var hosts []string

		c := &cobra.Command{
			Use:   "worker",
			Short: "Actor worker",
			Long:  `Starts a worker node for the cluster`,
			RunE: func(cmd *cobra.Command, args []string) error {
				if cPort < 1000 {
					return errors.New("cluster port number should be given and greater than 1000")
				}

				if sPort < 1000 {
					return errors.New("service port number should be given and greater than 1000")
				}

				fmt.Printf("Starting worker at cluster-port %v, service-port %v", cPort, sPort)
				cluster.StartWorker(cPort, sPort, hosts)

				return nil
			},
		}

		c.Flags().IntVarP(&cPort, "cluster-port", "c", 0, "cluster port number, E.g. 6331")
		c.Flags().IntVarP(&sPort, "service-port", "s", 0, "service (or grpc) port number, E.g. 8080")
		c.Flags().StringSliceVarP(&hosts, "host", "m", []string{}, "cluster memebers to connect, E.g. 127.0.0.1:6331")
		return c
	}())
}
