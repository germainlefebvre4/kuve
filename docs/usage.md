# Usage Guide

This comprehensive guide covers all aspects of using Kuve to manage kubectl versions.

## Table of Contents

- [Basic Concepts](#basic-concepts)
- [Managing kubectl Versions](#managing-kubectl-versions)
  - [Installing Versions](#installing-versions)
  - [Listing Versions](#listing-versions)
  - [Switching Versions](#switching-versions)
  - [Checking Current Version](#checking-current-version)
  - [Uninstalling Versions](#uninstalling-versions)
- [Version Files](#version-files)
  - [Creating Version Files](#creating-version-files)
  - [Using Version Files](#using-version-files)
  - [Directory Hierarchy](#directory-hierarchy)
- [Cluster Detection](#cluster-detection)
  - [Auto-Detect from Cluster](#auto-detect-from-cluster)
  - [Version Normalization](#version-normalization)
- [Common Workflows](#common-workflows)
- [Best Practices](#best-practices)

## Basic Concepts

### How Kuve Works

Kuve manages kubectl versions using:

1. **Version Storage**: Each kubectl version is stored in `~/.kuve/versions/<version>/`
2. **Symbolic Links**: The active version is linked via `~/.kuve/bin/kubectl`
3. **PATH Priority**: By placing `~/.kuve/bin` first in PATH, it takes precedence

### Version Format

Kuve accepts versions in multiple formats:
- `v1.29.1` (with 'v' prefix)
- `1.29.1` (without 'v' prefix)
- Both formats are normalized internally

## Managing kubectl Versions

### Installing Versions

Install specific kubectl versions from official Kubernetes releases.

#### Install Latest Stable Version

```bash
kuve list remote
# Shows: Latest stable version: v1.29.1

kuve install v1.29.1
```

#### Install Specific Version

```bash
# With 'v' prefix
kuve install v1.28.0

# Without 'v' prefix (both work)
kuve install 1.28.0
```

#### What Happens During Installation

1. Downloads the kubectl binary from `dl.k8s.io`
2. Verifies the download
3. Stores it in `~/.kuve/versions/<version>/kubectl`
4. Makes it executable
5. Confirms successful installation

**Example Output:**

```
Downloading kubectl v1.28.0 for linux/amd64...
Successfully installed kubectl v1.28.0
```

### Listing Versions

#### List Installed Versions

View all kubectl versions installed on your system:

```bash
kuve list installed
```

**Example Output:**

```
Installed kubectl versions:
  v1.26.3
* v1.28.0
  v1.29.1
```

The `*` indicates the currently active version.

#### List Available Remote Versions

Check the latest stable version available for download:

```bash
kuve list remote
```

**Example Output:**

```
Latest stable version: v1.29.1
```

### Switching Versions

Change the active kubectl version instantly using symbolic links.

#### Switch to Specific Version

```bash
kuve switch v1.28.0
```

**Example Output:**

```
Switched to kubectl v1.28.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

#### Verify the Switch

```bash
kuve current
# Output: Current kubectl version: v1.28.0

kubectl version --client
# Output: Client Version: v1.28.0
```

### Checking Current Version

Display the currently active kubectl version:

```bash
kuve current
```

**Example Output:**

```
Current kubectl version: v1.28.0
```

### Uninstalling Versions

Remove kubectl versions you no longer need.

#### Uninstall a Version

```bash
kuve uninstall v1.27.0
```

**Example Output:**

```
Successfully uninstalled kubectl v1.27.0
```

#### Important Restrictions

- **Cannot uninstall active version**: You must switch to a different version first

**Example Error:**

```bash
kuve uninstall v1.28.0  # Current version
# Output: Error: cannot uninstall the currently active version (v1.28.0)
```

**Solution:**

```bash
# Switch to another version first
kuve switch v1.29.1

# Now uninstall the previous version
kuve uninstall v1.28.0
```

## Version Files

Use `.kubernetes-version` files to specify project-specific kubectl versions.

### Creating Version Files

#### Initialize with Current Version

Create a `.kubernetes-version` file using the currently active kubectl version:

```bash
cd ~/my-project
kuve init
```

This creates a file containing the current version:

```bash
cat .kubernetes-version
# Output: v1.28.0
```

#### Initialize with Specific Version

Specify a version directly:

```bash
kuve init v1.29.1
```

The file will contain:

```
v1.29.1
```

### Using Version Files

#### Automatic Version Detection

When you run `kuve use` without arguments, it:

1. Searches for `.kubernetes-version` in the current directory
2. If not found, searches parent directories up to home
3. Reads the version from the file
4. Installs the version if not already installed
5. Switches to that version

**Example:**

```bash
cd ~/my-project
cat .kubernetes-version
# Output: v1.28.0

kuve use
```

**Example Output:**

```
Found .kubernetes-version file specifying v1.28.0
Switched to kubectl v1.28.0
```

#### Auto-Install Missing Versions

If the specified version isn't installed:

```bash
kuve use
```

**Example Output:**

```
Found .kubernetes-version file specifying v1.30.0
Version v1.30.0 is not installed. Installing...
Downloading kubectl v1.30.0 for linux/amd64...
Successfully installed kubectl v1.30.0
Switched to kubectl v1.30.0
```

### Directory Hierarchy

Kuve searches up the directory tree for `.kubernetes-version` files:

```
/home/user/projects/team-a/microservice-x/
├── .kubernetes-version (v1.28.0) ← Found first
└── ...

/home/user/projects/team-a/
├── .kubernetes-version (v1.27.0) ← Ignored
└── ...
```

**Example:**

```bash
cd /home/user/projects/team-a/microservice-x
kuve use
# Uses v1.28.0 (from closest .kubernetes-version file)

cd /home/user/projects/team-a
kuve use
# Uses v1.27.0 (from parent directory)
```

## Cluster Detection

Automatically detect and use the kubectl version matching your Kubernetes cluster.

### Auto-Detect from Cluster

Use the `--from-cluster` (or `-c`) flag to detect the cluster version:

```bash
kuve use --from-cluster
# or shorthand
kuve use -c
```

#### What Happens

1. Connects to your current Kubernetes context
2. Queries the cluster API server version
3. Normalizes the version (removes vendor suffixes)
4. Installs the matching kubectl version if needed
5. Switches to that version

**Example Output:**

```
Detecting Kubernetes version from current cluster context...
Detected cluster version: v1.29.3 (using kubectl v1.29.0)
Version v1.29.0 is not installed. Installing...
Downloading kubectl v1.29.0 for linux/amd64...
Successfully installed kubectl v1.29.0
Switched to kubectl v1.29.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

### Version Normalization

Different Kubernetes distributions report versions with vendor-specific suffixes. Kuve automatically normalizes these to match official kubectl releases.

#### Supported Distributions

| Distribution | Cluster Version Example | Normalized Version |
|--------------|-------------------------|-------------------|
| **GKE** | `v1.33.5-gke.1308000` | `v1.33.0` |
| **EKS** | `v1.28.3-eks-123456` | `v1.28.0` |
| **AKS** | `v1.27.5-aks-20231015` | `v1.27.0` |
| **K3s** | `v1.26.8+k3s1` | `v1.26.0` |
| **MicroK8s** | `v1.25.5+microk8s1` | `v1.25.0` |
| **Vanilla** | `v1.29.3` | `v1.29.0` |

#### Why Normalization

Kubectl binaries are released per minor version (e.g., v1.28.0, v1.29.0). Normalization ensures:
- Compatibility with official kubectl releases
- Consistent behavior across distributions
- Proper version matching

## Common Workflows

### Workflow 1: Project-Based Version Management

```bash
# Set up a new project
mkdir ~/new-project
cd ~/new-project

# Initialize with specific version
kuve init v1.28.0

# Later, when returning to the project
cd ~/new-project
kuve use  # Automatically switches to v1.28.0

# Work with kubectl
kubectl get pods
```

### Workflow 2: Multi-Cluster Environment

```bash
# Working with production cluster (v1.29.0)
kubectl config use-context prod-cluster
kuve use --from-cluster
kubectl get nodes

# Switch to staging cluster (v1.28.0)
kubectl config use-context staging-cluster
kuve use --from-cluster
kubectl get nodes
```

### Workflow 3: Testing Across Versions

```bash
# Install multiple versions
kuve install v1.27.0
kuve install v1.28.0
kuve install v1.29.0

# Test with v1.27.0
kuve switch v1.27.0
kubectl apply -f deployment.yaml

# Test with v1.28.0
kuve switch v1.28.0
kubectl apply -f deployment.yaml

# Test with v1.29.0
kuve switch v1.29.0
kubectl apply -f deployment.yaml

# Check currently active
kuve current
```

### Workflow 4: Team Standardization

```bash
# Team lead sets up project
cd ~/team-project
kuve init v1.28.0
git add .kubernetes-version
git commit -m "Add kubectl version specification"
git push

# Team members clone and use
git clone <repository>
cd team-project
kuve use  # Automatically installs and switches to v1.28.0
```

## Best Practices

### 1. Use Version Files for Projects

Always create a `.kubernetes-version` file for projects:

```bash
kuve init v1.28.0
git add .kubernetes-version
git commit -m "Add kubectl version specification"
```

**Benefits:**
- Ensures team uses same kubectl version
- Documents version requirements
- Enables automatic version switching

### 2. Match Cluster Versions

Use cluster detection when working with different environments:

```bash
kubectl config use-context production
kuve use --from-cluster
```

**Benefits:**
- Ensures compatibility
- Reduces version skew issues
- Automates version management

### 3. Keep Versions Clean

Regularly remove unused versions:

```bash
# List installed versions
kuve list installed

# Remove old versions
kuve switch v1.29.0  # Switch away from old version
kuve uninstall v1.26.0
kuve uninstall v1.27.0
```

### 4. Verify After Switching

Always verify the switch worked:

```bash
kuve switch v1.28.0
kuve current
kubectl version --client
```

### 5. Use Shell Integration for Auto-Switching

Set up automatic version switching when entering directories. See [Shell Integration Guide](../SHELL_INTEGRATION.md).

### 6. Document Version Requirements

In your project README, document the required kubectl version:

```markdown
## Prerequisites

- kubectl v1.28.0 (managed via Kuve)

### Setup

1. Install Kuve: [Installation Guide](https://github.com/germainlefebvre4/kuve)
2. Install kubectl version: `kuve use`
```

## Troubleshooting

### kubectl: command not found

Ensure `~/.kuve/bin` is in your PATH:

```bash
echo $PATH | grep -q ".kuve/bin" && echo "Found" || echo "Not found"
export PATH="$HOME/.kuve/bin:$PATH"
```

### Version Not Switching

Verify the symlink:

```bash
ls -la ~/.kuve/bin/kubectl
# Should show: kubectl -> ../versions/v1.28.0/kubectl

# Fix if needed
kuve switch v1.28.0
```

### Cannot Connect to Cluster

When using `--from-cluster`, ensure:

```bash
# Check cluster connectivity
kubectl cluster-info

# Check current context
kubectl config current-context

# Verify credentials
kubectl get nodes
```

### Version File Not Found

Check file location:

```bash
# Search for version files
find . -name ".kubernetes-version" -type f

# Create if missing
kuve init v1.28.0
```

## Next Steps

- Learn about [Cluster Detection](../CLUSTER_DETECTION.md) in detail
- Set up [Shell Integration](../SHELL_INTEGRATION.md) for auto-switching
- Explore [CLI Reference](./cli-reference.md) for all command options
- Read [Troubleshooting Guide](./troubleshooting.md) for common issues
