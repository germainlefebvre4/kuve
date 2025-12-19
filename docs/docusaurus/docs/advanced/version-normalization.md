---
sidebar_position: 3
---

# Version Normalization

Understanding how Kuve handles version formats and normalizes vendor-specific versions.

## Overview

Kubernetes clusters from different vendors (GKE, EKS, AKS, K3s, etc.) often report versions with vendor-specific suffixes. Kuve normalizes these to match official kubectl release versions.

## Why Normalization?

### The Problem

Official kubectl binaries are released with standard versions:
```
v1.28.0, v1.28.1, v1.28.2, v1.29.0, etc.
```

But managed Kubernetes services add vendor suffixes:
```
v1.28.3-gke.1234    (Google GKE)
v1.28.5-eks-abc123  (AWS EKS)
v1.27.9-aks-456def  (Azure AKS)
v1.26.8+k3s1        (K3s)
```

These vendor-specific versions don't exist as kubectl releases!

### The Solution

Kuve normalizes vendor versions to the nearest official kubectl release:

```
v1.28.3-gke.1234 → v1.28.0  (compatible kubectl version)
```

## Normalization Rules

### General Algorithm

1. **Extract base version**: Remove everything after `-` or `+`
2. **Normalize to minor**: Use `.0` for patch version
3. **Maintain major.minor**: Keep Kubernetes API version compatibility

### Examples by Distribution

#### Google Kubernetes Engine (GKE)

| Cluster Version        | kubectl Version |
|------------------------|-----------------|
| v1.33.5-gke.1308000    | v1.33.0         |
| v1.28.3-gke.1234       | v1.28.0         |
| v1.27.8-gke.567        | v1.27.0         |

#### Amazon Elastic Kubernetes Service (EKS)

| Cluster Version        | kubectl Version |
|------------------------|-----------------|
| v1.28.5-eks-abc123     | v1.28.0         |
| v1.27.9-eks-xyz789     | v1.27.0         |
| v1.26.12-eks-def456    | v1.26.0         |

#### Azure Kubernetes Service (AKS)

| Cluster Version        | kubectl Version |
|------------------------|-----------------|
| v1.27.9-aks-20231015   | v1.27.0         |
| v1.28.5-aks-20240201   | v1.28.0         |
| v1.29.1-aks-20240315   | v1.29.0         |

#### K3s (Lightweight Kubernetes)

| Cluster Version  | kubectl Version |
|------------------|-----------------|
| v1.26.8+k3s1     | v1.26.0         |
| v1.27.5+k3s2     | v1.27.0         |
| v1.28.3+k3s1     | v1.28.0         |

#### MicroK8s

| Cluster Version        | kubectl Version |
|------------------------|-----------------|
| v1.29.0+microk8s1      | v1.29.0         |
| v1.28.5+microk8s2      | v1.28.0         |

#### K0s

| Cluster Version       | kubectl Version |
|-----------------------|-----------------|
| v1.28.4+k0s.0         | v1.28.0         |
| v1.27.8+k0s.1         | v1.27.0         |

#### Standard Kubernetes

| Cluster Version | kubectl Version |
|-----------------|-----------------|
| v1.29.1         | v1.29.0         |
| v1.28.5         | v1.28.0         |
| v1.27.9         | v1.27.0         |

## Version Compatibility

### Kubernetes Version Skew Policy

kubectl is compatible with clusters that are:
- **Same minor version**: Perfect compatibility
- **±1 minor version**: Supported

Kuve's normalization ensures you get a compatible kubectl version.

### Examples

| Cluster Version | kubectl (Kuve) | Compatibility |
|----------------|----------------|---------------|
| v1.28.5-gke.1234 | v1.28.0 | ✅ Same minor |
| v1.29.1-eks-abc | v1.29.0 | ✅ Same minor |
| v1.27.8+k3s1 | v1.27.0 | ✅ Same minor |

All normalized versions maintain API compatibility!

## Technical Details

### Pattern Matching

Kuve uses the following regex patterns:

