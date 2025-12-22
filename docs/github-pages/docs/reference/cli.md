---
sidebar_position: 1
---

# CLI Reference

Complete reference for all Kuve commands and options.

## Synopsis

```bash
kuve [command] [flags]
```

## Global Flags

These flags are available for all commands:

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--verbose` | `-v` | Enable verbose output | `false` |
| `--help` | `-h` | Show help for command | - |

### Example

```bash
kuve --verbose install v1.28.0
kuve -v switch v1.29.0
```

## Commands

### kuve install

Install a specific kubectl version from official Kubernetes releases.

#### Synopsis

```bash
kuve install <version> [flags]
```

#### Description

Downloads and installs kubectl binary for the specified version from `dl.k8s.io`. The version can be specified with or without the 'v' prefix.

#### Arguments

- `version` (required): The kubectl version to install (e.g., `v1.29.0` or `1.29.0`)

#### Examples

```bash
# Install with 'v' prefix
kuve install v1.28.0

# Install without 'v' prefix
kuve install 1.28.0

# Verbose output
kuve install v1.29.0 --verbose
```

#### Output

```
Downloading kubectl v1.28.0 for linux/amd64...
Successfully installed kubectl v1.28.0
```

---

### kuve uninstall

Remove an installed kubectl version.

#### Synopsis

```bash
kuve uninstall <version> [flags]
```

#### Description

Removes the specified kubectl version from `~/.kuve/versions/`. You cannot uninstall the currently active version.

#### Arguments

- `version` (required): The kubectl version to uninstall

#### Examples

```bash
# Uninstall a version
kuve uninstall v1.27.0

# Uninstall without 'v' prefix
kuve uninstall 1.27.0
```

#### Output

```
Successfully uninstalled kubectl v1.27.0
```

---

### kuve switch

Switch the active kubectl version.

#### Synopsis

```bash
kuve switch <version> [flags]
```

#### Description

Changes the active kubectl version by updating the symbolic link in `~/.kuve/bin/kubectl`.

#### Arguments

- `version` (required): The kubectl version to switch to

#### Examples

```bash
# Switch to specific version
kuve switch v1.28.0

# Switch without 'v' prefix
kuve switch 1.28.0
```

#### Output

```
Switched to kubectl v1.28.0
Note: Make sure /home/user/.kuve/bin is in your PATH
```

---

### kuve current

Show the currently active kubectl version.

#### Synopsis

```bash
kuve current [flags]
```

#### Description

Displays the kubectl version currently in use.

#### Examples

```bash
kuve current
```

#### Output

```
Current kubectl version: v1.28.0
```

---

### kuve list

List kubectl versions.

#### Synopsis

```bash
kuve list [installed|remote] [flags]
```

#### Description

Lists either installed kubectl versions or the latest remote stable version.

#### Subcommands

##### kuve list installed

List all locally installed kubectl versions.

```bash
kuve list installed
```

**Output:**
```
Installed kubectl versions:
  v1.26.3
* v1.28.0
  v1.29.1

* = current version (v1.28.0)
```

##### kuve list remote

Show the latest stable kubectl version available.

```bash
kuve list remote
```

**Output:**
```
Latest stable version: v1.29.1
```

---

### kuve use

Use kubectl version from file or detect from cluster.

#### Synopsis

```bash
kuve use [flags]
```

#### Description

Switches kubectl version based on:
- `.kubernetes-version` file in current directory (default)
- Kubernetes cluster version (with `--from-cluster` flag)

#### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--from-cluster` | `-c` | Detect version from current cluster context |

#### Examples

```bash
# Use version from .kubernetes-version file
kuve use

# Detect from cluster
kuve use --from-cluster
kuve use -c
```

#### Output

```bash
# From file
Found .kubernetes-version file with version v1.28.0
Switched to kubectl v1.28.0

# From cluster
Detecting Kubernetes version from current cluster context...
Detected cluster version: v1.28.3-gke.1234 (using kubectl v1.28.0)
Switched to kubectl v1.28.0
```

---

### kuve init

Create a `.kubernetes-version` file.

#### Synopsis

```bash
kuve init [version] [flags]
```

#### Description

Creates a `.kubernetes-version` file in the current directory with either the specified version or the currently active version.

#### Arguments

- `version` (optional): The kubectl version to write to file. If omitted, uses current version.

#### Examples

```bash
# Use current version
kuve init

# Specify version
kuve init v1.28.0

# Without 'v' prefix
kuve init 1.28.0
```

#### Output

```
Created .kubernetes-version file with version v1.28.0
```

---

### kuve completion

Generate shell completion scripts.

#### Synopsis

```bash
kuve completion <shell> [flags]
```

#### Description

Generates completion scripts for the specified shell.

#### Supported Shells

- `bash`
- `zsh`
- `fish`
- `powershell`

#### Examples

```bash
# Bash - current session
source <(kuve completion bash)

# Bash - system-wide
kuve completion bash | sudo tee /etc/bash_completion.d/kuve

# Zsh
kuve completion zsh > "${fpath[1]}/_kuve"

# Fish
kuve completion fish > ~/.config/fish/completions/kuve.fish

# PowerShell
kuve completion powershell | Out-String | Invoke-Expression
```

---

### kuve version

Show Kuve version information.

#### Synopsis

```bash
kuve version [flags]
kuve --version [flags]
```

#### Description

Displays the version of Kuve.

#### Examples

```bash
kuve version
kuve --version
```

#### Output

```
kuve version dev
```

---

### kuve help

Show help information.

#### Synopsis

```bash
kuve help [command] [flags]
kuve [command] --help
```

#### Description

Displays help information for Kuve or a specific command.

#### Examples

```bash
# General help
kuve help
kuve --help

# Command-specific help
kuve help install
kuve install --help
```

## Version Formats

Kuve accepts versions in multiple formats:

### With 'v' Prefix

```bash
kuve install v1.29.0
kuve switch v1.28.0
```

### Without 'v' Prefix

```bash
kuve install 1.29.0
kuve switch 1.28.0
```

### Normalization

Both formats are normalized internally to use the 'v' prefix.

## Exit Codes

Kuve uses standard exit codes:

| Code | Meaning | Example Scenarios |
|------|---------|-------------------|
| `0` | Success | Command completed successfully |
| `1` | General error | Invalid arguments, version not found |
| `2` | Misuse | Incorrect command syntax |

### Examples

```bash
# Success
kuve install v1.28.0
echo $?  # Output: 0

# Error - version not installed
kuve switch v1.99.0
echo $?  # Output: 1

# Error - cannot uninstall current version
kuve uninstall v1.28.0  # if v1.28.0 is current
echo $?  # Output: 1
```

## Environment Variables

### PATH

The `PATH` environment variable must include `~/.kuve/bin`:

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

### HOME

Kuve uses `$HOME` to locate the `.kuve` directory:

```bash
~/.kuve/
├── bin/
└── versions/
```

### HTTP_PROXY / HTTPS_PROXY

Kuve respects proxy settings for downloads:

```bash
export HTTP_PROXY="http://proxy.example.com:8080"
export HTTPS_PROXY="https://proxy.example.com:8443"
```

## Configuration Files

### .kubernetes-version

Project-level version file.

**Location:** Project root directory

**Format:**
```
v1.28.0
```

**Usage:**
```bash
kuve use  # Reads from .kubernetes-version
```

## See Also

- [Commands](./commands) - Detailed command documentation
- [Configuration](./configuration) - Configuration options
- [Troubleshooting](./troubleshooting) - Common issues and solutions
