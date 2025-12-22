---
sidebar_position: 3
---

# Configuration

Kuve configuration, directory structure, and customization options.

## Directory Structure

Kuve uses a well-defined directory structure in the user's home directory.

### Default Layout

```
$HOME/
└── .kuve/                              # Base directory
    ├── bin/                            # Binary directory
    │   ├── kuve                        # Kuve binary
    │   └── kubectl -> ../versions/v1.28.0/kubectl  # Symlink to active version
    └── versions/                       # Installed kubectl versions
        ├── v1.26.3/
        │   └── kubectl                 # kubectl v1.26.3 binary
        ├── v1.28.0/
        │   └── kubectl                 # kubectl v1.28.0 binary
        └── v1.29.1/
            └── kubectl                 # kubectl v1.29.1 binary
```

### Directory Paths

| Directory | Path | Purpose |
|-----------|------|---------|
| **Base** | `~/.kuve/` | Root directory for all Kuve data |
| **Bin** | `~/.kuve/bin/` | Executable binaries (kuve, kubectl symlink) |
| **Versions** | `~/.kuve/versions/` | Installed kubectl versions |
| **Version Dir** | `~/.kuve/versions/v1.28.0/` | Specific version directory |

### Disk Usage

Each kubectl version uses approximately:
- **Linux/amd64**: ~50 MB
- **macOS/amd64**: ~52 MB
- **macOS/arm64**: ~50 MB

**Example** with 3 versions installed:
```bash
$ du -sh ~/.kuve/versions/*
50M    ~/.kuve/versions/v1.26.3
52M    ~/.kuve/versions/v1.28.0
52M    ~/.kuve/versions/v1.29.1
```

## Environment Variables

### PATH

**Required**: Must include `~/.kuve/bin`

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

**Verification:**
```bash
echo $PATH | grep -q ".kuve/bin" && echo "OK" || echo "Missing"
```

### HOME

Used to locate the `.kuve` directory:

```bash
# Kuve looks for: $HOME/.kuve/
```

### HTTP_PROXY / HTTPS_PROXY

Respected for kubectl downloads:

```bash
export HTTP_PROXY="http://proxy.example.com:8080"
export HTTPS_PROXY="https://proxy.example.com:8443"
```

## Version Files

### .kubernetes-version

Project-level version specification.

**Location:** Project root directory

**Format:**
```
v1.28.0
```

**Usage:**
```bash
kuve use  # Reads from .kubernetes-version
```

**Best Practices:**
- Commit to version control
- Place at project root
- Document in README
- Update with cluster upgrades

## Shell Configuration

### Bash

Add to `~/.bashrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"

# Optional: Shell completion
source <(kuve completion bash)
```

### Zsh

Add to `~/.zshrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"

# Optional: Shell completion
autoload -U compinit && compinit
```

### Fish

Add to `~/.config/fish/config.fish`:

```fish
# Kuve - Kubernetes Client Switcher
set -gx PATH "$HOME/.kuve/bin" $PATH

# Optional: Shell completion already in place
```

## Platform Support

### Operating Systems

| OS | Supported | Architecture |
|----|-----------|--------------|
| Linux | ✅ | amd64, arm64 |
| macOS | ✅ | amd64, arm64 |
| Windows | ❌ | Not supported |

### Download URLs

Kuve downloads kubectl from official sources:

```
https://dl.k8s.io/release/v{version}/bin/{os}/{arch}/kubectl
```

**Examples:**
- Linux: `https://dl.k8s.io/release/v1.28.0/bin/linux/amd64/kubectl`
- macOS: `https://dl.k8s.io/release/v1.28.0/bin/darwin/amd64/kubectl`

## Customization

### Custom Installation Path

Currently, Kuve uses a fixed path: `~/.kuve/`

To use a different location, create a symbolic link:

```bash
mkdir -p /custom/path/kuve
ln -s /custom/path/kuve ~/.kuve
```

