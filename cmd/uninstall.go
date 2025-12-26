package cmd

import (
	"fmt"

	"github.com/germainlefebvre4/kuve/internal/kubectl"
	"github.com/germainlefebvre4/kuve/pkg/config"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall <version>",
	Short: "Uninstall a specific kubectl version",
	Long: `Remove a specific kubectl version from your system.

Example:
  kuve uninstall v1.28.0
  kuve uninstall 1.28.0`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version := args[0]

		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		installer := kubectl.NewInstaller(cfg)
		if err := installer.Uninstall(version); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
