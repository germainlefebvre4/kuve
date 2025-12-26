---
sidebar_position: 4
---

# Workflows

Learn advanced usage patterns and workflows for common scenarios.

## Multi-Project Management

### Scenario: Working with Multiple Projects

Manage kubectl versions across multiple projects with different requirements.

```bash
# Project A - Production (LTS version)
cd ~/projects/prod-app
kuve init v1.27.9
kuve use

# Project B - Staging (Current stable)
cd ~/projects/staging-app
kuve init v1.28.5
kuve use

# Project C - Development (Latest)
cd ~/projects/dev-app
kuve init v1.29.1
kuve use
```

## Multi-Cluster Management

### Scenario: Different Clusters, Different Versions

Work with multiple Kubernetes clusters requiring different kubectl versions.

```bash
# List your clusters
kubectl config get-contexts

# Production cluster (GKE v1.28)
kubectl config use-context prod-gke
kuve use --from-cluster  # Installs and uses v1.28.0

# Staging cluster (EKS v1.29)
kubectl config use-context staging-eks
kuve use --from-cluster  # Installs and uses v1.29.0

# Development cluster (K3s v1.30)
kubectl config use-context dev-k3s
kuve use --from-cluster  # Installs and uses v1.30.0
```

### Create Helper Script

```bash
#!/bin/bash
# switch-cluster.sh - Switch cluster and kubectl version

CONTEXT=$1

if [ -z "$CONTEXT" ]; then
    echo "Usage: $0 <context-name>"
    exit 1
fi

# Switch context
kubectl config use-context "$CONTEXT"

# Use matching kubectl version
kuve use --from-cluster

echo "Switched to context: $CONTEXT"
kuve current
```

Usage:
```bash
./switch-cluster.sh prod-gke
./switch-cluster.sh staging-eks
```

## Testing Across Versions

### Scenario: Manifest Compatibility Testing

Test Kubernetes manifests across multiple kubectl versions.

```bash
#!/bin/bash
# test-manifests.sh - Test manifests with multiple kubectl versions

VERSIONS=("v1.27.0" "v1.28.0" "v1.29.0" "v1.30.0")
MANIFESTS="manifests/*.yaml"

echo "Testing manifests across kubectl versions..."

for version in "${VERSIONS[@]}"; do
    echo ""
    echo "=== Testing with kubectl $version ==="

    # Install and switch version
    kuve install "$version" 2>/dev/null
    kuve switch "$version"

    # Validate manifests
    kubectl apply -f $MANIFESTS --dry-run=client

    if [ $? -eq 0 ]; then
        echo "✓ Manifests valid with $version"
    else
        echo "✗ Manifests failed with $version"
    fi
done
```

### Automated Testing Workflow

```yaml
# .github/workflows/test-kubectl-versions.yml
name: Test kubectl compatibility

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        kubectl-version: ['v1.27.0', 'v1.28.0', 'v1.29.0']

    steps:
      - uses: actions/checkout@v3

      - name: Install Kuve
        run: |
          wget https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64
          chmod +x kuve-linux-amd64
          sudo mv kuve-linux-amd64 /usr/local/bin/kuve

      - name: Setup kubectl
        run: |
          export PATH="$HOME/.kuve/bin:$PATH"
          kuve install ${{ matrix.kubectl-version }}
          kuve switch ${{ matrix.kubectl-version }}

      - name: Validate manifests
        run: kubectl apply -f manifests/ --dry-run=client
```

## CI/CD Integration

### GitHub Actions

Complete workflow for deploying with Kuve:

```yaml
name: Deploy to Kubernetes

on:
  push:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Install Kuve
        run: |
          curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
          chmod +x kuve
          sudo mv kuve /usr/local/bin/
          echo "$HOME/.kuve/bin" >> $GITHUB_PATH

      - name: Setup kubectl from version file
        run: kuve use

      - name: Configure kubeconfig
        run: |
          mkdir -p $HOME/.kube
          echo "${{ secrets.KUBECONFIG }}" > $HOME/.kube/config

      - name: Deploy
        run: |
          kubectl apply -f manifests/
          kubectl rollout status deployment/app
```

### GitLab CI/CD

```yaml
stages:
  - deploy

deploy_production:
  stage: deploy
  image: ubuntu:latest
  before_script:
    - apt-get update && apt-get install -y curl
    - curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
    - chmod +x kuve && mv kuve /usr/local/bin/
    - export PATH="$HOME/.kuve/bin:$PATH"
    - kuve use
  script:
    - kubectl apply -f manifests/
  only:
    - main
```

## Version Migration

### Scenario: Upgrading Kubectl Across Projects

Systematically upgrade kubectl version across all projects.

