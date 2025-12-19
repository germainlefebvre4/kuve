# Troubleshooting Guide

Common issues and solutions when using Kuve.

## Table of Contents

- [Installation Issues](#installation-issues)
- [PATH and Command Issues](#path-and-command-issues)
- [Version Management Issues](#version-management-issues)
- [Network and Download Issues](#network-and-download-issues)
- [Cluster Detection Issues](#cluster-detection-issues)
- [Version File Issues](#version-file-issues)
- [Permission Issues](#permission-issues)
- [Platform-Specific Issues](#platform-specific-issues)
- [Getting Help](#getting-help)

## Installation Issues

### Go Version Too Old

**Problem:**
```
go: module github.com/germainlefebvre4/kuve requires go >= 1.25
```

**Solution:**

Update Go to version 1.25 or later:

```bash
# Check current version
go version

# Download and install Go 1.25+ (Linux)
wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz

# Verify installation
go version
```

### Build Fails with Missing Dependencies

**Problem:**
```
cannot find package "github.com/spf13/cobra"
```

**Solution:**

Download dependencies:

```bash
cd kuve
go mod download
go mod tidy
make build
```

### Make Command Not Found

**Problem:**
```
bash: make: command not found
```

**Solution:**

Install make or build manually:

```bash
# Install make (Ubuntu/Debian)
sudo apt-get install build-essential

# Install make (macOS)
xcode-select --install

# Or build without make
go build -o kuve main.go
mkdir -p ~/.kuve/bin
mv kuve ~/.kuve/bin/
```

## PATH and Command Issues

### kuve: command not found

**Problem:**

Shell cannot find `kuve` command after installation.

**Diagnosis:**

```bash
# Check if kuve exists
ls -la ~/.kuve/bin/kuve

# Check if directory is in PATH
echo $PATH | grep -q ".kuve/bin" && echo "Found" || echo "Not found"
```

**Solution:**

Add `~/.kuve/bin` to your PATH:

```bash
# Bash
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Zsh
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# Fish
echo 'set -gx PATH "$HOME/.kuve/bin" $PATH' >> ~/.config/fish/config.fish
source ~/.config/fish/config.fish

# Verify
which kuve
kuve --version
```

### kubectl: command not found

**Problem:**

After installing and switching kubectl versions, `kubectl` command is not found.

**Diagnosis:**

```bash
# Check if kubectl symlink exists
ls -la ~/.kuve/bin/kubectl

# Check if any kubectl version is installed
kuve list installed

# Check PATH
echo $PATH | grep -q ".kuve/bin" && echo "Found" || echo "Not found"
```

**Solution:**

1. Ensure versions are installed:
   ```bash
   kuve install v1.28.0
   kuve switch v1.28.0
   ```

2. Verify symlink:
   ```bash
   ls -la ~/.kuve/bin/kubectl
   # Should show: kubectl -> ../versions/v1.28.0/kubectl
   ```

3. Add to PATH if missing:
   ```bash
   export PATH="$HOME/.kuve/bin:$PATH"
   ```

4. Verify:
   ```bash
   which kubectl
   kubectl version --client
   ```

### kubectl Uses Wrong Version

**Problem:**

`kubectl version --client` shows different version than `kuve current`.

**Diagnosis:**

```bash
# Check kuve's kubectl
kuve current
ls -la ~/.kuve/bin/kubectl

# Check which kubectl is being used
which kubectl
type kubectl

# Check all kubectl binaries in PATH
which -a kubectl  # Linux/macOS
```

**Solution:**

Ensure `~/.kuve/bin` is **first** in PATH:

```bash
# Check current PATH order
echo $PATH

# Fix PATH order (Bash/Zsh)
export PATH="$HOME/.kuve/bin:$PATH"

# Make permanent
# Edit ~/.bashrc or ~/.zshrc
# Place PATH export at the END of the file (after other PATH modifications)
```

Verify:
```bash
which kubectl
# Should be: /home/user/.kuve/bin/kubectl

kubectl version --client
kuve current
# Should match
```

## Version Management Issues

### Cannot Uninstall Current Version

**Problem:**
```
Error: cannot uninstall the currently active version (v1.28.0)
```

**Solution:**

Switch to a different version first:

```bash
# List installed versions
kuve list installed

# Switch to another version
kuve switch v1.29.0

# Now uninstall
kuve uninstall v1.28.0
```

### Version Not Found After Installation

**Problem:**

Installed a version but `kuve list installed` doesn't show it.

**Diagnosis:**

```bash
# Check versions directory
ls -la ~/.kuve/versions/

# Check specific version
ls -la ~/.kuve/versions/v1.28.0/
```

**Solution:**

Reinstall the version:

```bash
# Remove potentially corrupted installation
rm -rf ~/.kuve/versions/v1.28.0

# Reinstall
kuve install v1.28.0
```

### Cannot Switch to Installed Version

**Problem:**
```
Error: version v1.28.0 is not installed
```

But the version appears in `kuve list installed`.

**Diagnosis:**

```bash
# Check if binary exists and is executable
ls -la ~/.kuve/versions/v1.28.0/kubectl
file ~/.kuve/versions/v1.28.0/kubectl
```

**Solution:**

Ensure binary is executable:

```bash
chmod +x ~/.kuve/versions/v1.28.0/kubectl

# Try switching again
kuve switch v1.28.0
```

## Network and Download Issues

### Download Fails: Connection Timeout

**Problem:**
```
Error: failed to download kubectl: dial tcp: i/o timeout
```

**Solution:**

1. Check internet connection:
   ```bash
   ping -c 3 dl.k8s.io
   ```

2. Check proxy settings (if behind proxy):
   ```bash
   export HTTP_PROXY=http://proxy.example.com:8080
   export HTTPS_PROXY=http://proxy.example.com:8080
   kuve install v1.28.0
   ```

3. Try again later (temporary network issue)

### Download Fails: 404 Not Found

**Problem:**
```
Error: failed to download kubectl: 404 Not Found
```

**Solution:**

Verify the version exists:

```bash
# Check available version
kuve list remote

# Use exact version format
kuve install v1.28.0  # Not v1.28 or 1.28.0
```

### Slow Download

**Problem:**

Download takes very long or stalls.

**Solution:**

1. Check network speed:
   ```bash
   curl -o /dev/null https://dl.k8s.io/release/v1.28.0/bin/linux/amd64/kubectl
   ```

2. Use different network (mobile hotspot, VPN, etc.)

3. Download manually and install:
   ```bash
   # Download manually
   curl -LO "https://dl.k8s.io/release/v1.28.0/bin/linux/amd64/kubectl"
   
   # Install manually
   mkdir -p ~/.kuve/versions/v1.28.0
   mv kubectl ~/.kuve/versions/v1.28.0/
   chmod +x ~/.kuve/versions/v1.28.0/kubectl
   
   # Switch to it
   kuve switch v1.28.0
   ```

## Cluster Detection Issues

### Cannot Detect Cluster Version

**Problem:**
```
Error: failed to detect cluster version: unable to execute kubectl
```

**Diagnosis:**

```bash
# Check cluster connectivity
kubectl cluster-info

# Check current context
kubectl config current-context

# Check server version manually
kubectl version --short
```

**Solution:**

1. Ensure cluster is accessible:
   ```bash
   kubectl get nodes
   ```

2. Verify kubeconfig:
   ```bash
   echo $KUBECONFIG
   kubectl config view
   ```

3. Use correct context:
   ```bash
   kubectl config use-context <your-context>
   kuve use --from-cluster
   ```

### Cluster Version Detection Returns Wrong Version

**Problem:**

Detected version doesn't match expected cluster version.

**Diagnosis:**

```bash
# Check cluster version manually
kubectl version --short

# Check what kuve detects
kuve use --from-cluster --verbose
```

**Solution:**

This is usually due to version normalization. Kuve normalizes vendor-specific versions:

- `v1.28.3-eks-123456` → `v1.28.0`
- `v1.29.5-gke.100` → `v1.29.0`

This is **expected behavior** to match official kubectl releases.

### Cluster Access Denied

**Problem:**
```
Error: failed to detect cluster version: Unauthorized
```

**Solution:**

1. Check credentials:
   ```bash
   kubectl get nodes
   ```

2. Refresh credentials (cloud providers):
   ```bash
   # AWS EKS
   aws eks update-kubeconfig --name <cluster-name>
   
   # GCP GKE
   gcloud container clusters get-credentials <cluster-name>
   
   # Azure AKS
   az aks get-credentials --name <cluster-name> --resource-group <rg>
   ```

3. Try again:
   ```bash
   kuve use --from-cluster
   ```

## Version File Issues

### .kubernetes-version File Not Found

**Problem:**
```
Error: no .kubernetes-version file found in current or parent directories
```

**Solution:**

Create a version file:

```bash
# Create with specific version
kuve init v1.28.0

# Or create with current version
kuve init

# Verify
cat .kubernetes-version
```

### Version File Exists But Not Detected

**Problem:**

Created `.kubernetes-version` but `kuve use` doesn't find it.

**Diagnosis:**

```bash
# Check file exists
ls -la .kubernetes-version

# Check file content
cat .kubernetes-version

# Check current directory
pwd
```

**Solution:**

1. Ensure file is named correctly (with leading dot):
   ```bash
   # Correct
   .kubernetes-version
   
   # Incorrect
   kubernetes-version
   ```

2. Ensure file contains valid version:
   ```bash
   echo "v1.28.0" > .kubernetes-version
   ```

3. Run from correct directory:
   ```bash
   cd /path/to/project
   kuve use
   ```

### Invalid Version in File

**Problem:**
```
Error: invalid version format in .kubernetes-version file
```

**Solution:**

Fix version format:

```bash
# Check current content
cat .kubernetes-version

# Fix format (should be just the version)
echo "v1.28.0" > .kubernetes-version

# Or without 'v' prefix
echo "1.28.0" > .kubernetes-version

# Verify
cat .kubernetes-version
kuve use
```

## Permission Issues

### Permission Denied: Cannot Create Directory

**Problem:**
```
Error: failed to create directories: permission denied
```

**Solution:**

Ensure you have write permissions to home directory:

```bash
# Check permissions
ls -ld ~/

# Fix if needed (shouldn't be necessary)
chmod 755 ~/

# Create directories manually
mkdir -p ~/.kuve/bin
mkdir -p ~/.kuve/versions

# Try again
kuve install v1.28.0
```

### Permission Denied: Cannot Execute Binary

**Problem:**
```
bash: /home/user/.kuve/bin/kubectl: Permission denied
```

**Solution:**

Make binary executable:

```bash
# Fix permissions
chmod +x ~/.kuve/bin/kubectl
chmod +x ~/.kuve/versions/*/kubectl

# Verify
ls -la ~/.kuve/bin/kubectl
```

### Cannot Write to /usr/local/bin

**Problem:**

Trying to install Kuve system-wide without sudo.

**Solution:**

Use user-space installation (recommended):

```bash
# Don't install to /usr/local/bin
# Instead, use ~/.kuve/bin

make install  # Installs to ~/.kuve/bin
export PATH="$HOME/.kuve/bin:$PATH"
```

Or use sudo for system-wide:

```bash
# Build first
make build

# Install system-wide
sudo mv kuve /usr/local/bin/
```

## Platform-Specific Issues

### macOS: "kuve" cannot be opened

**Problem:**

macOS security prevents running unsigned binary.

**Solution:**

1. Allow the binary:
   ```bash
   xattr -d com.apple.quarantine ~/.kuve/bin/kuve
   ```

2. Or allow in System Preferences:
   - System Preferences → Security & Privacy → General
   - Click "Allow Anyway" next to kuve message

### macOS: Wrong Architecture Downloaded

**Problem:**

Downloaded version doesn't work on Apple Silicon (M1/M2).

**Diagnosis:**

```bash
# Check system architecture
uname -m

# Check binary architecture
file ~/.kuve/versions/v1.28.0/kubectl
```

**Solution:**

Kuve automatically detects architecture. If issues persist:

```bash
# Verify Go environment
go env GOOS GOARCH

# Rebuild Kuve
cd kuve
make clean
make build
```

### Linux: SELinux Prevents Execution

**Problem:**

SELinux blocks kubectl execution.

**Solution:**

```bash
# Check SELinux status
getenforce

# Allow execution
chcon -t bin_t ~/.kuve/bin/kubectl
chcon -t bin_t ~/.kuve/versions/*/kubectl

# Or temporarily disable SELinux (not recommended for production)
sudo setenforce 0
```

## Getting Help

### Enable Verbose Mode

Get more detailed output:

```bash
kuve --verbose install v1.28.0
kuve -v use --from-cluster
```

### Check Logs

Review command output for error messages:

```bash
kuve install v1.28.0 2>&1 | tee kuve-install.log
```

### Verify Installation

Run diagnostics:

```bash
# Check Kuve installation
which kuve
kuve --version

# Check directory structure
ls -laR ~/.kuve/

# Check PATH
echo $PATH

# Check current kubectl
which kubectl
kubectl version --client

# Check kuve's view
kuve current
kuve list installed
```

### Clean Reinstall

If all else fails, clean reinstall:

```bash
# Backup version files (optional)
cp .kubernetes-version .kubernetes-version.bak

# Remove Kuve completely
rm -rf ~/.kuve

# Remove PATH entry from shell config
# Edit ~/.bashrc or ~/.zshrc
# Remove: export PATH="$HOME/.kuve/bin:$PATH"

# Reinstall
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve
make install

# Re-add PATH
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Restore version file
mv .kubernetes-version.bak .kubernetes-version

# Reinstall kubectl versions
kuve install v1.28.0
kuve switch v1.28.0
```

### Report Issues

If you encounter a bug:

1. Check existing issues: https://github.com/germainlefebvre4/kuve/issues

2. Create new issue with:
   - Kuve version (`kuve --version`)
   - OS and version (`uname -a`)
   - Go version (`go version`)
   - Complete error message
   - Steps to reproduce
   - Output of diagnostic commands above

3. Include logs:
   ```bash
   kuve --verbose <command> 2>&1 | tee debug.log
   ```

## Common Solutions Summary

| Problem | Quick Fix |
|---------|-----------|
| Command not found | Add `~/.kuve/bin` to PATH |
| Wrong kubectl version | Ensure `~/.kuve/bin` is **first** in PATH |
| Cannot uninstall | Switch to different version first |
| Download fails | Check network, verify version exists |
| Cluster detection fails | Check `kubectl cluster-info` |
| Permission denied | Run `chmod +x` on binaries |
| Version file not found | Run `kuve init v1.28.0` |

## Related Documentation

- [Installation Guide](./installation.md) - Setup instructions
- [Usage Guide](./usage.md) - How to use Kuve
- [CLI Reference](./cli-reference.md) - All commands and options
- [Configuration](./configuration.md) - Configuration options