```regex
# Vendor suffix removal
^v?(\d+)\.(\d+)\.\d+[-+].*$  → v$1.$2.0

# Examples:
v1.28.3-gke.1234      → v1.28.0
v1.27.5+k3s1          → v1.27.0
```

### Implementation

```go
// Simplified version
func NormalizeVersion(version string) string {
    // Remove 'v' prefix if present
    version = strings.TrimPrefix(version, "v")
    
    // Split on '-' or '+'
    parts := strings.FieldsFunc(version, func(r rune) bool {
        return r == '-' || r == '+'
    })
    
    // Get base version (major.minor.patch)
    base := parts[0]
    
    // Extract major.minor
    versionParts := strings.Split(base, ".")
    if len(versionParts) >= 2 {
        return fmt.Sprintf("v%s.%s.0", versionParts[0], versionParts[1])
    }
    
    return "v" + base
}
```

## Use Cases

### Multi-Cloud Deployments

Working with clusters across different cloud providers:

```bash
# GKE cluster
kubectl config use-context prod-gke
kuve use --from-cluster
# Uses kubectl v1.28.0

# EKS cluster
kubectl config use-context prod-eks
kuve use --from-cluster
# Uses kubectl v1.28.0

# Same kubectl version for both!
```

### Hybrid Environments

Mix of managed and self-hosted Kubernetes:

```bash
# GKE managed cluster
kuve use --from-cluster  # v1.28.3-gke.1234 → v1.28.0

# Self-hosted K3s
kuve use --from-cluster  # v1.28.5+k3s1 → v1.28.0

# Standard Kubernetes
kuve use --from-cluster  # v1.28.1 → v1.28.0
```

### Testing Compatibility

Test manifests across distributions:

```bash
#!/bin/bash
# test-across-clouds.sh

CLUSTERS=(
    "prod-gke:v1.28.3-gke.1234"
    "prod-eks:v1.28.5-eks-abc"
    "prod-aks:v1.28.9-aks-xyz"
)

for cluster_info in "${CLUSTERS[@]}"; do
    context="${cluster_info%%:*}"
    
    echo "Testing on $context..."
    kubectl config use-context "$context"
    kuve use --from-cluster
    
    kubectl apply -f manifests/ --dry-run=client
done
```

## Verification

### Check Normalized Version

```bash
# See what version Kuve will use
kuve use --from-cluster --verbose

# Example output:
# Detected cluster version: v1.28.5-gke.1234 (using kubectl v1.28.0)
```

### Verify Compatibility

```bash
# Switch to cluster version
kuve use --from-cluster

# Check versions
kubectl version --short

# Example output:
# Client Version: v1.28.0
# Server Version: v1.28.5-gke.1234
```

Minor versions match! (1.28.x)

## Troubleshooting

### Unexpected Normalization

**Problem:** Kuve selected wrong kubectl version.

**Diagnosis:**

```bash
# Check cluster version
kubectl version --short

# See Kuve's normalization
kuve use --from-cluster --verbose
```

**Report Issues:**

If normalization seems incorrect, report on [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) with:
- Cloud provider/distribution
- Full cluster version string
- Expected kubectl version
- Actual version Kuve selected

### Version Not Available

**Problem:**
```
Error: version v1.33.0 is not available
```

**Cause:** Very new Kubernetes version not yet released as kubectl.

**Solution:**

Use the latest available:

```bash
# Check latest
kuve list remote

# Install latest
kuve install v1.29.1
kuve switch v1.29.1
```

## Related Topics

- [Cluster Detection](./cluster-detection) - Auto-detect cluster versions
- [Basic Usage](../user-guide/basic-usage) - kubectl version management
- [Workflows](../user-guide/workflows) - Multi-cluster patterns

## Next Steps

- Try: `kuve use --from-cluster`
- Read: [Cluster Detection](./cluster-detection) for full details
- Learn: [Kubernetes Version Skew Policy](https://kubernetes.io/releases/version-skew-policy/)
