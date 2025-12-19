---
sidebar_position: 1
slug: /
---

# Introduction to Kuve

Welcome to **Kuve** - the Kubernetes Client Switcher that makes managing multiple kubectl versions effortless!

## What is Kuve?

Kuve is a command-line tool that allows you to install, manage, and seamlessly switch between multiple kubectl versions on your system. Whether you're working with different Kubernetes clusters or testing across versions, Kuve simplifies version management.

## Key Features

### 📦 Version Management
- **Install** specific kubectl versions from official Kubernetes releases
- **Switch** instantly between installed versions
- **List** available and installed versions
- **Uninstall** versions you no longer need

### 📄 Project-Specific Versions
Use `.kubernetes-version` files to specify kubectl versions per project:
```bash
kuve init v1.28.0  # Creates .kubernetes-version file
kuve use           # Switches to version in file
```

### 🔍 Cluster Auto-Detection
Automatically detect and use the kubectl version matching your Kubernetes cluster:
```bash
kuve use --from-cluster
```

### 🚀 Smart Installation
Automatically installs missing versions when switching:
```bash
kuve switch v1.29.0  # Installs if not present
```

## Why Kuve?

### Simple & Safe
- No root access required after initial setup
- Cannot uninstall active versions
- Clear error messages
- Non-destructive operations

### Fast & Lightweight
- Instant switching using symbolic links
- Single binary with no dependencies
- Minimal disk space usage

### Flexible
- Works with multiple Kubernetes distributions (GKE, EKS, AKS, K3s, etc.)
- Supports version files for project consistency
- Shell completion for enhanced productivity

## Quick Example

```bash
# Install a kubectl version
kuve install v1.29.0

# Switch to it
kuve switch v1.29.0

# Verify
kubectl version --client
```

## Platform Support

- **Operating Systems**: Linux, macOS
- **Architectures**: amd64, arm64
- **Shells**: Bash, Zsh, Fish, PowerShell

## Next Steps

Ready to get started? Check out these guides:

- [**Installation**](./getting-started/installation) - Install Kuve on your system
- [**Quick Start**](./getting-started/quickstart) - Get up and running in minutes
- [**Basic Usage**](./user-guide/basic-usage) - Learn the essential commands

## Open Source

Kuve is open source and welcomes contributions! Visit our [GitHub repository](https://github.com/germainlefebvre4/kuve) to:
- Report issues
- Suggest features
- Contribute code
- Join discussions

## Getting Help

- 📖 [Documentation](./getting-started/installation)
- 🐛 [Issue Tracker](https://github.com/germainlefebvre4/kuve/issues)
- 💬 [Discussions](https://github.com/germainlefebvre4/kuve/discussions)
- 📝 [Troubleshooting Guide](./reference/troubleshooting)