### Custom Download Mirror

Not currently supported. Feature request: [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues)

## Security Considerations

### Binary Verification

Kuve downloads from official Kubernetes releases but doesn't currently verify checksums.

**Manual verification:**

```bash
# Download checksum
curl -LO "https://dl.k8s.io/release/v1.28.0/bin/linux/amd64/kubectl.sha256"

# Verify
echo "$(cat kubectl.sha256) ~/.kuve/versions/v1.28.0/kubectl" | sha256sum --check
```

### File Permissions

Kuve sets appropriate permissions:

```bash
# Binaries are executable
ls -l ~/.kuve/versions/v1.28.0/kubectl
# -rwxr-xr-x

# Directories are readable
ls -ld ~/.kuve/versions/v1.28.0/
# drwxr-xr-x
```

### User-Space Installation

Kuve operates entirely in user space:
- No system-wide changes
- No sudo required (after initial kuve installation)
- Safe for multi-user systems

## Advanced Configuration

### Multiple Kuve Installations

Not recommended, but possible:

```bash
# User A
export PATH="$HOME/.kuve/bin:$PATH"

# User B (different user on same system)
export PATH="$HOME/.kuve/bin:$PATH"
```

Each user has independent installations.

### CI/CD Configuration

#### GitHub Actions

```yaml
env:
  KUVE_VERSION: "latest"
  
steps:
  - name: Setup Kuve
    run: |
      curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
      chmod +x kuve
      sudo mv kuve /usr/local/bin/
      export PATH="$HOME/.kuve/bin:$PATH"
```

#### GitLab CI

```yaml
variables:
  KUVE_VERSION: "latest"

before_script:
  - curl -L https://github.com/germainlefebvre4/kuve/releases/$KUVE_VERSION/download/kuve-linux-amd64 -o kuve
  - chmod +x kuve && mv kuve /usr/local/bin/
  - export PATH="$HOME/.kuve/bin:$PATH"
```

## Maintenance

### Cleanup Old Versions

```bash
# List installed versions
kuve list installed

# Remove old versions
kuve uninstall v1.26.0
kuve uninstall v1.27.0
```

### Check Disk Usage

```bash
# Total Kuve disk usage
du -sh ~/.kuve/

# Per-version usage
du -sh ~/.kuve/versions/*

# Number of versions
ls ~/.kuve/versions/ | wc -l
```

### Verify Installation Health

```bash
#!/bin/bash
# health-check.sh - Verify Kuve installation

echo "Checking Kuve installation..."

# Check kuve binary
if command -v kuve &> /dev/null; then
    echo "✓ Kuve binary found"
    kuve --version
else
    echo "✗ Kuve binary not found"
fi

# Check PATH
if echo $PATH | grep -q ".kuve/bin"; then
    echo "✓ PATH configured"
else
    echo "✗ PATH not configured"
fi

# Check directory structure
if [ -d ~/.kuve/bin ] && [ -d ~/.kuve/versions ]; then
    echo "✓ Directory structure OK"
else
    echo "✗ Directory structure missing"
fi

# Check installed versions
echo ""
echo "Installed versions:"
kuve list installed
```

## Troubleshooting Configuration

### PATH Not Working

```bash
# Verify PATH
echo $PATH | grep .kuve

# Re-source shell config
source ~/.bashrc  # or ~/.zshrc
```

### Directory Permissions

```bash
# Fix permissions
chmod -R u+w ~/.kuve/
chmod +x ~/.kuve/bin/kuve
find ~/.kuve/versions -name kubectl -exec chmod +x {} \;
```

### Broken Symlinks

```bash
# Check symlink
ls -l ~/.kuve/bin/kubectl

# Recreate if broken
kuve switch <version>
```

## Next Steps

- [Troubleshooting](./troubleshooting) - Fix common issues
- [CLI Reference](./cli) - Command reference
- [Installation](../getting-started/installation) - Install Kuve
