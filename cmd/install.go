package cmd

import (
	"fmt"

	"github.com/germainlefebvre4/kuve/internal/kubectl"
	"github.com/germainlefebvre4/kuve/pkg/config"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install <version>",
	Short: "Install a specific kubectl version",
	Long: `Download and install a specific kubectl version.
	
Example:
  kuve install v1.28.0
  kuve install 1.28.0`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		version := args[0]

		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		// Ensure directories exist
		if err := cfg.EnsureDirectories(); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
		}

		installer := kubectl.NewInstaller(cfg)
		if err := installer.Install(version); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
