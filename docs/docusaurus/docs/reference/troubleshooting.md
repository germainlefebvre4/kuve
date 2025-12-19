---
sidebar_position: 4
---

# Troubleshooting

Common issues and solutions when using Kuve.

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

# Install Go 1.25+ (Linux)
wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz

# Verify
go version
```

### Build Fails with Missing Dependencies

**Problem:**
```
cannot find package "github.com/spf13/cobra"
```

**Solution:**

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

**Problem:** Shell cannot find `kuve` command after installation.

**Diagnosis:**

```bash
# Check if kuve exists
ls -la ~/.kuve/bin/kuve

# Check if directory is in PATH
echo $PATH | grep -q ".kuve/bin" && echo "Found" || echo "Not found"
```

**Solution:**

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

**Problem:** After installing kubectl via Kuve, `kubectl` command is not found.

**Diagnosis:**

```bash
# Check if kubectl symlink exists
ls -la ~/.kuve/bin/kubectl

# Check if any version is installed
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
   ls -l ~/.kuve/bin/kubectl
   ```

3. Check PATH:
   ```bash
   export PATH="$HOME/.kuve/bin:$PATH"
   ```

### Wrong kubectl Version

**Problem:** Running `kubectl version --client` shows wrong version.

**Diagnosis:**

```bash
# Check what kubectl is being used
which kubectl

# Check Kuve's current version
kuve current

# Check all kubectl in PATH
which -a kubectl
```

**Solution:**

Ensure `~/.kuve/bin` is first in PATH:

```bash
# Bash/Zsh
export PATH="$HOME/.kuve/bin:$PATH"

# Verify
which kubectl
# Should output: /home/username/.kuve/bin/kubectl
```

## Version Management Issues

### Cannot Uninstall Current Version

**Problem:**
```
Error: cannot uninstall the currently active version (v1.28.0)
```

**Solution:**

Switch to another version first:

```bash
# List installed versions
kuve list installed

# Switch to different version
kuve switch v1.29.0

# Now uninstall
kuve uninstall v1.28.0
```

### Version Already Installed

**Problem:**
```
Version v1.28.0 is already installed
```

**Solution:**

If you need to reinstall:

```bash
kuve uninstall v1.28.0
kuve install v1.28.0
```

### Version Not Found

**Problem:**
```
Error: version v1.99.0 is not available
```

**Solution:**

```bash
# Check available versions
kuve list remote

# Install a valid version
kuve install v1.29.1
```

## Network and Download Issues

### Download Fails

**Problem:**
```
Error: failed to download kubectl: connection timeout
```

**Diagnosis:**

```bash
# Test connectivity
curl -I https://dl.k8s.io/

# Check if using proxy
echo $HTTP_PROXY
echo $HTTPS_PROXY
```

**Solution:**

1. Check internet connection
2. Configure proxy if needed:
   ```bash
   export HTTP_PROXY="http://proxy.example.com:8080"
   export HTTPS_PROXY="https://proxy.example.com:8443"
   ```

3. Try again:
   ```bash
   kuve install v1.28.0
   ```

### Slow Downloads

**Problem:** Downloads are very slow.

**Solution:**

1. Check network speed
2. Try during off-peak hours
3. Consider using a different network
4. Check if firewall/proxy is throttling

## Cluster Detection Issues

### Cannot Detect Cluster Version

**Problem:**
```
Error: failed to detect cluster version
```

**Diagnosis:**

```bash
# Verify cluster access
kubectl cluster-info

# Check current context
kubectl config current-context

# Test version query
kubectl version --output=json
```

**Solution:**

1. Verify kubeconfig:
   ```bash
   kubectl config view
   ```

2. Ensure cluster access:
   ```bash
   kubectl get nodes
   ```

3. Try again:
   ```bash
   kuve use --from-cluster
   ```

### Wrong Version Detected

**Problem:** Cluster version detected incorrectly.

**Diagnosis:**

```bash
# Check what kubectl reports
kubectl version --short

# Check cluster version
kubectl version --output=json | jq '.serverVersion'
```

**Solution:**

This might be a normalization issue. Report it on [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) with:
- Cluster version output
- Expected kubectl version
- Actual version Kuve selected

## Version File Issues

### Version File Not Found

**Problem:**
```
Error: .kubernetes-version file not found in current directory
```

**Solution:**

```bash
# Create version file
kuve init v1.28.0

# Or manually
echo "v1.28.0" > .kubernetes-version
```

### Invalid Version in File

**Problem:**
```
Error: invalid version format in .kubernetes-version
```

**Diagnosis:**

```bash
# Check file contents
cat .kubernetes-version

# Check for hidden characters
cat -A .kubernetes-version
```

**Solution:**

Fix the format:

```bash
# Correct format
echo "v1.28.0" > .kubernetes-version

