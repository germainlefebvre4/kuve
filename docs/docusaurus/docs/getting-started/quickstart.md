---
sidebar_position: 2
---

# Quick Start Guide

Get started with Kuve in just a few minutes!

## Step 1: Build and Install

```bash
# Clone the repository (if not already done)
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve

# Build and install
make install
```

Or build without installing:
```bash
make build
./kuve --help
```

## Step 2: Set Up Your Shell

Add Kuve's bin directory to your PATH:

```bash
# For Bash
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# For Zsh
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## Step 3: Install Your First kubectl Version

```bash
# Install the latest stable version
kuve install v1.29.1

# Or install a specific version
kuve install v1.28.0
```

**Example output:**
```
Downloading kubectl v1.29.1 for linux/amd64...
Successfully installed kubectl v1.29.1
```

## Step 4: Switch to a Version

```bash
# Switch to the installed version
kuve switch v1.29.1

# Verify it's active
kuve current
kubectl version --client
```

**Example output:**
```
Switched to kubectl v1.29.1
Note: Make sure /home/user/.kuve/bin is in your PATH
```

## Step 5: Try Version Files

Version files allow you to specify kubectl versions per project.

```bash
# Go to your project directory
cd ~/my-k8s-project

# Create a version file
kuve init v1.28.0

# The file is created
cat .kubernetes-version
# Output: v1.28.0

# Use the version from the file
kuve use
```

## Step 6: Try Cluster Detection

Automatically detect and use your cluster's kubectl version.

```bash
# Make sure you have a Kubernetes cluster configured
kubectl cluster-info

# Auto-detect and use the cluster version
kuve use --from-cluster

# Verify it worked
kuve current
```

## Common Commands

### List Versions

```bash
# List installed versions
kuve list installed

# List remote versions (shows latest stable)
kuve list remote
```

### Show Current Version

```bash
# Show current version
kuve current
```

### Uninstall a Version

```bash
# Uninstall a version (cannot uninstall current version)
kuve uninstall v1.27.0
```

## Typical Workflow

Here's a typical workflow when working with multiple projects:

```bash
# Project A needs kubectl v1.28.0
cd ~/projects/project-a
kuve init v1.28.0
kuve use

# Project B needs kubectl v1.29.0
cd ~/projects/project-b
kuve init v1.29.0
kuve use

# Back to project A - automatically uses v1.28.0
cd ~/projects/project-a
kuve use
```

## What's Next?

Now that you're up and running, explore these topics:

- **[Basic Usage](../user-guide/basic-usage)** - Learn all commands in detail
- **[Version Files](../user-guide/version-files)** - Master project-specific versions
- **[Cluster Detection](../advanced/cluster-detection)** - Understand auto-detection

## Need Help?

- View all commands: `kuve --help`
- Get command help: `kuve <command> --help`
- Check [Troubleshooting](../reference/troubleshooting) for common issues
