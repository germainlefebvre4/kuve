package kubectl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/germainlefebvre4/kuve/pkg/config"
)

// Installer handles kubectl installation
type Installer struct {
	config *config.Config
}

// NewInstaller creates a new kubectl installer
func NewInstaller(cfg *config.Config) *Installer {
	return &Installer{
		config: cfg,
	}
}

// Install downloads and installs a specific kubectl version
func (i *Installer) Install(version string) error {
	if version == "" {
		return fmt.Errorf("version cannot be empty")
	}

	// Normalize version (ensure it starts with 'v')
	if version[0] != 'v' {
		version = "v" + version
	}

	versionDir := filepath.Join(i.config.VersionsDir, version)
	kubectlPath := filepath.Join(versionDir, config.KubectlBinaryName)

	// Check if already installed
	if _, err := os.Stat(kubectlPath); err == nil {
		return fmt.Errorf("version %s is already installed", version)
	}

	// Create version directory
	if err := os.MkdirAll(versionDir, 0755); err != nil {
		return fmt.Errorf("failed to create version directory: %w", err)
	}

	// Build download URL
	downloadURL := fmt.Sprintf("https://dl.k8s.io/release/%s/bin/%s/%s/kubectl",
		version, runtime.GOOS, runtime.GOARCH)

	// Download kubectl binary
	fmt.Printf("Downloading kubectl %s for %s/%s...\n", version, runtime.GOOS, runtime.GOARCH)
	if err := i.downloadFile(downloadURL, kubectlPath); err != nil {
		os.RemoveAll(versionDir) // Cleanup on failure
		return fmt.Errorf("failed to download kubectl: %w", err)
	}

	// Make binary executable
	if err := os.Chmod(kubectlPath, 0755); err != nil {
		os.RemoveAll(versionDir) // Cleanup on failure
		return fmt.Errorf("failed to make kubectl executable: %w", err)
	}

	fmt.Printf("Successfully installed kubectl %s\n", version)
	return nil
}

// Uninstall removes a specific kubectl version
func (i *Installer) Uninstall(version string) error {
	if version == "" {
		return fmt.Errorf("version cannot be empty")
	}

	// Normalize version
	if version[0] != 'v' {
		version = "v" + version
	}

	versionDir := filepath.Join(i.config.VersionsDir, version)

	// Check if version is installed
	if _, err := os.Stat(versionDir); os.IsNotExist(err) {
		return fmt.Errorf("version %s is not installed", version)
	}

	// Check if this is the current version
	currentSymlink := i.config.CurrentSymlink
	if target, err := os.Readlink(currentSymlink); err == nil {
		if filepath.Dir(target) == versionDir {
			return fmt.Errorf("cannot uninstall %s as it is currently active. Switch to another version first", version)
		}
	}

	// Remove version directory
	if err := os.RemoveAll(versionDir); err != nil {
		return fmt.Errorf("failed to remove version directory: %w", err)
	}

	fmt.Printf("Successfully uninstalled kubectl %s\n", version)
	return nil
}

// Switch changes the active kubectl version
func (i *Installer) Switch(version string) error {
	if version == "" {
		return fmt.Errorf("version cannot be empty")
	}

	// Normalize version
	if version[0] != 'v' {
		version = "v" + version
	}

	versionDir := filepath.Join(i.config.VersionsDir, version)
	kubectlPath := filepath.Join(versionDir, config.KubectlBinaryName)

	// Check if version is installed
	if _, err := os.Stat(kubectlPath); os.IsNotExist(err) {
		return fmt.Errorf("version %s is not installed. Run 'kuve install %s' first", version, version)
	}

	// Remove existing symlink if it exists
	currentSymlink := i.config.CurrentSymlink
	if _, err := os.Lstat(currentSymlink); err == nil {
		if err := os.Remove(currentSymlink); err != nil {
			return fmt.Errorf("failed to remove existing symlink: %w", err)
		}
	}

	// Create new symlink
	if err := os.Symlink(kubectlPath, currentSymlink); err != nil {
		return fmt.Errorf("failed to create symlink: %w", err)
	}

	fmt.Printf("Switched to kubectl %s\n", version)
	fmt.Printf("Note: Make sure %s is in your PATH\n", i.config.BinDir)
	return nil
}

// downloadFile downloads a file from a URL and saves it to destPath
func (i *Installer) downloadFile(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download: HTTP %d", resp.StatusCode)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
