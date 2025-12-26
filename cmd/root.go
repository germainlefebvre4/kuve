package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	appVersion = "dev"
)

var rootCmd = &cobra.Command{
	Use:   "kuve",
	Short: "Kubernetes Client Switcher",
	Long: `Kuve is a CLI tool to easily switch versions of kubectl.

It helps you manage multiple kubectl versions on your system,
allowing you to install, switch, and use different versions
based on your needs or project requirements.`,
	Version: appVersion,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flags can be added here
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