```bash
#!/bin/bash
# migrate-kubectl-version.sh - Upgrade all projects to new version

OLD_VERSION="v1.28.0"
NEW_VERSION="v1.29.0"

echo "Migrating projects from $OLD_VERSION to $NEW_VERSION"

# Find all .kubernetes-version files
for file in $(find ~/projects -name .kubernetes-version); do
    current=$(cat "$file" | tr -d '[:space:]')

    if [ "$current" = "$OLD_VERSION" ]; then
        echo "Updating: $file"
        echo "$NEW_VERSION" > "$file"

        # Optionally commit the change
        dir=$(dirname "$file")
        (cd "$dir" && git add .kubernetes-version && \
         git commit -m "Update kubectl to $NEW_VERSION")
    fi
done

# Install new version
kuve install "$NEW_VERSION"

echo "Migration complete!"
```

## Team Collaboration

### Scenario: Onboarding New Team Members

Standardize kubectl setup for new team members.

#### Setup Script

```bash
#!/bin/bash
# setup-dev-environment.sh - Onboard new developers

echo "Setting up development environment..."

# Install Kuve
if ! command -v kuve &> /dev/null; then
    echo "Installing Kuve..."
    curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-$(uname -s | tr '[:upper:]' '[:lower:]')-amd64 -o kuve
    chmod +x kuve
    mkdir -p ~/.local/bin
    mv kuve ~/.local/bin/
    export PATH="$HOME/.local/bin:$PATH"
    export PATH="$HOME/.kuve/bin:$PATH"
fi

# Clone project repositories
PROJECTS=(
    "prod-cluster"
    "staging-cluster"
    "dev-cluster"
)

for project in "${PROJECTS[@]}"; do
    echo "Setting up $project..."
    cd ~/projects

    if [ ! -d "$project" ]; then
        git clone "https://github.com/company/$project.git"
    fi

    cd "$project"

    # Install required kubectl version
    if [ -f .kubernetes-version ]; then
        kuve use
    fi
done

echo "Environment setup complete!"
echo "Run 'source ~/.bashrc' to update your PATH"
```

### Documentation for Team

```markdown
# Developer Setup Guide

## Prerequisites
- Git installed
- Internet connection

## Setup

1. Run the setup script:
   ```bash
   curl -sSL https://company.com/scripts/setup-dev.sh | bash
   ```

2. Reload shell:
   ```bash
   source ~/.bashrc
   ```

3. Verify installation:
   ```bash
   kuve --version
   kuve list installed
   ```
```

## Disaster Recovery

### Scenario: Restore Kuve Setup

Backup and restore your Kuve configuration.

#### Backup

```bash
#!/bin/bash
# backup-kuve.sh - Backup Kuve installation

BACKUP_DIR="$HOME/backups/kuve-$(date +%Y%m%d)"

mkdir -p "$BACKUP_DIR"

# Backup installed versions
echo "Backing up installed versions..."
cp -r ~/.kuve/versions "$BACKUP_DIR/"

# Backup version files from projects
echo "Backing up version files..."
find ~/projects -name .kubernetes-version -exec cp --parents {} "$BACKUP_DIR/" \;

# Create manifest
kuve list installed > "$BACKUP_DIR/installed-versions.txt"

echo "Backup complete: $BACKUP_DIR"
```

#### Restore

```bash
#!/bin/bash
# restore-kuve.sh - Restore Kuve from backup

BACKUP_DIR="$1"

if [ -z "$BACKUP_DIR" ]; then
    echo "Usage: $0 <backup-directory>"
    exit 1
fi

# Restore versions
echo "Restoring kubectl versions..."
cp -r "$BACKUP_DIR/versions" ~/.kuve/

# Fix permissions
chmod -R +x ~/.kuve/versions/*/kubectl

# Restore version files
echo "Restoring version files..."
cd "$BACKUP_DIR" && find . -name .kubernetes-version -exec cp --parents {} ~/ \;

echo "Restore complete!"
echo "Reinstall Kuve binary if needed: make install"
```

## Performance Optimization

### Scenario: Minimize Version Switches

Organize projects to minimize kubectl version switches.

```bash
# Group projects by kubectl version
~/projects/
├── kubectl-v1.27/
│   ├── legacy-app/
│   ├── old-service/
│   └── .kubernetes-version  # v1.27.0
├── kubectl-v1.28/
│   ├── prod-app/
│   ├── api-service/
│   └── .kubernetes-version  # v1.28.0
└── kubectl-v1.29/
    ├── new-app/
    ├── experimental/
    └── .kubernetes-version  # v1.29.0
```

### Pre-install Required Versions

```bash
#!/bin/bash
# preinstall-versions.sh - Install all required versions

# Read versions from all projects
VERSIONS=$(find ~/projects -name .kubernetes-version -exec cat {} \; | sort -u)

echo "Installing required kubectl versions..."

for version in $VERSIONS; do
    echo "Installing $version..."
    kuve install "$version"
done

echo "All versions installed!"
kuve list installed
```

## Next Steps

- [**Cluster Detection**](../advanced/cluster-detection) - Auto-detect cluster versions
- [**Version Normalization**](../advanced/version-normalization) - Version compatibility
- [**Troubleshooting**](../reference/troubleshooting) - Fix common issues
