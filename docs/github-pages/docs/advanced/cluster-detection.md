---
sidebar_position: 1
---

# Cluster Detection

Automatically detect and use kubectl versions matching your Kubernetes cluster.

## Overview

Kuve can automatically detect the Kubernetes version from your current cluster context and install/switch to the matching kubectl version. This ensures compatibility between your kubectl client and the Kubernetes API server.

## How It Works

The cluster detection feature follows these steps:

1. **Locate kubectl**: Finds kubectl binary (from Kuve or system PATH)
2. **Query Cluster**: Executes `kubectl version --output=json` to get server version
3. **Fallback**: Uses `kubectl version --short` for older kubectl versions
4. **Parse Version**: Extracts the Kubernetes cluster version
5. **Normalize**: Removes vendor-specific suffixes (GKE, EKS, AKS, etc.)
6. **Install**: Downloads matching kubectl version if not installed
7. **Switch**: Changes active kubectl to the matching version

## Version Normalization

Kubernetes clusters often report versions with vendor-specific suffixes that don't correspond to official kubectl releases. Kuve automatically normalizes these.

### Supported Distributions

| Distribution | Example Cluster Version | Normalized Version |
|--------------|-------------------------|-------------------|
| **GKE** (Google) | `v1.33.5-gke.1308000` | `v1.33.0` |
| **EKS** (AWS) | `v1.28.3-eks-123456` | `v1.28.0` |
| **AKS** (Azure) | `v1.27.5-aks-20231015` | `v1.27.0` |
| **K3s** | `v1.26.8+k3s1` | `v1.26.0` |
| **MicroK8s** | `v1.29.0+microk8s1` | `v1.29.0` |
| **K0s** | `v1.28.4+k0s.0` | `v1.28.0` |
| **Standard** | `v1.29.1` | `v1.29.0` |

### Normalization Rules

1. **Remove vendor suffix**: Everything after `-` or `+` is removed
2. **Minor version**: Normalizes to `.0` patch version for compatibility
3. **Preserve major.minor**: Keeps the core Kubernetes version

**Examples:**

| Cluster Version        | Normalized Version |
|------------------------|--------------------|
| v1.33.5-gke.1308000    | v1.33.0            |
| v1.28.3-eks-123456     | v1.28.0            |
| v1.27.5-aks-20231015   | v1.27.0            |
| v1.26.8+k3s1           | v1.26.0            |
| v1.29.0+microk8s1      | v1.29.0            |

## Usage

### Basic Usage

Detect and switch to cluster version:

```bash
kuve use --from-cluster
```

### Short Form

```bash
kuve use -c
```

### Example Output

```bash
$ kuve use --from-cluster
Detecting Kubernetes version from current cluster context...
Detected cluster version: v1.33.5-gke.1308000 (using kubectl v1.33.0)
Version v1.33.0 is not installed. Installing...
Downloading kubectl v1.33.0 for linux/amd64...
Successfully installed kubectl v1.33.0
Switched to kubectl v1.33.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

### With Existing Version

If the version is already installed:

```bash
$ kuve use --from-cluster
Detecting Kubernetes version from current cluster context...
Detected cluster version: v1.28.5-eks-abc123 (using kubectl v1.28.0)
Switched to kubectl v1.28.0
```

## Prerequisites

For cluster detection to work, you need:

### Configured Kubernetes Context

```bash
# List contexts
kubectl config get-contexts

# Verify current context
kubectl config current-context

# Ensure context is set
kubectl config use-context my-cluster
```

### Cluster Access

```bash
# Verify connectivity
kubectl cluster-info

# Test API access
kubectl version --short
```

### Network Connectivity

- Access to Kubernetes API server
- Internet access for downloading kubectl (if not installed)

## Use Cases

### Working with Multiple Clusters

Switch between clusters with different versions:

```bash
# Production cluster (GKE v1.28)
kubectl config use-context prod-gke
kuve use --from-cluster
# Installs and uses kubectl v1.28.0

# Staging cluster (EKS v1.29)
kubectl config use-context staging-eks
kuve use --from-cluster
# Installs and uses kubectl v1.29.0

# Development cluster (K3s v1.30)
kubectl config use-context dev-k3s
kuve use --from-cluster
# Installs and uses kubectl v1.30.0
```

### New Cluster Onboarding

Quickly set up kubectl for a new cluster:

```bash
# Just got access to a new cluster
kubectl config use-context new-cluster

# Auto-configure kubectl version
kuve use --from-cluster

# Start working immediately
kubectl get nodes
```

### CI/CD Pipelines

Ensure correct kubectl version in automated workflows:

```bash
#!/bin/bash
# deploy.sh - Deploy with correct kubectl version

# Configure kubeconfig
export KUBECONFIG=/path/to/kubeconfig

# Auto-detect and use cluster version
kuve use --from-cluster

