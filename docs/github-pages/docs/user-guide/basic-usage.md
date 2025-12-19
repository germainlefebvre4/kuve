---
sidebar_position: 1
---

# Basic Usage

Learn the essential Kuve commands to manage kubectl versions effectively.

## Core Concepts

### How Kuve Works

Kuve manages kubectl versions using:

1. **Version Storage**: Each kubectl version is stored in `~/.kuve/versions/<version>/`
2. **Symbolic Links**: The active version is linked via `~/.kuve/bin/kubectl`
3. **PATH Priority**: By placing `~/.kuve/bin` first in PATH, it takes precedence

### Version Format

Kuve accepts versions in multiple formats:
- `v1.29.1` (with 'v' prefix) ✅
- `1.29.1` (without 'v' prefix) ✅
- Both formats are normalized internally

## Essential Commands

### Install a Version

Install specific kubectl versions from official Kubernetes releases.

```bash
# Install with 'v' prefix
kuve install v1.29.1

# Install without 'v' prefix (both work)
kuve install 1.29.1
```

**Example output:**
```
Downloading kubectl v1.29.1 for linux/amd64...
Successfully installed kubectl v1.29.1
```

:::tip
Use `kuve list remote` to find the latest stable version before installing.
:::

### List Installed Versions

View all kubectl versions installed on your system:

```bash
kuve list installed
```

**Example output:**
```raw
Installed kubectl versions:
  v1.29.1
  v1.33.0
* v1.33.5
  v1.34.3

* = current version (v1.33.5)
```

The `*` indicates the currently active version.

### List Remote Versions

Check the latest stable version available for download:

```bash
kuve list remote
```

**Example output:**
```raw
Last 10 stable kubectl versions:
  v1.35.0
  v1.34.3
  v1.34.2
  v1.34.1
  v1.33.7
  v1.33.6
  v1.33.5
  v1.32.11
  v1.32.10
  v1.31.14
```

### Switch Versions

Change the active kubectl version instantly:

```bash
kuve switch v1.28.0
```

**Example output:**
```
Switched to kubectl v1.28.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

:::info
Switching is instant - it just updates a symbolic link!
:::

### Check Current Version

Display the currently active kubectl version:

```bash
kuve current
```

**Example output:**
```
Current kubectl version: v1.28.0
```

Verify with kubectl:
```bash
kubectl version --client
```

### Uninstall a Version

Remove kubectl versions you no longer need:

```bash
kuve uninstall v1.27.0
```

**Example output:**
```
Successfully uninstalled kubectl v1.27.0
```

:::warning Important
You cannot uninstall the currently active version. Switch to another version first.
:::

## Command Aliases

Some commands have convenient aliases:

| Full Command | Alias | Description |
|-------------|-------|-------------|
| `kuve switch` | `kuve use` | Switch to a version |

Example:
```bash
kuve use v1.28.0  # Same as: kuve switch v1.28.0
```

## Working with Version Files

Create and use `.kubernetes-version` files for project-specific versions.

### Create a Version File

```bash
# Use current version
kuve init

# Specify a version
kuve init v1.28.0
```

This creates a `.kubernetes-version` file in the current directory.

### Use Version from File

```bash
kuve use
```

Kuve will:
1. Look for `.kubernetes-version` in current directory
2. Install the version if not present
3. Switch to that version

## Common Workflows

### Workflow 1: First-Time Setup

```bash
# 1. Install latest stable version
kuve list remote
kuve install v1.29.1

# 2. Switch to it
kuve switch v1.29.1

# 3. Verify
kuve current
kubectl version --client
```

### Workflow 2: Testing Across Versions

```bash
# Install multiple versions
kuve install v1.28.0
kuve install v1.29.0
kuve install v1.30.0

# Test with each version
kuve switch v1.28.0
kubectl apply -f manifest.yaml

kuve switch v1.29.0
kubectl apply -f manifest.yaml

kuve switch v1.30.0
kubectl apply -f manifest.yaml
```

### Workflow 3: Project-Specific Versions

```bash
# Project A
cd ~/projects/project-a
kuve init v1.28.0
kuve use

# Project B
cd ~/projects/project-b
kuve init v1.29.0
kuve use

# Back to A - easily switch
cd ~/projects/project-a
kuve use
```

### Workflow 4: Cleanup Old Versions

```bash
# List installed versions
kuve list installed

# Switch to latest
kuve switch v1.29.1

# Clean up old versions
kuve uninstall v1.26.0
kuve uninstall v1.27.0
```

## Best Practices

### Use Version Files

Always create `.kubernetes-version` files for projects:

```bash
cd ~/my-k8s-project
kuve init v1.28.0
git add .kubernetes-version
git commit -m "chore: Add kubectl version requirement"
```

**Benefits:**
- Team consistency
- Prevents version conflicts
- Self-documenting

### Match Cluster Versions

Use kubectl versions matching your Kubernetes cluster:

```bash
# Check cluster version
kubectl version --short

# Install matching kubectl version
kuve install v1.28.0
kuve switch v1.28.0
```

### Keep Minimal Versions

Only install versions you actively use:

```bash
# Regularly review and uninstall
kuve list installed
kuve uninstall <unused-version>
```

### Stay Updated

Periodically check for newer kubectl versions:

```bash
kuve list remote
kuve install <new-version>
```

## Command Reference Quick Guide

| Command | Purpose | Example |
|---------|---------|---------|
| `kuve install` | Install a version | `kuve install v1.29.0` |
| `kuve uninstall` | Remove a version | `kuve uninstall v1.28.0` |
| `kuve switch` | Change active version | `kuve switch v1.29.0` |
| `kuve current` | Show active version | `kuve current` |
| `kuve list installed` | Show installed versions | `kuve list installed` |
| `kuve list remote` | Show latest stable | `kuve list remote` |
| `kuve use` | Use version from file | `kuve use` |
| `kuve init` | Create version file | `kuve init v1.28.0` |

## Getting Help

### View All Commands

```bash
kuve --help
```

### Get Command-Specific Help

```bash
kuve install --help
kuve switch --help
kuve list --help
```

## Next Steps

- [**Managing Versions**](./managing-versions) - Deep dive into version management
- [**Version Files**](./version-files) - Master project-specific versions
- [**Workflows**](./workflows) - Learn advanced usage patterns
- [**CLI Reference**](../reference/cli) - Complete command documentation