# Or use kuve init
kuve init v1.28.0
```

## Permission Issues

### Permission Denied on Install

**Problem:**
```
Error: permission denied writing to /home/user/.kuve/versions/
```

**Solution:**

```bash
# Fix ownership
sudo chown -R $USER:$USER ~/.kuve/

# Fix permissions
chmod -R u+w ~/.kuve/
```

### Cannot Execute kubectl

**Problem:**
```
bash: /home/user/.kuve/bin/kubectl: Permission denied
```

**Solution:**

```bash
# Make executable
chmod +x ~/.kuve/bin/kubectl

# Or reinstall version
kuve switch v1.28.0
```

## Platform-Specific Issues

### macOS: Gatekeeper Warning

**Problem:**
```
"kubectl" cannot be opened because it is from an unidentified developer
```

**Solution:**

```bash
# Remove quarantine attribute
xattr -d com.apple.quarantine ~/.kuve/versions/v1.28.0/kubectl

# Or for all versions
find ~/.kuve/versions -name kubectl -exec xattr -d com.apple.quarantine {} \;
```

### macOS: Wrong Architecture

**Problem:** Running on Apple Silicon but got wrong binary.

**Solution:**

```bash
# Check architecture
uname -m
# Should be: arm64

# Reinstall (Kuve should auto-detect)
kuve uninstall v1.28.0
kuve install v1.28.0
```

### Linux: Library Missing

**Problem:**
```
error while loading shared libraries
```

**Solution:**

```bash
# Install required libraries (Ubuntu/Debian)
sudo apt-get update
sudo apt-get install -y ca-certificates

# Verify
ldd ~/.kuve/versions/v1.28.0/kubectl
```

## Symbolic Link Issues

### Broken Symlink

**Problem:**
```
ls: cannot access '/home/user/.kuve/bin/kubectl': No such file or directory
```

**Diagnosis:**

```bash
# Check symlink
ls -l ~/.kuve/bin/kubectl

# Should show target like:
# kubectl -> ../versions/v1.28.0/kubectl
```

**Solution:**

```bash
# Recreate symlink
kuve switch v1.28.0

# Or manually
ln -sf ~/.kuve/versions/v1.28.0/kubectl ~/.kuve/bin/kubectl
```

### Symlink Points to Wrong Version

**Problem:** Symlink doesn't match current version.

**Solution:**

```bash
# Fix by switching
kuve switch v1.28.0

# Verify
ls -l ~/.kuve/bin/kubectl
kuve current
```

## Shell Integration Issues

### Auto-Switch Not Working

**Problem:** Version doesn't change when entering directory.

**Diagnosis:**

```bash
# Check if function is defined
type kuve_auto_switch

# Check if .kubernetes-version exists
cat .kubernetes-version
```

**Solution:**

1. Verify shell integration is configured
2. Source shell config:
   ```bash
   source ~/.bashrc  # or ~/.zshrc
   ```

3. Test manually:
   ```bash
   kuve_auto_switch
   ```

### Completion Not Working

**Problem:** Tab completion doesn't work.

**Solution:**

```bash
# Bash
source <(kuve completion bash)

# Zsh
autoload -U compinit && compinit

# Verify
kuve <TAB>
# Should show: install, uninstall, switch, list, use, init, etc.
```

## Getting Help

### Enable Verbose Mode

Get more information about what's happening:

```bash
kuve --verbose install v1.28.0
kuve -v switch v1.29.0
```

### Check Version

```bash
kuve --version
```

### View Logs

Kuve doesn't have log files, but you can capture output:

```bash
kuve install v1.28.0 2>&1 | tee kuve-install.log
```

### Report Issues

If you encounter a bug:

1. Check [existing issues](https://github.com/germainlefebvre4/kuve/issues)
2. Gather information:
   ```bash
   kuve --version
   kuve list installed
   echo $PATH
   ls -la ~/.kuve/
   ```
3. [Open a new issue](https://github.com/germainlefebvre4/kuve/issues/new) with:
   - Kuve version
   - OS and architecture
   - Command you ran
   - Error output (verbose mode)
   - Expected vs actual behavior

## Emergency Reset

If Kuve is completely broken:

```bash
# Backup version files
find ~/projects -name .kubernetes-version -exec cp --parents {} ~/kuve-backup/ \;

# Remove Kuve
rm -rf ~/.kuve/

# Reinstall
cd kuve-source
make install

# Restore PATH
export PATH="$HOME/.kuve/bin:$PATH"

# Reinstall needed versions
kuve install v1.28.0
kuve switch v1.28.0
```

## Next Steps

- [Configuration](./configuration) - Review configuration
- [CLI Reference](./cli) - Command reference
- [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) - Report bugs
