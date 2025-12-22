package version

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/germainlefebvre4/kuve/pkg/config"
)

func TestReadVersionFile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "kuve-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name        string
		content     string
		want        string
		shouldExist bool
	}{
		{
			name:        "valid version file",
			content:     "v1.28.0\n",
			want:        "v1.28.0",
			shouldExist: true,
		},
		{
			name:        "version without newline",
			content:     "v1.27.5",
			want:        "v1.27.5",
			shouldExist: true,
		},
		{
			name:        "empty version file",
			content:     "",
			want:        "",
			shouldExist: true,
		},
		{
			name:        "version with spaces",
			content:     "  v1.26.0  \n",
			want:        "v1.26.0",
			shouldExist: true,
		},
		{
			name:        "non-existent file",
			shouldExist: false,
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDir := filepath.Join(tmpDir, tt.name)
			os.MkdirAll(testDir, 0755)

			if tt.shouldExist {
				versionFile := filepath.Join(testDir, config.VersionFileName)
				if err := os.WriteFile(versionFile, []byte(tt.content), 0644); err != nil {
					t.Fatalf("Failed to write test file: %v", err)
				}
			}

			got, err := ReadVersionFile(testDir)
			if err != nil {
				t.Errorf("ReadVersionFile() error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("ReadVersionFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListInstalledVersions(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "kuve-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	versionsDir := filepath.Join(tmpDir, "versions")
	os.MkdirAll(versionsDir, 0755)

	cfg := &config.Config{
		HomeDir:        tmpDir,
		KuveDir:        tmpDir,
		BinDir:         filepath.Join(tmpDir, "bin"),
		VersionsDir:    versionsDir,
		CurrentSymlink: filepath.Join(tmpDir, "bin", "kubectl"),
	}

	manager := NewManager(cfg)

	// Test empty directory
	versions, err := manager.ListInstalledVersions()
	if err != nil {
		t.Errorf("ListInstalledVersions() error = %v", err)
	}
	if len(versions) != 0 {
		t.Errorf("Expected 0 versions, got %d", len(versions))
	}

	// Create some version directories
	testVersions := []string{"v1.28.0", "v1.27.5", "v1.26.3"}
	for _, v := range testVersions {
		vDir := filepath.Join(versionsDir, v)
		os.MkdirAll(vDir, 0755)
	}

	// Test with versions
	versions, err = manager.ListInstalledVersions()
	if err != nil {
		t.Errorf("ListInstalledVersions() error = %v", err)
	}
	if len(versions) != len(testVersions) {
		t.Errorf("Expected %d versions, got %d", len(testVersions), len(versions))
	}

	// Verify versions are sorted
	for i := 1; i < len(versions); i++ {
		if versions[i-1] > versions[i] {
			t.Errorf("Versions not sorted: %v", versions)
			break
		}
	}
}

func TestIsVersionInstalled(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "kuve-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	versionsDir := filepath.Join(tmpDir, "versions")
	os.MkdirAll(versionsDir, 0755)

	cfg := &config.Config{
		HomeDir:        tmpDir,
		KuveDir:        tmpDir,
		BinDir:         filepath.Join(tmpDir, "bin"),
		VersionsDir:    versionsDir,
		CurrentSymlink: filepath.Join(tmpDir, "bin", "kubectl"),
	}

	manager := NewManager(cfg)

	// Create a version directory with kubectl binary
	version := "v1.28.0"
	vDir := filepath.Join(versionsDir, version)
	os.MkdirAll(vDir, 0755)
	kubectlPath := filepath.Join(vDir, config.KubectlBinaryName)
	os.WriteFile(kubectlPath, []byte("fake binary"), 0755)

	// Test installed version
	if !manager.IsVersionInstalled(version) {
		t.Errorf("Expected version %s to be installed", version)
	}

	// Test non-existent version
	if manager.IsVersionInstalled("v1.27.0") {
		t.Errorf("Expected version v1.27.0 to not be installed")
	}
}

func TestFindKubectlBinary(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "kuve-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	binDir := filepath.Join(tmpDir, "bin")
	os.MkdirAll(binDir, 0755)

	cfg := &config.Config{
		HomeDir:        tmpDir,
		KuveDir:        tmpDir,
		BinDir:         binDir,
		VersionsDir:    filepath.Join(tmpDir, "versions"),
		CurrentSymlink: filepath.Join(binDir, "kubectl"),
	}

	manager := NewManager(cfg)

	// Test when kubectl doesn't exist in kuve bin (it might exist in system PATH)
	// We don't test for error here since system kubectl might exist
	_, _ = manager.findKubectlBinary()

	// Create kubectl symlink
	kubectlPath := cfg.CurrentSymlink
	targetPath := filepath.Join(tmpDir, "versions", "v1.28.0", "kubectl")
	os.MkdirAll(filepath.Dir(targetPath), 0755)
	os.WriteFile(targetPath, []byte("fake kubectl"), 0755)
	os.Symlink(targetPath, kubectlPath)

	// Test when kuve's kubectl exists - it should prefer kuve's version
	foundPath, err := manager.findKubectlBinary()
	if err != nil {
		t.Errorf("findKubectlBinary() error = %v", err)
	}
	if foundPath != kubectlPath {
		t.Errorf("Expected path %s, got %s", kubectlPath, foundPath)
	}
}

func TestNormalizeClusterVersion(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "GKE version with build suffix",
			input: "v1.33.5-gke.1308000",
			want:  "v1.33.5",
		},
		{
			name:  "EKS version with suffix",
			input: "v1.28.3-eks-123456",
			want:  "v1.28.3",
		},
		{
			name:  "AKS version with suffix",
			input: "v1.27.5-aks-20231015",
			want:  "v1.27.5",
		},
		{
			name:  "Standard version without suffix",
			input: "v1.29.2",
			want:  "v1.29.2",
		},
		{
			name:  "Version with plus suffix",
			input: "v1.26.8+k3s1",
			want:  "v1.26.8",
		},
		{
			name:  "Version without v prefix",
			input: "1.25.4-gke.1000",
			want:  "v1.25.4",
		},
		{
			name:  "Version with spaces",
			input: "  v1.30.1-custom  ",
			want:  "v1.30.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeClusterVersion(tt.input)
			if got != tt.want {
				t.Errorf("normalizeClusterVersion(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
