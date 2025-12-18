package config

import (
	"os"
	"path/filepath"
)

const (
	// AppName is the name of the application
	AppName = "kuve"

	// VersionFileName is the name of the version file
	VersionFileName = ".kubernetes-version"

	// KubectlBinaryName is the name of the kubectl binary
	KubectlBinaryName = "kubectl"
)

// Config holds the application configuration
type Config struct {
	HomeDir        string
	KuveDir        string
	BinDir         string
	VersionsDir    string
	CurrentSymlink string
}

// New creates a new configuration
func New() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	kuveDir := filepath.Join(homeDir, "."+AppName)
	binDir := filepath.Join(kuveDir, "bin")
	versionsDir := filepath.Join(kuveDir, "versions")
	currentSymlink := filepath.Join(binDir, KubectlBinaryName)

	return &Config{
		HomeDir:        homeDir,
		KuveDir:        kuveDir,
		BinDir:         binDir,
		VersionsDir:    versionsDir,
		CurrentSymlink: currentSymlink,
	}, nil
}

// EnsureDirectories creates necessary directories if they don't exist
func (c *Config) EnsureDirectories() error {
	dirs := []string{c.KuveDir, c.BinDir, c.VersionsDir}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
