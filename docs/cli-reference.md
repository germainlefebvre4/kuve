# CLI Reference

Complete reference for all Kuve commands and options.

## Table of Contents

- [Global Options](#global-options)
- [Commands](#commands)
  - [kuve install](#kuve-install)
  - [kuve uninstall](#kuve-uninstall)
  - [kuve switch](#kuve-switch)
  - [kuve current](#kuve-current)
  - [kuve list](#kuve-list)
  - [kuve use](#kuve-use)
  - [kuve init](#kuve-init)
  - [kuve completion](#kuve-completion)
  - [kuve version](#kuve-version)
  - [kuve help](#kuve-help)
- [Version Formats](#version-formats)
- [Exit Codes](#exit-codes)
- [Environment Variables](#environment-variables)

## Global Options

These options are available for all commands:

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--verbose` | `-v` | Enable verbose output | `false` |
| `--help` | `-h` | Show help for command | - |

### Usage

```bash
kuve --verbose install v1.28.0
kuve -v switch v1.29.0
```

## Commands

### kuve install

Install a specific kubectl version from official Kubernetes releases.

#### Syntax

```bash
kuve install <version>
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `version` | Yes | The kubectl version to install |

#### Description

Downloads the kubectl binary for the specified version from `dl.k8s.io` and installs it to `~/.kuve/versions/<version>/`. The version can be specified with or without the 'v' prefix.

#### Examples

```bash
# Install with 'v' prefix
kuve install v1.28.0

# Install without 'v' prefix (both work)
kuve install 1.28.0

# Install latest stable version (get from 'kuve list remote')
kuve install v1.29.1
```

#### Output

**Success:**
```
Downloading kubectl v1.28.0 for linux/amd64...
Successfully installed kubectl v1.28.0
```

**Error (already installed):**
```
Version v1.28.0 is already installed
```

**Error (download failed):**
```
Error: failed to download kubectl: <reason>
```

#### Notes

- Downloads are platform-specific (linux/amd64, darwin/amd64, etc.)
- Requires internet connection
- Binary is stored in `~/.kuve/versions/<version>/kubectl`
- Does not automatically switch to the installed version

---

### kuve uninstall

Remove an installed kubectl version.

#### Syntax

```bash
kuve uninstall <version>
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `version` | Yes | The kubectl version to uninstall |

#### Description

Removes the specified kubectl version from `~/.kuve/versions/`. You cannot uninstall the currently active version - switch to another version first.

#### Examples

```bash
# Uninstall a version
kuve uninstall v1.27.0

# Uninstall without 'v' prefix
kuve uninstall 1.27.0
```

#### Output

**Success:**
```
Successfully uninstalled kubectl v1.27.0
```

**Error (currently active):**
```
Error: cannot uninstall the currently active version (v1.27.0)
Switch to another version first using 'kuve switch <version>'
```

**Error (not installed):**
```
Error: version v1.27.0 is not installed
```

#### Notes

- Permanently deletes the version directory
- Cannot uninstall the active version (safety feature)
- Use `kuve switch` to change active version first

---

### kuve switch

Switch to a different installed kubectl version.

#### Syntax

```bash
kuve switch <version>
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `version` | Yes | The kubectl version to switch to |

#### Description

Changes the active kubectl version by updating the symbolic link at `~/.kuve/bin/kubectl` to point to the specified version. The version must be installed first.

#### Examples

```bash
# Switch to a version
kuve switch v1.28.0

# Switch without 'v' prefix
kuve switch 1.28.0
```

#### Output

**Success:**
```
Switched to kubectl v1.28.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

**Error (not installed):**
```
Error: version v1.28.0 is not installed
Run 'kuve install v1.28.0' to install it first
```

#### Notes

- Instant operation (just updates symlink)
- Requires version to be installed
- Use `kuve current` to verify the switch
- Ensure `~/.kuve/bin` is in your PATH

---

### kuve current

Show the currently active kubectl version.

#### Syntax

```bash
kuve current
```

#### Arguments

None

#### Description

Displays the currently active kubectl version managed by Kuve by reading the symbolic link at `~/.kuve/bin/kubectl`.

#### Examples

```bash
kuve current
```

#### Output

**Success:**
```
Current kubectl version: v1.28.0
```

**Error (no version active):**
```
Error: no kubectl version is currently active
Run 'kuve install <version>' and 'kuve switch <version>' to set one up
```

#### Notes

- Fast operation (reads symlink)
- Compare with `kubectl version --client` for verification
- Returns error if no version is installed/active

---

### kuve list

List kubectl versions (installed or available).

#### Syntax

```bash
kuve list <subcommand>
```

#### Subcommands

- `installed` - List locally installed versions
- `remote` - List available remote versions

---

#### kuve list installed

List all kubectl versions installed locally.

##### Syntax

```bash
kuve list installed
```

##### Arguments

None

##### Description

Shows all kubectl versions installed in `~/.kuve/versions/`. The currently active version is marked with an asterisk (`*`).

##### Examples

```bash
kuve list installed
```

##### Output

**Success (with versions):**
```
Installed kubectl versions:
  v1.26.3
* v1.28.0
  v1.29.1

* = current version (v1.28.0)
```

**Success (no versions):**
```
No kubectl versions installed.
Run 'kuve install <version>' to install one.
```

##### Notes

- Sorted by version number
- Current version marked with `*`
- Fast operation (reads directory)

---

#### kuve list remote

Show the latest stable kubectl version available for download.

##### Syntax

```bash
kuve list remote
```

##### Arguments

None

##### Description

Fetches the latest stable kubectl version from the official Kubernetes releases. This is the recommended version for most users.

##### Examples

```bash
kuve list remote
```

##### Output

**Success:**
```
Latest stable version: v1.29.1
```

**Error (network issue):**
```
Error: failed to fetch stable version: <reason>
```

##### Notes

- Requires internet connection
- Fetches from official Kubernetes releases
- Use this version for `kuve install`

---

### kuve use

Use kubectl version from `.kubernetes-version` file or detect from cluster.

#### Syntax

```bash
kuve use [options]
```

#### Options

| Flag | Short | Description |
|------|-------|-------------|
| `--from-cluster` | `-c` | Detect version from current Kubernetes cluster |

#### Description

Without flags, searches for a `.kubernetes-version` file in the current and parent directories, then switches to that version. With `--from-cluster`, detects the Kubernetes version from the current cluster context and switches to the matching kubectl version.

If the specified version is not installed, it will be installed automatically.

#### Examples

##### Using Version File

```bash
# Create a version file first
kuve init v1.28.0

# Use the version from file
kuve use
```

##### Using Cluster Detection

```bash
# Detect from current cluster context
kuve use --from-cluster

# Shorthand
kuve use -c
```

#### Output

**Success (from file):**
```
Found version v1.28.0 in .kubernetes-version file
Switched to kubectl v1.28.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

**Success (from cluster):**
```
Detecting Kubernetes version from current cluster context...
Detected cluster version: v1.29.3 (using kubectl v1.29.0)
Version v1.29.0 is not installed. Installing...
Downloading kubectl v1.29.0 for linux/amd64...
Successfully installed kubectl v1.29.0
Switched to kubectl v1.29.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

**Error (no version file):**
```
Error: no .kubernetes-version file found in current or parent directories
Run 'kuve init <version>' to create one
```

**Error (cluster detection failed):**
```
Error: failed to detect cluster version: <reason>
Make sure you have a valid kubeconfig and cluster access
```

#### Notes

- Auto-installs missing versions
- Searches up directory tree for version files
- Cluster detection normalizes vendor-specific versions
- Requires `kubectl` access for cluster detection

---

### kuve init

Create a `.kubernetes-version` file in the current directory.

#### Syntax

```bash
kuve init [version]
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `version` | No | Version to write to file (uses current if omitted) |

#### Description

Creates a `.kubernetes-version` file in the current directory. If no version is specified, uses the currently active kubectl version.

#### Examples

```bash
# Create with specific version
kuve init v1.28.0

# Create with current version
kuve init
```

#### Output

**Success (with version):**
```
Created .kubernetes-version file with v1.28.0
```

**Success (current version):**
```
Created .kubernetes-version file with v1.28.0 (current version)
```

**Error (no active version):**
```
Error: no kubectl version is currently active
Please specify a version: kuve init <version>
```

**Error (file exists):**
```
Error: .kubernetes-version file already exists
```

#### Notes

- Creates file in current directory
- File contains only the version string
- Commit to version control for team sharing
- Use `kuve use` to apply the version

---

### kuve completion

Generate shell completion scripts.

#### Syntax

```bash
kuve completion <shell>
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `shell` | Yes | Shell type: `bash`, `zsh`, `fish`, or `powershell` |

#### Description

Generates shell completion scripts for enhanced command-line experience. Provides tab completion for commands, flags, and arguments.

#### Examples

##### Bash

```bash
# Current session
source <(kuve completion bash)

# System-wide (requires sudo)
kuve completion bash | sudo tee /etc/bash_completion.d/kuve

# Add to ~/.bashrc
echo 'source <(kuve completion bash)' >> ~/.bashrc
```

##### Zsh

```zsh
# Generate to completions directory
kuve completion zsh > "${fpath[1]}/_kuve"

# Reload completions
autoload -U compinit && compinit
```

##### Fish

```fish
# Generate to completions directory
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

##### PowerShell

```powershell
kuve completion powershell | Out-String | Invoke-Expression
```

#### Output

Outputs the completion script to stdout.

#### Notes

- Shell-specific installation
- Enables tab completion
- Includes command and flag completion
- Restart shell or source config after setup

---

### kuve version

Show Kuve version information.

#### Syntax

```bash
kuve version
kuve --version
```

#### Arguments

None

#### Description

Displays the version of Kuve currently installed.

#### Examples

```bash
kuve version
kuve --version
```

#### Output

```
kuve version dev
```

#### Notes

- Shows Kuve version, not kubectl version
- Use `kuve current` for kubectl version
- Use `kubectl version --client` for full kubectl info

---

### kuve help

Show help information.

#### Syntax

```bash
kuve help [command]
kuve [command] --help
```

#### Arguments

| Argument | Required | Description |
|----------|----------|-------------|
| `command` | No | Specific command to show help for |

#### Description

Displays help information for Kuve commands. Without arguments, shows general help. With a command name, shows detailed help for that command.

#### Examples

```bash
# General help
kuve help
kuve --help

# Command-specific help
kuve help install
kuve install --help

# Subcommand help
kuve help list
kuve list --help
```

#### Output

Shows usage, description, available commands, flags, and examples.

---

## Version Formats

Kuve accepts versions in multiple formats:

### With 'v' Prefix

```bash
kuve install v1.28.0
kuve switch v1.29.1
```

### Without 'v' Prefix

```bash
kuve install 1.28.0
kuve switch 1.29.1
```

### Normalization

Both formats are normalized internally:
- Input: `1.28.0` → Stored as: `v1.28.0`
- Input: `v1.28.0` → Stored as: `v1.28.0`

### Cluster Versions

When using `--from-cluster`, vendor-specific versions are normalized:

| Input | Output |
|-------|--------|
| `v1.28.3-eks-123456` | `v1.28.0` |
| `v1.29.5-gke.1308000` | `v1.29.0` |
| `v1.27.2+k3s1` | `v1.27.0` |
| `v1.26.8` | `v1.26.0` |

## Exit Codes

Kuve uses standard Unix exit codes:

| Code | Meaning | Example |
|------|---------|---------|
| `0` | Success | Command completed successfully |
| `1` | General error | Version not found, network error, etc. |

### Checking Exit Codes

```bash
# Bash/Zsh
kuve install v1.28.0
echo $?  # Prints exit code

# In scripts
if kuve switch v1.28.0; then
    echo "Switch successful"
else
    echo "Switch failed"
fi
```

## Environment Variables

Currently, Kuve does not use custom environment variables. It respects standard environment variables:

| Variable | Purpose |
|----------|---------|
| `HOME` | Used to determine `~/.kuve` directory location |
| `PATH` | Must include `~/.kuve/bin` for kubectl access |

### Setting PATH

```bash
# Bash/Zsh
export PATH="$HOME/.kuve/bin:$PATH"

# Fish
set -gx PATH "$HOME/.kuve/bin" $PATH
```

## Command Chaining

Kuve commands can be chained in shell scripts:

### Sequential Commands

```bash
# Install and switch
kuve install v1.28.0 && kuve switch v1.28.0

# List and install latest
LATEST=$(kuve list remote | grep -oP 'v\d+\.\d+\.\d+')
kuve install $LATEST && kuve switch $LATEST
```

### Conditional Execution

```bash
# Switch only if installed
if kuve list installed | grep -q "v1.28.0"; then
    kuve switch v1.28.0
else
    kuve install v1.28.0 && kuve switch v1.28.0
fi
```

### Error Handling

```bash
# Stop on error
set -e
kuve install v1.28.0
kuve switch v1.28.0

# Continue on error
kuve install v1.28.0 || echo "Install failed, continuing..."
kuve switch v1.28.0
```

## Related Documentation

- [Usage Guide](./usage.md) - Detailed usage examples
- [Installation Guide](./installation.md) - Setup instructions
- [Troubleshooting](./troubleshooting.md) - Common issues and solutions
