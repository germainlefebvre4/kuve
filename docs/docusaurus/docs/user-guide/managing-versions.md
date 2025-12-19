---
sidebar_position: 2
---

# Managing Versions

Complete guide to installing, switching, and managing kubectl versions with Kuve.

## Installing Versions

### Basic Installation

Install any kubectl version from official Kubernetes releases:

```bash
kuve install v1.29.1
```

### Installation Process

When you install a version, Kuve:

1. **Downloads** the kubectl binary from `dl.k8s.io`
2. **Verifies** the download was successful
3. **Stores** it in `~/.kuve/versions/<version>/kubectl`
4. **Makes** it executable
5. **Confirms** successful installation

### Installation Requirements

- **Internet connection** to download from dl.k8s.io
- **Disk space** (~50MB per version)
- **Write permissions** in `~/.kuve/` directory

### Platform Detection

Kuve automatically detects your platform:
- **Linux**: `linux/amd64`, `linux/arm64`
- **macOS**: `darwin/amd64`, `darwin/arm64`

### Version Already Installed

If the version is already installed:

```bash
$ kuve install v1.28.0
Version v1.28.0 is already installed
```

To reinstall, uninstall first:

```bash
kuve uninstall v1.28.0
kuve install v1.28.0
```

## Switching Versions

### Instant Switching

Kuve uses symbolic links for instant version switching:

```bash
kuve switch v1.29.0
```

### What Happens During Switch

1. **Validates** the version is installed
2. **Updates** the symlink: `~/.kuve/bin/kubectl → ~/.kuve/versions/v1.29.0/kubectl`
3. **Confirms** the switch

### Auto-Install on Switch

If a version isn't installed, Kuve prompts to install it:

```bash
$ kuve switch v1.30.0
Version v1.30.0 is not installed. Would you like to install it? [y/N]
```

### Verify Switch

Always verify after switching:

```bash
kuve current
kubectl version --client --short
```

## Listing Versions

### List Installed Versions

See all locally installed kubectl versions:

```bash
$ kuve list installed
Installed kubectl versions:
  v1.26.3
  v1.27.5
* v1.28.0
  v1.29.1

* = current version (v1.28.0)
```

#### Output Details

- **Listed versions**: All installed kubectl versions
- **Current marker** (`*`): Indicates the active version
- **Sorted**: Displayed in version order

### List Remote Versions

Check the latest stable kubectl version:

```bash
$ kuve list remote
Latest stable version: v1.29.1
```

This queries the official Kubernetes release information.

### No Versions Installed

If no versions are installed:

```bash
$ kuve list installed
No kubectl versions installed

Install a version using: kuve install <version>
```

## Uninstalling Versions

### Basic Uninstall

Remove versions you no longer need:

```bash
kuve uninstall v1.27.0
```

### Uninstall Process

1. **Validates** the version is installed
2. **Checks** if it's the current version
3. **Removes** the version directory
4. **Confirms** successful removal

### Cannot Uninstall Current Version

Kuve prevents uninstalling the active version:

```bash
$ kuve uninstall v1.28.0
Error: cannot uninstall the currently active version (v1.28.0)
Switch to another version first using 'kuve switch <version>'
```

**Solution:**

```bash
# Switch to another version first
kuve switch v1.29.0

# Now uninstall
kuve uninstall v1.28.0
```

### Version Not Installed

If trying to uninstall a non-existent version:

```bash
$ kuve uninstall v1.25.0
Error: version v1.25.0 is not installed
```

## Checking Current Version

### Show Active Version

Display the currently active kubectl version:

```bash
$ kuve current
Current kubectl version: v1.28.0
```

### No Version Active

If no version is set:

```bash
$ kuve current
No kubectl version is currently active

Install and switch to a version using:
  kuve install <version>
  kuve switch <version>
```

### Cross-Verify with kubectl

Always verify with kubectl itself:

```bash
kubectl version --client --short
# Output: Client Version: v1.28.0
```

