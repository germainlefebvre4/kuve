package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNew(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Fatalf("New() failed: %v", err)
	}

	if cfg.HomeDir == "" {
		t.Error("HomeDir should not be empty")
	}

	if cfg.KuveDir == "" {
		t.Error("KuveDir should not be empty")
	}

	expectedKuveDir := filepath.Join(cfg.HomeDir, "."+AppName)
	if cfg.KuveDir != expectedKuveDir {
		t.Errorf("KuveDir = %s, want %s", cfg.KuveDir, expectedKuveDir)
	}

	expectedBinDir := filepath.Join(cfg.KuveDir, "bin")
	if cfg.BinDir != expectedBinDir {
		t.Errorf("BinDir = %s, want %s", cfg.BinDir, expectedBinDir)
	}

	expectedVersionsDir := filepath.Join(cfg.KuveDir, "versions")
	if cfg.VersionsDir != expectedVersionsDir {
		t.Errorf("VersionsDir = %s, want %s", cfg.VersionsDir, expectedVersionsDir)
	}
}

func TestEnsureDirectories(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "kuve-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	cfg := &Config{
		HomeDir:        tmpDir,
		KuveDir:        filepath.Join(tmpDir, ".kuve"),
		BinDir:         filepath.Join(tmpDir, ".kuve", "bin"),
		VersionsDir:    filepath.Join(tmpDir, ".kuve", "versions"),
		CurrentSymlink: filepath.Join(tmpDir, ".kuve", "bin", "kubectl"),
	}

	err = cfg.EnsureDirectories()
	if err != nil {
		t.Fatalf("EnsureDirectories() failed: %v", err)
	}

	// Check if directories were created
	dirs := []string{cfg.KuveDir, cfg.BinDir, cfg.VersionsDir}
	for _, dir := range dirs {
		info, err := os.Stat(dir)
		if err != nil {
			t.Errorf("Directory %s was not created: %v", dir, err)
		}
		if !info.IsDir() {
			t.Errorf("%s is not a directory", dir)
		}
	}
}
