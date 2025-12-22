package version

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/germainlefebvre4/kuve/pkg/config"
)

const (
	// KubectlReleasesURL is the URL to fetch kubectl releases
	KubectlReleasesURL         = "https://dl.k8s.io/release/stable.txt"
	KubectlDownloadURLTemplate = "https://dl.k8s.io/release/%s/bin/%s/%s/kubectl"
)

// Manager handles version operations
type Manager struct {
	config *config.Config
}

// NewManager creates a new version manager
func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		config: cfg,
	}
}

// GetStableVersion fetches the latest stable kubectl version
func (m *Manager) GetStableVersion() (string, error) {
	resp, err := http.Get(KubectlReleasesURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch stable version: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	version := strings.TrimSpace(string(body))
	return version, nil
}

// ListRemoteVersions fetches available kubectl versions
// Returns the last 10 stable versions from GitHub releases
func (m *Manager) ListRemoteVersions() ([]string, error) {
	// Fetch releases from GitHub API
	resp, err := http.Get("https://api.github.com/repos/kubernetes/kubernetes/releases?per_page=50")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch releases from GitHub: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse JSON response
	var releases []struct {
		TagName    string `json:"tag_name"`
		Draft      bool   `json:"draft"`
		Prerelease bool   `json:"prerelease"`
	}

	if err := json.Unmarshal(body, &releases); err != nil {
		return nil, fmt.Errorf("failed to parse releases: %w", err)
	}

	// Filter and collect stable versions (exclude pre-releases, drafts, and RC versions)
	versions := []string{}
	versionRegex := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)

	for _, release := range releases {
		if !release.Draft && !release.Prerelease && versionRegex.MatchString(release.TagName) {
			versions = append(versions, release.TagName)
			if len(versions) == 10 {
				break
			}
		}
	}

	// Sort versions in descending order (newest first)
	sort.Slice(versions, func(i, j int) bool {
		return versions[i] > versions[j]
	})

	return versions, nil
}

// ListInstalledVersions lists all locally installed kubectl versions
func (m *Manager) ListInstalledVersions() ([]string, error) {
	entries, err := os.ReadDir(m.config.VersionsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, fmt.Errorf("failed to read versions directory: %w", err)
	}

	versions := []string{}
	versionRegex := regexp.MustCompile(`^v\d+\.\d+\.\d+$`)

	for _, entry := range entries {
		if entry.IsDir() && versionRegex.MatchString(entry.Name()) {
			versions = append(versions, entry.Name())
		}
	}

	sort.Strings(versions)
	return versions, nil
}

// GetCurrentVersion returns the currently active kubectl version
func (m *Manager) GetCurrentVersion() (string, error) {
	target, err := os.Readlink(m.config.CurrentSymlink)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("no kubectl version is currently active")
		}
		return "", fmt.Errorf("failed to read symlink: %w", err)
	}

	// Extract version from path (e.g., /home/user/.kuve/versions/v1.28.0/kubectl -> v1.28.0)
	parts := strings.Split(target, string(filepath.Separator))
	for i, part := range parts {
		if part == "versions" && i+1 < len(parts) {
			return parts[i+1], nil
		}
	}

	return "", fmt.Errorf("could not determine version from symlink")
}

// IsVersionInstalled checks if a specific version is installed
func (m *Manager) IsVersionInstalled(version string) bool {
	versionDir := filepath.Join(m.config.VersionsDir, version)
	kubectlBinary := filepath.Join(versionDir, config.KubectlBinaryName)

	info, err := os.Stat(kubectlBinary)
	if err != nil {
		return false
	}

	return !info.IsDir() && info.Mode()&0111 != 0 // Check if it's executable
}

// ReadVersionFile reads the .kubernetes-version file
func ReadVersionFile(dir string) (string, error) {
	versionFile := filepath.Join(dir, config.VersionFileName)
	data, err := os.ReadFile(versionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("failed to read version file: %w", err)
	}

	version := strings.TrimSpace(string(data))
	if version == "" {
		return "", nil
	}

	return version, nil
}

// FindVersionFile searches for .kubernetes-version file in current and parent directories
func FindVersionFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		version, err := ReadVersionFile(currentDir)
		if err != nil {
			return "", err
		}
		if version != "" {
			return version, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// Reached root directory
			break
		}
		currentDir = parentDir
	}

	return "", nil
}

