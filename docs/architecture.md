# Architecture

This document describes the architecture, design decisions, and implementation details of Kuve.

## Table of Contents

- [Overview](#overview)
- [Design Principles](#design-principles)
- [System Architecture](#system-architecture)
- [Core Components](#core-components)
- [Data Flow](#data-flow)
- [File System Layout](#file-system-layout)
- [Technology Stack](#technology-stack)
- [Design Decisions](#design-decisions)

## Overview

Kuve (Kubernetes Client Switcher) is a command-line tool written in Go that manages multiple kubectl versions on a single system. It uses symbolic links and PATH manipulation to provide instant version switching without modifying system directories or requiring elevated privileges.

### Key Characteristics

- **Language**: Go 1.25
- **CLI Framework**: Cobra
- **Distribution**: Single binary
- **Installation**: User-space only (no sudo required after initial setup)
- **Architecture**: Modular, testable design

## Design Principles

### 1. Simplicity

- Single binary with no external dependencies
- User-space installation (no system-wide changes)
- Intuitive command structure
- Clear, actionable error messages

### 2. Safety

- Cannot uninstall active version
- Validates versions before switching
- Non-destructive operations
- Graceful error handling

### 3. Flexibility

- Support for multiple version formats
- Project-specific version files
- Cluster version detection
- Shell integration support

### 4. Performance

- Instant version switching via symlinks
- Cached version information
- Minimal disk space usage
- Parallel downloads (future enhancement)

### 5. User Experience

- Consistent command patterns
- Helpful output messages
- Shell completion support
- Progressive disclosure of features

## System Architecture

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                        User Interface                        │
│                      (CLI Commands)                          │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────────┐  │
│  │   Install   │  │    Switch    │  │      List        │  │
│  │   Command   │  │   Command    │  │    Command       │  │
│  └─────────────┘  └──────────────┘  └──────────────────┘  │
│                                                               │
│  ┌─────────────┐  ┌──────────────┐  ┌──────────────────┐  │
│  │     Use     │  │  Uninstall   │  │      Init        │  │
│  │   Command   │  │   Command    │  │    Command       │  │
│  └─────────────┘  └──────────────┘  └──────────────────┘  │
│                                                               │
├─────────────────────────────────────────────────────────────┤
│                       Business Logic                         │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────────────┐      ┌──────────────────────┐    │
│  │  Version Manager     │      │  kubectl Installer   │    │
│  │                      │      │                      │    │
│  │  - List versions     │      │  - Download binary   │    │
│  │  - Detect version    │      │  - Install to dir    │    │
│  │  - Search files      │      │  - Create symlinks   │    │
│  │  - Normalize version │      │  - Uninstall         │    │
│  └──────────────────────┘      └──────────────────────┘    │
│                                                               │
│  ┌──────────────────────┐                                    │
│  │  Config Manager      │                                    │
│  │                      │                                    │
│  │  - App directories   │                                    │
│  │  - Constants         │                                    │
│  │  - Initialization    │                                    │
│  └──────────────────────┘                                    │
│                                                               │
├─────────────────────────────────────────────────────────────┤
│                     External Services                        │
├─────────────────────────────────────────────────────────────┤
│                                                               │
│  ┌──────────────────────┐      ┌──────────────────────┐    │
│  │  Kubernetes API      │      │  dl.k8s.io           │    │
│  │  (Cluster Version)   │      │  (kubectl binaries)  │    │
│  └──────────────────────┘      └──────────────────────┘    │
│                                                               │
└─────────────────────────────────────────────────────────────┘
```

## Core Components

### 1. Command Layer (`cmd/`)

The command layer implements the CLI interface using Cobra.

#### Root Command (`cmd/root.go`)

- Entry point for all commands
- Global flags and configuration
- Version information
- Help system

#### Command Implementations

| Command | File | Purpose |
|---------|------|---------|
| `install` | `cmd/install.go` | Install kubectl versions |
| `uninstall` | `cmd/uninstall.go` | Remove kubectl versions |
| `switch` | `cmd/switch.go` | Switch active version |
| `current` | `cmd/switch.go` | Show current version |
| `list` | `cmd/list.go` | List versions |
| `use` | `cmd/use.go` | Use version from file/cluster |
| `init` | `cmd/use.go` | Create version file |

### 2. Business Logic Layer

#### Version Manager (`internal/version/manager.go`)

Responsible for version-related operations:

**Functions:**
- `GetStableVersion()` - Fetch latest stable kubectl version
- `ListInstalledVersions()` - List locally installed versions
- `GetCurrentVersion()` - Determine active version
- `ReadVersionFile()` - Read `.kubernetes-version` files
- `SearchVersionFile()` - Search directory tree for version files
- `DetectClusterVersion()` - Query cluster for version
- `NormalizeVersion()` - Normalize vendor-specific versions

**Implementation Details:**

```go
type VersionManager struct {
    config *config.Config
}

// Example: Version normalization
func (vm *VersionManager) NormalizeVersion(version string) string {
    // Remove vendor suffixes
    // v1.28.3-eks-123 -> v1.28.0
    // v1.29.5-gke.100 -> v1.29.0
    // v1.27.2+k3s1 -> v1.27.0
}
```

#### kubectl Installer (`internal/kubectl/installer.go`)

Manages kubectl binary installation and lifecycle:

**Functions:**
- `Install(version)` - Download and install kubectl
- `Uninstall(version)` - Remove kubectl version
- `Switch(version)` - Update symlink to version
- `DownloadBinary(version)` - Fetch from dl.k8s.io
- `CreateSymlink(target)` - Create/update symlink

**Download Logic:**

```go
// URL pattern for kubectl downloads
url := fmt.Sprintf(
    "https://dl.k8s.io/release/%s/bin/%s/%s/kubectl",
    version,    // e.g., v1.28.0
    runtime.GOOS,    // linux, darwin
    runtime.GOARCH,  // amd64, arm64
)
```

#### Config Manager (`pkg/config/config.go`)

Manages application configuration and directories:

**Constants:**
- `KuveDir` - Base directory (`~/.kuve`)
- `BinDir` - Binary directory (`~/.kuve/bin`)
- `VersionsDir` - Versions directory (`~/.kuve/versions`)
- `KubectlBinary` - Binary name (`kubectl`)
- `VersionFile` - Version file name (`.kubernetes-version`)

**Functions:**
- `NewConfig()` - Initialize configuration
- `EnsureDirectories()` - Create directory structure
- `GetVersionPath(version)` - Get path to specific version

### 3. Testing

Unit tests ensure reliability:

- `pkg/config/config_test.go` - Configuration tests
- `internal/version/manager_test.go` - Version manager tests

**Test Coverage:**
- Version normalization logic
- Directory path generation
- Version parsing and validation
- Error handling

## Data Flow

### Installing a kubectl Version

```
User Command: kuve install v1.28.0
       ↓
1. Parse version argument
       ↓
2. Normalize version (add/remove 'v' prefix)
       ↓
3. Check if already installed
       ↓
4. Download binary from dl.k8s.io
       ↓
5. Create version directory (~/.kuve/versions/v1.28.0/)
       ↓
6. Save binary to directory
       ↓
7. Set executable permissions
       ↓
8. Confirm installation
```

### Switching kubectl Versions

```
User Command: kuve switch v1.28.0
       ↓
1. Parse version argument
       ↓
2. Verify version is installed
       ↓
3. Resolve symlink target path
       ↓
4. Remove existing symlink (~/.kuve/bin/kubectl)
       ↓
5. Create new symlink pointing to version
       ↓
6. Verify symlink creation
       ↓
7. Confirm switch
```

### Using Version from File

```
User Command: kuve use
       ↓
1. Search for .kubernetes-version file
   ├─ Check current directory
   ├─ Check parent directories
   └─ Up to home directory
       ↓
2. Read version from file
       ↓
3. Check if version is installed
   ├─ If NO: Install version first
   └─ If YES: Continue
       ↓
4. Switch to version (symlink update)
       ↓
5. Confirm switch
```

### Using Version from Cluster

```
User Command: kuve use --from-cluster
       ↓
1. Locate kubectl binary (kuve or system)
       ↓
2. Execute: kubectl version --output=json
   └─ Fallback: kubectl version --short
       ↓
3. Parse server version from output
       ↓
4. Normalize version (remove vendor suffix)
       ↓
5. Check if version is installed
   ├─ If NO: Install version first
   └─ If YES: Continue
       ↓
6. Switch to version (symlink update)
       ↓
7. Confirm switch
```

## File System Layout

### Directory Structure

```
$HOME/
└── .kuve/                           # Base directory
    ├── bin/                         # Binaries directory
    │   ├── kuve                     # Kuve binary
    │   └── kubectl -> ../versions/v1.28.0/kubectl  # Symlink
    └── versions/                    # Installed versions
        ├── v1.26.3/
        │   └── kubectl              # kubectl v1.26.3 binary
        ├── v1.28.0/
        │   └── kubectl              # kubectl v1.28.0 binary
        └── v1.29.1/
            └── kubectl              # kubectl v1.29.1 binary
```

### Path Resolution

When `kubectl` is executed:

1. Shell looks for `kubectl` in PATH
2. Finds `~/.kuve/bin/kubectl` (if `.kuve/bin` is first in PATH)
3. Follows symlink to `~/.kuve/versions/v1.28.0/kubectl`
4. Executes the target binary

**Advantages:**
- Instant switching (just update symlink)
- No PATH changes needed
- Works with all shells
- Multiple versions coexist

## Technology Stack

### Core Technologies

| Technology | Version | Purpose |
|------------|---------|---------|
| **Go** | 1.25+ | Programming language |
| **Cobra** | Latest | CLI framework |
| **Go Modules** | - | Dependency management |

### Dependencies

```go
// go.mod
module github.com/germainlefebvre4/kuve

require (
    github.com/spf13/cobra v1.10.2
    github.com/spf13/pflag v1.0.10
    github.com/inconshreveable/mousetrap v1.1.0
)
```

### Standard Library Usage

- `os` - File system operations
- `path/filepath` - Path manipulation
- `io` - I/O operations
- `net/http` - HTTP downloads
- `encoding/json` - JSON parsing
- `runtime` - Platform detection
- `testing` - Unit tests

## Design Decisions

### 1. Why Symbolic Links?

**Decision**: Use symlinks instead of shell wrappers or PATH manipulation.

**Rationale:**
- Instant switching (no process overhead)
- Shell-agnostic solution
- Simple implementation
- Easy debugging (just check symlink)

**Alternatives Considered:**
- Shell function wrappers (shell-specific, complex)
- Copying binaries (slow, disk space)
- Dynamic PATH modification (fragile, state management)

### 2. Why User-Space Installation?

**Decision**: Install everything in `~/.kuve/` without requiring sudo.

**Rationale:**
- Security (no system-wide changes)
- Simplicity (no permission issues)
- Multi-user support (isolated installations)
- Easy uninstallation

### 3. Why Go?

**Decision**: Implement in Go instead of Bash or Python.

**Rationale:**
- Single binary distribution (no runtime dependencies)
- Cross-platform support (Linux, macOS, Windows future)
- Strong standard library (HTTP, JSON, file operations)
- Easy testing and CI/CD
- Performance (fast execution)

### 4. Why Cobra?

**Decision**: Use Cobra for CLI framework.

**Rationale:**
- Standard in Go ecosystem
- Built-in completion support
- Subcommand structure
- Flag parsing and validation
- Help generation

### 5. Version Normalization Strategy

**Decision**: Normalize cluster versions to minor version (v1.28.3 → v1.28.0).

**Rationale:**
- kubectl binaries released per minor version
- Compatibility across patch versions
- Vendor suffixes removed automatically
- Consistent user experience

**Examples:**
```
v1.28.3-eks-123456  →  v1.28.0
v1.29.5-gke.100     →  v1.29.0
v1.27.2+k3s1        →  v1.27.0
v1.26.8             →  v1.26.0
```

### 6. Configuration Management

**Decision**: Use file-based configuration (`.kubernetes-version`).

**Rationale:**
- Version control friendly (git commit version files)
- Project-specific versioning
- Team collaboration
- Clear documentation
- Shell-agnostic

### 7. Error Handling Strategy

**Decision**: Fail fast with clear error messages.

**Rationale:**
- User safety (prevent destructive operations)
- Clear debugging information
- Actionable error messages
- No silent failures

**Example:**
```go
if currentVersion == targetVersion {
    return fmt.Errorf("cannot uninstall the currently active version (%s)", version)
}
```

## Performance Considerations

### Fast Operations

- **Version Switching**: O(1) - Just update symlink
- **Current Version Check**: O(1) - Read symlink target
- **List Installed**: O(n) - Read directory entries

### Slow Operations

- **Install Version**: Network-bound (download binary)
- **List Remote**: Network-bound (fetch from k8s releases)
- **Cluster Detection**: Network-bound (query cluster API)

### Optimization Strategies

1. **Caching**: Future enhancement for remote version list
2. **Parallel Downloads**: Future enhancement for batch installs
3. **Local Index**: Future enhancement for offline capability

## Security Considerations

### Download Verification

**Current**: Downloads from official `dl.k8s.io`

**Future Enhancements:**
- Checksum verification
- Signature validation
- HTTPS enforcement

### Permission Model

- No sudo required (user-space only)
- Binary execution controlled by user
- No system-wide modifications

### Input Validation

- Version string validation
- Path sanitization
- URL validation

## Extensibility

### Adding New Commands

1. Create command file in `cmd/`
2. Implement Cobra command structure
3. Register with root command
4. Add tests
5. Update documentation

### Adding New Platforms

1. Add platform detection in installer
2. Update download URL logic
3. Handle platform-specific paths
4. Test on target platform

### Adding New Features

1. Implement in appropriate layer
2. Add tests
3. Update CLI commands if needed
4. Document in user guides

## Future Enhancements

### Planned Features

- [ ] Checksum verification for downloads
- [ ] Parallel version installs
- [ ] Version caching for offline use
- [ ] Custom download sources
- [ ] Version aliases (latest, stable, etc.)
- [ ] Windows support
- [ ] Plugin system for extensions

### Possible Improvements

- Configuration file support
- Automatic cleanup of old versions
- Version recommendations based on cluster
- Integration with version managers (asdf, mise)
- Telemetry for version usage statistics

## Related Documentation

- [Development Guide](../DEVELOPMENT.md) - Developer setup and workflows
- [Contributing Guide](../CONTRIBUTING.md) - How to contribute
- [CLI Reference](./cli-reference.md) - All commands and options
- [Usage Guide](./usage.md) - How to use Kuve
