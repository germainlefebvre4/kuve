---
sidebar_position: 2
---

# Commands

Detailed documentation for each Kuve command.

This page provides in-depth information about each command's behavior, use cases, and examples.

## install

Install a specific kubectl version.

### Usage

```bash
kuve install <version>
```

### Behavior

1. Validates version format
2. Checks if version already installed
3. Determines platform (OS/architecture)
4. Downloads kubectl binary from `dl.k8s.io`
5. Saves to `~/.kuve/versions/<version>/kubectl`
6. Makes binary executable
7. Confirms installation

### Requirements

- Internet connection
- ~50MB disk space per version
- Write permissions in `~/.kuve/`

### Error Conditions

| Error | Cause | Solution |
|-------|-------|----------|
| Already installed | Version exists | Use existing or uninstall first |
| Download failed | Network/URL issue | Check connection, verify version exists |
| Permission denied | No write access | Fix `~/.kuve/` permissions |
| Disk full | Insufficient space | Free up disk space |

### See Also

- [uninstall](#uninstall) - Remove versions
- [switch](#switch) - Activate versions

---

## uninstall

Remove an installed kubectl version.

### Usage

```bash
kuve uninstall <version>
```

### Behavior

1. Validates version is installed
2. Checks if version is currently active
3. Removes version directory
4. Confirms removal

### Restrictions

- Cannot uninstall the currently active version
- Must switch to another version first

### Error Conditions

| Error | Cause | Solution |
|-------|-------|----------|
| Not installed | Version doesn't exist | Check `kuve list installed` |
| Currently active | Version in use | Switch to another version first |
| Permission denied | No write access | Fix permissions |

### See Also

- [switch](#switch) - Change active version
- [list installed](#list-installed) - View installed versions

---

## switch

Change the active kubectl version.

### Usage

```bash
kuve switch <version>
```

### Aliases

- `kuve use <version>`

### Behavior

1. Validates version is installed
2. Updates symlink: `~/.kuve/bin/kubectl → ~/.kuve/versions/<version>/kubectl`
3. Confirms switch

### Technical Details

- Uses symbolic links for instant switching
- No file copying involved
- Changes take effect immediately

### Error Conditions

| Error | Cause | Solution |
|-------|-------|----------|
| Not installed | Version doesn't exist | Install version first |
| Invalid version | Bad format | Use correct format (v1.28.0) |

### Verification

After switching, verify:

```bash
kuve current
kubectl version --client
```

### See Also

- [current](#current) - Check active version
- [use](#use) - Use version from file/cluster

---

## current

Show the currently active kubectl version.

### Usage

```bash
kuve current
```

### Behavior

1. Reads symlink at `~/.kuve/bin/kubectl`
2. Extracts version from target path
3. Displays version

### Output Format

```
Current kubectl version: v1.28.0
```

### Use Cases

- Verify successful switch
- Check version before operations
- Debugging version issues
- Scripting version checks

### Error Conditions

| Error | Cause | Solution |
|-------|-------|----------|
| No version active | No symlink or broken | Install and switch to a version |

### See Also

- [switch](#switch) - Change active version
- [list installed](#list-installed) - View all versions

---

## list installed

List all locally installed kubectl versions.

### Usage

```bash
kuve list installed
```

### Behavior

1. Reads `~/.kuve/versions/` directory
2. Lists all version subdirectories
3. Marks currently active version with `*`
4. Sorts by version number

### Output Format

```
Installed kubectl versions:
  v1.26.3
* v1.28.0
  v1.29.1

* = current version (v1.28.0)
```

### Use Cases

- Audit installed versions
- Determine cleanup candidates
- Verify installations
- Check disk usage

### See Also

- [list remote](#list-remote) - Check latest version
- [uninstall](#uninstall) - Remove versions

---

## list remote

Show the latest stable kubectl version.

### Usage

```bash
kuve list remote
```

### Behavior

1. Queries official Kubernetes releases
2. Fetches latest stable version
3. Displays version

### Output Format

```
Latest stable version: v1.29.1
```

### Requirements

- Internet connection
- Access to Kubernetes release API

### Use Cases

- Check for updates
- Determine version to install
- Stay current with releases

### See Also

- [install](#install) - Install versions
- [list installed](#list-installed) - View local versions

---

## use

Use kubectl version from file or cluster.

### Usage

```bash
# From file
kuve use

# From cluster
kuve use --from-cluster
kuve use -c
```

### Mode: From File (Default)

#### Behavior

1. Searches for `.kubernetes-version` in current directory
2. Reads version from file
3. Installs version if needed
4. Switches to version

#### Requirements

- `.kubernetes-version` file must exist
- File must contain valid version

### Mode: From Cluster

#### Behavior

1. Connects to current Kubernetes cluster
2. Queries cluster version
3. Normalizes version (removes vendor suffixes)
4. Installs version if needed
5. Switches to version

#### Requirements

- Valid kubeconfig
- Cluster access
- `kubectl` binary available (from kuve or system)

### Version Normalization

Cluster versions are normalized:

- `v1.28.3-gke.1234` → `v1.28.0`
- `v1.27.5-eks-abc123` → `v1.27.0`
- `v1.29.1+k3s1` → `v1.29.0`

### See Also

- [init](#init) - Create version files
- [Cluster Detection](../advanced/cluster-detection) - Learn more

---

## init

Create a `.kubernetes-version` file.

### Usage

```bash
# Use current version
kuve init

# Specify version
kuve init <version>
```

### Behavior

1. Determines version (current or specified)
2. Creates `.kubernetes-version` in current directory
3. Writes version to file
4. Confirms creation

### File Format

```
v1.28.0
```

### Use Cases

- Set project kubectl requirement
- Document version dependency
- Enable version file workflow
- Team standardization

### Best Practices

1. Commit to version control:
   ```bash
   git add .kubernetes-version
   git commit -m "Add kubectl version requirement"
   ```

2. Create at project root

3. Document in README

### See Also

- [use](#use) - Use version from file
- [Version Files](../user-guide/version-files) - Learn more

---

## completion

Generate shell completion scripts.

### Usage

```bash
kuve completion <shell>
```

### Supported Shells

- `bash`
- `zsh`
- `fish`
- `powershell`

### Installation

#### Bash

```bash
# Current session
source <(kuve completion bash)

# Persistent
kuve completion bash | sudo tee /etc/bash_completion.d/kuve
```

#### Zsh

```bash
kuve completion zsh > "${fpath[1]}/_kuve"
autoload -U compinit && compinit
```

#### Fish

```bash
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

### Features

- Command completion
- Version completion (where applicable)
- Flag completion

### See Also

- [Shell Setup](../getting-started/shell-setup) - Configuration guide

---

## version

Show Kuve version information.

### Usage

```bash
kuve version
kuve --version
```

### Output Format

```
kuve version dev
```

### Use Cases

- Verify installation
- Report bugs (include version)
- Check for updates

---

## help

Display help information.

### Usage

```bash
kuve help [command]
kuve [command] --help
```

### Examples

```bash
# General help
kuve help

# Command help
kuve help install
kuve install --help
```

### Features

- Command descriptions
- Usage examples
- Available flags
- Related commands

## Command Comparison

| Command | Purpose | Requires Version | Modifies State |
|---------|---------|------------------|----------------|
| `install` | Add version | No | Yes (adds files) |
| `uninstall` | Remove version | Yes | Yes (removes files) |
| `switch` | Change active | Yes | Yes (symlink) |
| `current` | Show active | No | No |
| `list installed` | Show local | No | No |
| `list remote` | Show latest | No | No |
| `use` | Use from file/cluster | No* | Yes (symlink) |
| `init` | Create file | No | Yes (file) |

\* Auto-installs if needed

## Next Steps

- [Configuration](./configuration) - Configuration options
- [Troubleshooting](./troubleshooting) - Common issues
- [CLI Reference](./cli) - Quick reference
