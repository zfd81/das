package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

const (
	cliName        = "dasctl"
	cliDescription = "A simple command line client for DAS."

	defaultDialTimeout      = 2 * time.Second
	defaultCommandTimeOut   = 5 * time.Second
	defaultKeepAliveTime    = 2 * time.Second
	defaultKeepAliveTimeOut = 6 * time.Second
)

var rootCmd = &cobra.Command{
	Use:   cliName,
	Short: cliDescription,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DAS")
	},
}

func init() {
	rootCmd.AddCommand(NewVersionCommand())
	rootCmd.AddCommand(NewUserCommand())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ExitWithError(ExitError, err)
	}
}
