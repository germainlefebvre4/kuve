package cmd

import (
	"fmt"

	"github.com/germainlefebvre4/kuve/internal/version"
	"github.com/germainlefebvre4/kuve/pkg/config"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List kubectl versions",
	Long:  `List available kubectl versions (remote or installed).`,
}

var listRemoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "List available remote kubectl versions",
	Long:  `List all available kubectl versions that can be downloaded.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		manager := version.NewManager(cfg)

		// Get stable version (simplified implementation)
		stableVersion, err := manager.GetStableVersion()
		if err != nil {
			return fmt.Errorf("failed to get stable version: %w", err)
		}

		fmt.Println("Latest stable version:")
		fmt.Printf("  %s\n", stableVersion)
		fmt.Println("\nNote: For a full list of versions, visit https://github.com/kubernetes/kubernetes/releases")

		return nil
	},
}

var listInstalledCmd = &cobra.Command{
	Use:   "installed",
	Short: "List installed kubectl versions",
	Long:  `List all kubectl versions installed on this system.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		manager := version.NewManager(cfg)
		versions, err := manager.ListInstalledVersions()
		if err != nil {
			return fmt.Errorf("failed to list installed versions: %w", err)
		}

		if len(versions) == 0 {
			fmt.Println("No kubectl versions installed.")
			fmt.Println("Use 'kuve install <version>' to install a version.")
			return nil
		}

		// Get current version
		currentVersion, _ := manager.GetCurrentVersion()

		fmt.Println("Installed kubectl versions:")
		for _, v := range versions {
			marker := " "
			if v == currentVersion {
				marker = "*"
			}
			fmt.Printf("%s %s\n", marker, v)
		}

		if currentVersion != "" {
			fmt.Printf("\n* = current version (%s)\n", currentVersion)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listRemoteCmd)
	listCmd.AddCommand(listInstalledCmd)
}
