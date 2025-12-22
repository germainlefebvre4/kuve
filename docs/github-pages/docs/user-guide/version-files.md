---
sidebar_position: 3
---

# Version Files

Use `.kubernetes-version` files to specify project-specific kubectl versions.

## What are Version Files?

Version files (`.kubernetes-version`) allow you to specify the required kubectl version for a project or directory. This ensures everyone working on the project uses the same kubectl version.

### Benefits

- **Team Consistency**: Everyone uses the same kubectl version
- **Self-Documenting**: Version requirement is explicit
- **Version Control**: Commit to Git for team-wide consistency

## Creating Version Files

### Using Current Version

Create a version file with the currently active version:

```bash
cd ~/my-project
kuve init
```

This creates `.kubernetes-version` containing the current version:

```
v1.28.0
```

### Specifying a Version

Create a version file with a specific version:

```bash
cd ~/my-project
kuve init v1.29.0
```

### Manual Creation

You can also create the file manually:

```bash
echo "v1.29.0" > .kubernetes-version
```

## Using Version Files

### Basic Usage

Switch to the version specified in `.kubernetes-version`:

```bash
cd ~/my-project
kuve use
```

### What Happens

When you run `kuve use`:

1. **Searches** for `.kubernetes-version` in current directory
2. **Reads** the version from the file
3. **Installs** the version if not already installed
4. **Switches** to that version
5. **Confirms** the switch

### Example Output

```bash
$ kuve use
Found .kubernetes-version file with version v1.29.0
Version v1.29.0 is not installed. Installing...
Downloading kubectl v1.29.0 for linux/amd64...
Successfully installed kubectl v1.29.0
Switched to kubectl v1.29.0
```

## Directory Hierarchy

Kuve searches for `.kubernetes-version` files in the current directory only (not parent directories).

### Example Structure

```
~/projects/
├── project-a/
│   ├── .kubernetes-version    # v1.28.0
│   └── manifests/
├── project-b/
│   ├── .kubernetes-version    # v1.29.0
│   └── deployments/
└── project-c/                 # No version file
```

### Behavior

```bash
# In project-a
cd ~/projects/project-a
kuve use  # Uses v1.28.0

# In project-b
cd ~/projects/project-b
kuve use  # Uses v1.29.0

# In project-c
cd ~/projects/project-c
kuve use  # Error: no .kubernetes-version file found
```

## Version Control Integration

### Add to Git

Always commit version files to version control:

```bash
cd ~/my-project
kuve init v1.28.0
git add .kubernetes-version
git commit -m "Add kubectl version requirement"
git push
```

### Team Workflow

When team members clone the repository:

```bash
git clone https://github.com/team/project.git
cd project
kuve use  # Automatically uses the project's kubectl version
```

### Update Version

When updating kubectl requirements:

```bash
# Update version file
kuve init v1.29.0

# Commit the change
git add .kubernetes-version
git commit -m "Update kubectl to v1.29.0"
git push
```

## File Format

### Basic Format

The file contains a single line with the version:

```
v1.29.0
```

### With Prefix

Both formats are supported:

```
v1.29.0
```

or

```
1.29.0
```

### Whitespace Handling

Leading and trailing whitespace is ignored:

```
  v1.29.0  
```

This is valid and will be read as `v1.29.0`.

### Comments Not Supported

The file should contain only the version:

```
# This is NOT supported
v1.29.0
```

## Common Workflows

### Workflow 1: New Project

```bash
# Create project structure
mkdir ~/my-new-project
cd ~/my-new-project

# Initialize git
git init

# Set kubectl version
kuve init v1.29.0

# Commit
git add .kubernetes-version
git commit -m "Initial commit with kubectl version"
```

### Workflow 2: Existing Project

```bash
# Clone existing project
git clone https://github.com/team/project.git
cd project

# Check if version file exists
cat .kubernetes-version

# Use the specified version
kuve use
```

### Workflow 3: Multiple Environments

```bash
# Production environment
cd ~/projects/prod-cluster
kuve init v1.28.0  # Stable LTS version

# Staging environment
cd ~/projects/staging-cluster
kuve init v1.29.0  # Latest stable

# Development environment
cd ~/projects/dev-cluster
kuve init v1.30.0  # Cutting edge
```

## Integration with CI/CD

### GitHub Actions

```yaml
name: Deploy
on: [push]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Install Kuve
        run: |
          curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
          chmod +x kuve
          sudo mv kuve /usr/local/bin/
      
      - name: Setup kubectl version
        run: |
          export PATH="$HOME/.kuve/bin:$PATH"
          kuve use
      
      - name: Deploy
        run: kubectl apply -f manifests/
```

### GitLab CI

```yaml
deploy:
  stage: deploy
  script:
    - curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
    - chmod +x kuve
    - mv kuve /usr/local/bin/
    - export PATH="$HOME/.kuve/bin:$PATH"
    - kuve use
    - kubectl apply -f manifests/
```

### Jenkins

```groovy
pipeline {
    agent any
    stages {
        stage('Setup kubectl') {
            steps {
                sh '''
                    curl -L https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64 -o kuve
                    chmod +x kuve
                    mv kuve /usr/local/bin/
                    export PATH="$HOME/.kuve/bin:$PATH"
                    kuve use
                '''
            }
        }
        stage('Deploy') {
            steps {
                sh 'kubectl apply -f manifests/'
            }
        }
    }
}
```

## Troubleshooting

### File Not Found

If `kuve use` reports no file found:

```bash
$ kuve use
Error: .kubernetes-version file not found in current directory

# Solution: Create the file
kuve init v1.28.0
```

### Invalid Version

If the version in the file is invalid:

```bash
$ kuve use
Error: invalid version format in .kubernetes-version

# Solution: Fix the version format
echo "v1.28.0" > .kubernetes-version
```

### Version Not Available

If the version doesn't exist:

```bash
$ kuve use
Error: version v1.99.0 is not available

# Solution: Check available versions
kuve list remote
kuve init v1.29.0
```


## Next Steps

- [**Workflows**](./workflows) - Learn advanced usage patterns
- [**Cluster Detection**](../advanced/cluster-detection) - Auto-detect cluster versions
