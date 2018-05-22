package cmd

import (
	"github.com/spf13/cobra"
	"github.com/visheyra/prometheus-nats-gateway/prom"
	"go.uber.org/zap"
)

var forward = ""
var listen = ""
var user = ""
var pass = ""
var topic = ""

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

		prom.StartServer(listen, forward, user, pass, topic)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	rootCmd.PersistentFlags().StringVarP(&listen, "listen", "l", ":8080", "listen address of the prometheus receiver endpoint")
	rootCmd.PersistentFlags().StringVarP(&forward, "forward", "f", "http://localhost:4222", "address of the remote nats endpoint")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "u", "user", "user to authenticate to nats endpoint")
	rootCmd.PersistentFlags().StringVarP(&pass, "password", "p", "password", "password for nats endpoint")
	rootCmd.PersistentFlags().StringVarP(&topic, "topic", "t", "default", "topic on which subscribe")
}
