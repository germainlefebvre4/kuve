# Cluster Version Detection

Kuve can automatically detect the Kubernetes version from your current cluster context and install/switch to the matching kubectl version.

## How It Works

The cluster detection feature:
1. Locates the kubectl binary (from kuve or system PATH)
2. Executes `kubectl version --output=json` to get the server version
3. Falls back to `kubectl version --short` for older kubectl versions
4. Parses the Kubernetes cluster version
5. **Normalizes the version** to match kubectl binary naming (removes vendor suffixes)
6. Installs the matching kubectl version (if not already installed)
7. Switches to that version

### Version Normalization

Kubernetes clusters often report versions with vendor-specific suffixes:
- GKE: `v1.33.5-gke.1308000`
- EKS: `v1.28.3-eks-123456`
- AKS: `v1.27.5-aks-20231015`
- K3s: `v1.26.8+k3s1`

Kuve automatically normalizes these to the base minor version:
- `v1.33.5-gke.1308000` → `v1.33.0`
- `v1.28.3-eks-123456` → `v1.28.0`
- `v1.27.5` → `v1.27.0`

This ensures compatibility with official kubectl releases.

## Usage

### Basic Usage

```bash
# Detect cluster version and switch
kuve use --from-cluster

# Shorthand
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

## Prerequisites

For cluster detection to work, you need:

1. **A configured Kubernetes context**
   ```bash
   kubectl config get-contexts
   ```

2. **Access to the cluster**
   ```bash
   kubectl cluster-info
   ```

3. **Network connectivity** to the cluster API server

## Use Cases

### 1. Working with Multiple Clusters

```bash
# Switch to production cluster
kubectl config use-context prod-cluster

# Auto-detect and use the matching kubectl version
kuve use --from-cluster

# Now kubectl matches your cluster version
kubectl version --short
```

### 2. New Cluster Setup

```bash
# Just got access to a new cluster
kubectl config use-context new-cluster

# Install and use the correct kubectl version
kuve use -c

# Start working immediately
kubectl get nodes
```

### 3. CI/CD Pipelines

```bash
#!/bin/bash
# deployment script

# Ensure kubectl version matches the cluster
kuve use --from-cluster

# Deploy application
kubectl apply -f deployment.yaml
```

### 4. Version Mismatch Warnings

If you see warnings like this:
```
WARNING: version difference between client (1.27) and server (1.28) exceeds the supported minor version skew of +/-1
```

Simply run:
```bash
kuve use --from-cluster
```

## Combining with Version Files

You can use both features together:

```bash
# For local development, use version file
cd my-project
kuve use  # Uses .kubernetes-version

# For deployment to cluster, detect from cluster
kuve use --from-cluster  # Uses cluster version
```

## Troubleshooting

### "kubectl not found"

Make sure kubectl is installed:
```bash
# Check if kubectl exists
which kubectl

# Or install a version with kuve first
kuve install v1.28.0
kuve switch v1.28.0
```

### "failed to get cluster version"

Check cluster connectivity:
```bash
# Verify you can connect to the cluster
kubectl cluster-info

# Check your current context
kubectl config current-context
```

### "could not parse server version"

This might happen with very old kubectl versions. Try:
```bash
# Manually check your cluster version
kubectl version --short

# Then install that version
kuve install v1.28.3
kuve switch v1.28.3
```

## Supported kubectl Versions

The cluster detection feature works with kubectl versions that support:
- JSON output format (`--output=json`) - kubectl 1.20+
- Short output format (`--short`) - kubectl 1.2+

## Notes

- Cluster versions are automatically normalized to base kubectl versions
- Vendor suffixes (GKE, EKS, AKS, K3s, etc.) are removed
- Patch versions are normalized to `.0` to match kubectl minor releases
- If the version is not installed, it will be downloaded automatically
- The feature requires access to the Kubernetes API server
- Network timeouts may cause detection to fail - check your cluster connectivity

## Best Practices

1. **Always verify after switching**
   ```bash
   kuve use --from-cluster
   kuve current
   kubectl version --short
   ```

2. **Use in automation scripts**
   ```bash
   #!/bin/bash
   set -e
   kuve use --from-cluster || exit 1
   kubectl apply -f manifests/
   ```

3. **Combine with context switching**
   ```bash
   # Function to switch context and kubectl version
   switch_cluster() {
     kubectl config use-context "$1"
     kuve use --from-cluster
   }
   
   switch_cluster prod-cluster
   ```

4. **Keep kuve updated**
   ```bash
   # Regularly update kuve to get the latest features
   cd kuve && git pull && make install
   ```
