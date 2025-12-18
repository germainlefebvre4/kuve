package cmd

import (
	"fmt"

	"github.com/germainlefebvre4/kuve/internal/kubectl"
	"github.com/germainlefebvre4/kuve/internal/version"
	"github.com/germainlefebvre4/kuve/pkg/config"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <version>",
	Short: "Switch to a specific kubectl version",
	Long: `Switch to a specific kubectl version. The version must be installed first.
	
Example:
  kuve switch v1.28.0
  kuve switch 1.28.0`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version := args[0]

		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		installer := kubectl.NewInstaller(cfg)
		if err := installer.Switch(version); err != nil {
			return err
		}

		return nil
	},
}

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current kubectl version",
	Long:  `Display the currently active kubectl version managed by kuve.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		manager := version.NewManager(cfg)
		currentVersion, err := manager.GetCurrentVersion()
		if err != nil {
			return err
		}

		fmt.Printf("Current kubectl version: %s\n", currentVersion)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(currentCmd)
}