# Deploy
kubectl apply -f manifests/
```

### Multi-Cluster Management Script

```bash
#!/bin/bash
# switch-cluster.sh - Switch cluster and kubectl version

CLUSTER=$1

if [ -z "$CLUSTER" ]; then
    echo "Usage: $0 <cluster-name>"
    echo ""
    echo "Available clusters:"
    kubectl config get-contexts -o name
    exit 1
fi

# Switch context
kubectl config use-context "$CLUSTER"

# Auto-detect and use matching kubectl
kuve use --from-cluster

# Show current state
echo ""
echo "Active cluster: $CLUSTER"
echo "kubectl version:"
kubectl version --short
```

## Troubleshooting

### Cannot Detect Version

**Problem:**
```
Error: failed to detect cluster version
```

**Solutions:**

1. **Verify cluster access:**
   ```bash
   kubectl cluster-info
   kubectl version
   ```

2. **Check kubeconfig:**
   ```bash
   kubectl config view
   echo $KUBECONFIG
   ```

3. **Test manually:**
   ```bash
   kubectl version --output=json
   ```

### Wrong Version Detected

**Problem:** Detected version doesn't match cluster.

**Diagnosis:**

```bash
# Check what cluster reports
kubectl version --short

# Check normalized version
kuve use --from-cluster --verbose
```

**Solution:**

This might be a normalization bug. Report it on [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) with:
- Cluster distribution (GKE, EKS, AKS, etc.)
- Output of `kubectl version --short`
- Expected kubectl version
- Actual version Kuve selected

### Network Timeouts

**Problem:**
```
Error: timeout connecting to cluster
```

**Solutions:**

1. **Check connectivity:**
   ```bash
   kubectl cluster-info
   ```

2. **Verify credentials:**
   ```bash
   kubectl auth can-i get pods
   ```

3. **Check VPN/proxy:**
   - Ensure VPN is connected
   - Configure proxy if needed

## Advanced Usage

### Scripting

```bash
#!/bin/bash
# sync-kubectl-versions.sh - Sync kubectl for all clusters

CONTEXTS=$(kubectl config get-contexts -o name)

for context in $CONTEXTS; do
    echo "Processing cluster: $context"
    
    # Switch context
    kubectl config use-context "$context"
    
    # Detect and install version
    kuve use --from-cluster
    
    # Verify
    echo "  kubectl version: $(kuve current)"
    echo ""
done
```

### Integration with Cluster Switcher

```bash
# ~/.bashrc or ~/.zshrc
function kswitch() {
    local cluster=$1
    
    if [ -z "$cluster" ]; then
        echo "Available clusters:"
        kubectl config get-contexts -o name
        return 1
    fi
    
    # Switch Kubernetes context
    kubectl config use-context "$cluster"
    
    # Auto-switch kubectl version
    kuve use --from-cluster
    
    echo "Switched to cluster: $cluster"
}

# Usage
# kswitch prod-gke
# kswitch staging-eks
```

### CI/CD Integration

#### GitHub Actions

```yaml
- name: Setup kubectl for cluster
  run: |
    export PATH="$HOME/.kuve/bin:$PATH"
    kuve use --from-cluster
```

#### GitLab CI

```yaml
before_script:
  - export PATH="$HOME/.kuve/bin:$PATH"
  - kuve use --from-cluster
```

## Best Practices

### Use with Cluster Switching

Always run after switching clusters:

```bash
kubectl config use-context new-cluster
kuve use --from-cluster
```

### Verify Version Match

Check that kubectl matches cluster:

```bash
kubectl version --short
# Client: v1.28.0
# Server: v1.28.5-gke.1234
```

Minor versions should match (both v1.28.x).

### Document in Scripts

Make cluster detection explicit in scripts:

```bash
#!/bin/bash
# deploy.sh

# Ensure correct kubectl version
kuve use --from-cluster

# Now deploy
kubectl apply -f manifests/
```

### Cache for Performance

For repeated operations, cache version:

```bash
#!/bin/bash
# Switch kubectl once, then use multiple times

kuve use --from-cluster

# Multiple operations with same version
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl rollout status deployment/app
```

## Limitations

### Current Limitations

1. **Only detects from current context**: Doesn't detect from all contexts at once
2. **Requires cluster access**: Must be able to query cluster version
3. **Network dependency**: Needs internet to download kubectl
4. **Minor version normalization**: Always normalizes to `.0` patch version

### Future Enhancements

Potential improvements (not yet implemented):

- Detect from kubeconfig without cluster access
- Support for offline detection
- Custom normalization rules
- Batch detection for multiple clusters

## Related Topics

- [Version Normalization](./version-normalization) - Deep dive into version handling
- [Basic Usage](../user-guide/basic-usage) - Learn kubectl version management
- [Workflows](../user-guide/workflows) - Multi-cluster management patterns

## Next Steps

- Try it: `kuve use --from-cluster`
- Explore: [Workflows](../user-guide/workflows) for advanced patterns