## Version Storage

### Directory Structure

Kuve stores versions in a clean, organized structure:

```
~/.kuve/
├── bin/
│   ├── kuve          # Kuve executable
│   └── kubectl       # Symlink → versions/v1.28.0/kubectl
└── versions/
    ├── v1.27.5/
    │   └── kubectl   # kubectl v1.27.5 binary
    ├── v1.28.0/
    │   └── kubectl   # kubectl v1.28.0 binary
    └── v1.29.1/
        └── kubectl   # kubectl v1.29.1 binary
```

### Disk Space Usage

Each kubectl version uses approximately 50MB of disk space.

**Check usage:**

```bash
du -sh ~/.kuve/versions/*
# Output:
# 50M    ~/.kuve/versions/v1.27.5
# 50M    ~/.kuve/versions/v1.28.0
# 50M    ~/.kuve/versions/v1.29.1
```

### Manual Cleanup

You can manually remove version directories:

```bash
rm -rf ~/.kuve/versions/v1.27.0
```

:::warning
Always use `kuve uninstall` instead of manual removal to ensure proper cleanup.
:::

## Version Strategies

### Strategy 1: Minimal Versions

Keep only actively used versions:

```bash
# Install only what you need
kuve install v1.28.0  # Production cluster
kuve install v1.29.0  # Staging cluster

# Regularly clean up
kuve list installed
kuve uninstall <old-version>
```

**Benefits:**
- Minimal disk usage
- Easier management
- Clear purpose per version

### Strategy 2: Range Testing

Install a range for compatibility testing:

```bash
# Install multiple versions for testing
kuve install v1.27.0
kuve install v1.28.0
kuve install v1.29.0
kuve install v1.30.0
```

**Benefits:**
- Test across versions
- Validate manifests
- Ensure compatibility

### Strategy 3: Latest + LTS

Keep latest and long-term support versions:

```bash
# Latest stable
kuve install v1.29.1

# LTS versions
kuve install v1.27.9  # Previous LTS
kuve install v1.24.12 # Older LTS
```

**Benefits:**
- Support old clusters
- Use latest features
- Cover common cases

## Advanced Operations

### Batch Installation

Install multiple versions in sequence:

```bash
# Using a loop
for version in v1.27.0 v1.28.0 v1.29.0; do
    kuve install $version
done
```

### Version Migration

Move from one version to another across projects:

```bash
# Find all .kubernetes-version files
find ~/projects -name .kubernetes-version -exec cat {} \;

# Update them
find ~/projects -name .kubernetes-version -exec sh -c 'echo "v1.29.0" > {}' \;

# Verify and switch
cd ~/projects/my-project
kuve use
```

### Backup Versions

Backup your installed versions:

```bash
# Create backup
tar -czf kuve-versions-backup.tar.gz ~/.kuve/versions/

# Restore backup
tar -xzf kuve-versions-backup.tar.gz -C ~/
```

## Troubleshooting

### Download Failures

If download fails:

1. **Check internet connection**
2. **Verify version exists**: `kuve list remote`
3. **Check dl.k8s.io status**
4. **Try again**: Downloads may be transient

### Permission Issues

If you get permission errors:

```bash
# Fix permissions
chmod +x ~/.kuve/versions/v1.28.0/kubectl

# Fix directory permissions
chmod -R u+w ~/.kuve/
```

### Corrupted Installation

If a version seems corrupted:

```bash
# Reinstall
kuve uninstall v1.28.0
kuve install v1.28.0
```

### Symlink Issues

If kubectl command doesn't work:

```bash
# Check symlink
ls -l ~/.kuve/bin/kubectl

# Recreate if needed
kuve switch v1.28.0
```

## Next Steps

- [**Version Files**](./version-files) - Use project-specific versions
- [**Workflows**](./workflows) - Learn advanced patterns
- [**Cluster Detection**](../advanced/cluster-detection) - Auto-detect versions
