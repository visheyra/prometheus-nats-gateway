package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile = ""

var rootCmd = &cobra.Command{
	Use:   "png",
	Short: "png - prometheu nats gateway",
	Long:  "tool that listen prometheus 2.0 events translate them to json, then publish them to nats",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "config file, there is no default")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
}

//Execute the rootCmd to forward args
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
