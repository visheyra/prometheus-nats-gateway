package cmd

import (
	"github.com/spf13/cobra"
	"github.com/visheyra/prometheus-nats-gateway/prom"
	"go.uber.org/zap"
)

var forward = ""
var listen = ""
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the tool",
	Long:  "Start png with specified configuration",
	Run: func(cmd *cobra.Command, args []string) {

		l, err := zap.NewProduction()
		if err != nil {
			return
		}
		defer l.Sync()
		logger := l.Sugar()

		logger.Infow("Starting server",
			"listen", listen,
			"forward", forward,
		)

		prom.StartServer(listen, forward)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", ":8080", "listen address of the prometheus receiver endpoint")
	rootCmd.PersistentFlags().StringVarP(&forward, "forward", "f", "http://localhost:4222", "address of the remote nats endpoint")
}
