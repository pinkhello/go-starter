package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-starter/config"
	"os"
)

var (
	Version = "1.0.0"

	rootCmd = &cobra.Command{
		Use:     "go-starter",
		Version: Version,
		Short:   "go-starter Management CLI",
		Run: func(cmd *cobra.Command, args []string) {
			httpCmd.Run(cmd, args)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.AddCommand(httpCmd)
}
