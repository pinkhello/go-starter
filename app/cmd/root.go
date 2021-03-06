package cmd

import (
	"fmt"
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

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show Version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(rootCmd.Version)
		},
	}

	projectCmd = &cobra.Command{
		Use:   "project",
		Short: "Show project name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(rootCmd.Use)
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
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(projectCmd)
	rootCmd.AddCommand(httpCmd)
}