// DetectClusterVersion detects the Kubernetes version from the current cluster context
// Returns the normalized kubectl version to install
func (m *Manager) DetectClusterVersion() (string, error) {
	_, normalizedVersion, err := m.detectClusterVersionRaw()
	if err != nil {
		return "", err
	}
	return normalizedVersion, nil
}

// DetectClusterVersionWithRaw detects the Kubernetes version and returns both raw and normalized versions
func (m *Manager) DetectClusterVersionWithRaw() (rawVersion, normalizedVersion string, err error) {
	return m.detectClusterVersionRaw()
}

// detectClusterVersionRaw is the internal implementation
func (m *Manager) detectClusterVersionRaw() (rawVersion, normalizedVersion string, err error) {
	// Check if kubectl exists in PATH or in kuve's bin
	kubectlPath, err := m.findKubectlBinary()
	if err != nil {
		return "", "", fmt.Errorf("kubectl not found: %w", err)
	}

	// Run kubectl version to get server version
	rawVersion, err = m.getServerVersion(kubectlPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to get cluster version: %w", err)
	}

	// Normalize version to base kubectl version (remove vendor suffixes)
	normalizedVersion = normalizeClusterVersion(rawVersion)

	return rawVersion, normalizedVersion, nil
}

// findKubectlBinary locates kubectl binary
func (m *Manager) findKubectlBinary() (string, error) {
	// First check if kuve's kubectl exists
	if _, err := os.Stat(m.config.CurrentSymlink); err == nil {
		return m.config.CurrentSymlink, nil
	}

	// Try to find kubectl in PATH
	paths := filepath.SplitList(os.Getenv("PATH"))
	for _, dir := range paths {
		kubectlPath := filepath.Join(dir, config.KubectlBinaryName)
		if info, err := os.Stat(kubectlPath); err == nil && !info.IsDir() {
			return kubectlPath, nil
		}
	}

	return "", fmt.Errorf("kubectl binary not found in PATH or kuve bin directory")
}

// getServerVersion executes kubectl to get the server version
func (m *Manager) getServerVersion(kubectlPath string) (string, error) {
	// Run kubectl version with JSON output
	cmd := exec.Command(kubectlPath, "version", "--output=json")
	output, err := cmd.Output()
	if err != nil {
		// Try fallback with short output
		return m.getServerVersionFallback(kubectlPath)
	}

	// Parse JSON output
	var versionInfo struct {
		ServerVersion struct {
			GitVersion string `json:"gitVersion"`
		} `json:"serverVersion"`
	}

	if err := json.Unmarshal(output, &versionInfo); err != nil {
		return m.getServerVersionFallback(kubectlPath)
	}

	if versionInfo.ServerVersion.GitVersion == "" {
		return "", fmt.Errorf("could not parse server version from kubectl output")
	}

	return versionInfo.ServerVersion.GitVersion, nil
}

// getServerVersionFallback tries to get version using short output
func (m *Manager) getServerVersionFallback(kubectlPath string) (string, error) {
	cmd := exec.Command(kubectlPath, "version", "--short")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute kubectl version: %w", err)
	}

	// Parse output like "Server Version: v1.28.0"
	lines := bytes.Split(output, []byte("\n"))
	for _, line := range lines {
		if bytes.Contains(line, []byte("Server Version:")) {
			parts := bytes.Fields(line)
			if len(parts) >= 3 {
				return string(parts[2]), nil
			}
		}
	}

	return "", fmt.Errorf("could not find server version in kubectl output")
}

// normalizeClusterVersion extracts the base kubectl version from cluster version
// Examples:
//
//	v1.33.5-gke.1308000 -> v1.33.5
//	v1.28.3-eks-123456 -> v1.28.3
//	v1.27.5 -> v1.27.5
func normalizeClusterVersion(clusterVersion string) string {
	// Remove leading/trailing whitespace
	clusterVersion = strings.TrimSpace(clusterVersion)

	// Extract version using regex: vMAJOR.MINOR.PATCH(-suffix)?
	versionRegex := regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(?:[-+].*)?$`)
	matches := versionRegex.FindStringSubmatch(clusterVersion)

	if len(matches) < 4 {
		// If regex doesn't match, return as-is
		return clusterVersion
	}

	major := matches[1]
	minor := matches[2]
	patch := matches[3]

	// Return base version in format vMAJOR.MINOR.PATCH (without suffixes)
	return fmt.Sprintf("v%s.%s.%s", major, minor, patch)
}
