package cmd

import (
	"fmt"
	"os"

	"github.com/germainlefebvre4/kuve/internal/kubectl"
	"github.com/germainlefebvre4/kuve/internal/version"
	"github.com/germainlefebvre4/kuve/pkg/config"
	"github.com/spf13/cobra"
)

var (
	fromCluster bool
)

var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Use kubectl version from .kubernetes-version file or cluster",
	Long: `Switch to the kubectl version specified in the .kubernetes-version file or detect from cluster.
	
This command searches for a .kubernetes-version file in the current directory
and parent directories. If found, it switches to that version.

With --from-cluster flag, it detects the Kubernetes version from the current
cluster context and switches to the matching kubectl version.

If the version is not installed, it will be installed automatically.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.New()
		if err != nil {
			return fmt.Errorf("failed to create config: %w", err)
		}

		// Ensure directories exist
		if err := cfg.EnsureDirectories(); err != nil {
			return fmt.Errorf("failed to create directories: %w", err)
		}

		manager := version.NewManager(cfg)
		installer := kubectl.NewInstaller(cfg)

		var requestedVersion string

		if fromCluster {
			// Detect version from cluster
			fmt.Println("Detecting Kubernetes version from current cluster context...")
			rawVersion, normalizedVersion, err := manager.DetectClusterVersionWithRaw()
			if err != nil {
				return fmt.Errorf("failed to detect cluster version: %w", err)
			}
			if rawVersion != normalizedVersion {
				fmt.Printf("Detected cluster version: %s (using kubectl %s)\n", rawVersion, normalizedVersion)
			} else {
				fmt.Printf("Detected cluster version: %s\n", rawVersion)
			}
			requestedVersion = normalizedVersion
		} else {
			// Find .kubernetes-version file
			requestedVersion, err = version.FindVersionFile()
			if err != nil {
				return fmt.Errorf("error searching for version file: %w", err)
			}

			if requestedVersion == "" {
				return fmt.Errorf("no .kubernetes-version file found in current or parent directories")
			}
			fmt.Printf("Found version %s in .kubernetes-version file\n", requestedVersion)
		}

		// Normalize version
		if len(requestedVersion) > 0 && requestedVersion[0] != 'v' {
			requestedVersion = "v" + requestedVersion
		}

		// Check if version is installed
		if !manager.IsVersionInstalled(requestedVersion) {
			fmt.Printf("Version %s is not installed. Installing...\n", requestedVersion)
			if err := installer.Install(requestedVersion); err != nil {
				return fmt.Errorf("failed to install version: %w", err)
			}
		}

		// Switch to the version
		if err := installer.Switch(requestedVersion); err != nil {
			return fmt.Errorf("failed to switch version: %w", err)
		}

		return nil
	},
}

var initCmd = &cobra.Command{
	Use:   "init [version]",
	Short: "Create a .kubernetes-version file",
	Long: `Create a .kubernetes-version file in the current directory.
	
If no version is specified, the current active version will be used.`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var versionToWrite string

		if len(args) > 0 {
			versionToWrite = args[0]
		} else {
			// Use current version
			cfg, err := config.New()
			if err != nil {
				return fmt.Errorf("failed to create config: %w", err)
			}

			manager := version.NewManager(cfg)
			currentVersion, err := manager.GetCurrentVersion()
			if err != nil {
				return fmt.Errorf("no current version set and no version specified: %w", err)
			}
			versionToWrite = currentVersion
		}

		// Normalize version
		if versionToWrite[0] != 'v' {
			versionToWrite = "v" + versionToWrite
		}

		// Write version file
		versionFile := config.VersionFileName
		if err := os.WriteFile(versionFile, []byte(versionToWrite+"\n"), 0644); err != nil {
			return fmt.Errorf("failed to write version file: %w", err)
		}

		fmt.Printf("Created %s with version %s\n", versionFile, versionToWrite)
		return nil
	},
}

func init() {
	useCmd.Flags().BoolVarP(&fromCluster, "from-cluster", "c", false, "detect and use version from current Kubernetes cluster")
	rootCmd.AddCommand(useCmd)
	rootCmd.AddCommand(initCmd)
}
