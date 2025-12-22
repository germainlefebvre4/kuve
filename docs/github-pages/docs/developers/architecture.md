---
sidebar_position: 1
---

# Architecture

System design, architecture, and implementation details of Kuve.

## Overview

Kuve is a command-line tool written in Go that manages multiple kubectl versions using symbolic links and PATH manipulation. It operates entirely in user space without requiring system-wide changes or elevated privileges.

### Key Characteristics

- **Language**: Go 1.25
- **CLI Framework**: Cobra
- **Distribution**: Single binary
- **Installation**: User-space only (no sudo required after initial setup)
- **Architecture**: Modular, testable design

## Design Principles

### Simplicity

- Single binary with minimal dependencies
- User-space installation (no system-wide changes)
- Intuitive command structure
- Clear, actionable error messages

### Safety

- Cannot uninstall active version
- Validates versions before switching
- Non-destructive operations
- Graceful error handling

### Flexibility

- Support for multiple version formats
- Project-specific version files
- Cluster version detection
- Shell integration support

### Performance

- Instant version switching via symlinks
- Minimal disk space usage
- Fast operations
- No runtime overhead

## System Architecture

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

## Project Structure

```
kuve/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command
│   ├── install.go         # Install kubectl
│   ├── uninstall.go       # Uninstall kubectl
│   ├── switch.go          # Switch versions
│   ├── list.go            # List versions
│   └── use.go             # Use from file/cluster
├── internal/              # Internal packages
│   ├── kubectl/           # kubectl management
│   │   └── installer.go   # Installation logic
│   └── version/           # Version management
│       ├── manager.go     # Version operations
│       └── manager_test.go # Tests
├── pkg/                   # Public packages
│   └── config/            # Configuration
│       ├── config.go      # Config logic
│       └── config_test.go # Tests
├── main.go               # Entry point
├── Makefile              # Build automation
└── go.mod                # Dependencies
```

## Core Components

### Command Layer (`cmd/`)

Implements CLI interface using Cobra.

| Command | File | Purpose |
|---------|------|---------|
| `install` | `cmd/install.go` | Install kubectl versions |
| `uninstall` | `cmd/uninstall.go` | Remove kubectl versions |
| `switch` | `cmd/switch.go` | Switch active version |
| `current` | `cmd/switch.go` | Show current version |
| `list` | `cmd/list.go` | List versions |
| `use` | `cmd/use.go` | Use from file/cluster |
| `init` | `cmd/use.go` | Create version file |

### Version Manager (`internal/version/`)

Handles version-related operations:

- `GetStableVersion()` - Fetch latest stable version
- `ListInstalledVersions()` - List local versions
- `GetCurrentVersion()` - Determine active version
- `ReadVersionFile()` - Read `.kubernetes-version`
- `DetectClusterVersion()` - Query cluster
- `NormalizeVersion()` - Handle vendor versions

**Version Normalization:**

```go
// Removes vendor suffixes
v1.28.3-eks-123 → v1.28.0
v1.29.5-gke.100 → v1.29.0
v1.27.2+k3s1    → v1.27.0
```

### kubectl Installer (`internal/kubectl/`)

Manages kubectl binaries:

- `Install(version)` - Download and install
- `Uninstall(version)` - Remove version
- `Switch(version)` - Update symlink
- `DownloadBinary(version)` - Fetch from dl.k8s.io

**Download URLs:**

```
https://dl.k8s.io/release/{version}/bin/{os}/{arch}/kubectl
```

### Config Manager (`pkg/config/`)

Application configuration:

- `KuveDir` - `~/.kuve`
- `BinDir` - `~/.kuve/bin`
- `VersionsDir` - `~/.kuve/versions`
- `VersionFile` - `.kubernetes-version`

## Data Flow

### Install kubectl Version

```
kuve install v1.28.0
    ↓
Normalize version
    ↓
Check if installed
    ↓
Download from dl.k8s.io
    ↓
Save to ~/.kuve/versions/v1.28.0/
    ↓
Set executable permissions
    ↓
Confirm installation
```

### Switch kubectl Version

```
kuve switch v1.28.0
    ↓
Verify version installed
    ↓
Remove old symlink
    ↓
Create new symlink:
~/.kuve/bin/kubectl → versions/v1.28.0/kubectl
    ↓
Confirm switch
```

### Use from Version File

```
kuve use
    ↓
Search for .kubernetes-version
    ↓
Read version
    ↓
Install if needed
    ↓
Switch to version
    ↓
Confirm
```

### Use from Cluster

```
kuve use --from-cluster
    ↓
Execute: kubectl version --output=json
    ↓
Parse cluster version
    ↓
Normalize version
    ↓
Install if needed
    ↓
Switch to version
    ↓
Confirm
```

## Technology Stack

### Core Technologies

| Component | Technology | Version |
|-----------|-----------|---------|
| Language | Go | 1.25+ |
| CLI Framework | Cobra | v1.10+ |
| Modules | Go Modules | - |

### Dependencies

```go
require (
    github.com/spf13/cobra v1.10.2
    github.com/spf13/pflag v1.0.10
)
```

Minimal dependencies by design!

## File System Layout

```
~/.kuve/
├── bin/
│   ├── kuve              # Kuve binary
│   └── kubectl           # Symlink → active version
└── versions/
    ├── v1.27.0/
    │   └── kubectl       # kubectl v1.27.0
    ├── v1.28.0/
    │   └── kubectl       # kubectl v1.28.0
    └── v1.29.0/
        └── kubectl       # kubectl v1.29.0
```

## Design Decisions

### Symbolic Links

**Why:** Instant, atomic switching without shell overhead.

**Alternative considered:** Shell functions - rejected due to complexity.

### User-Space Installation

**Why:** No sudo required, safe for multi-user systems.

**Alternative considered:** System-wide - rejected for security.

### PATH-Based Approach

**Why:** Simple, standard Unix approach.

**Alternative considered:** Shell aliases - rejected for portability.

### Version Normalization

**Why:** Support multiple cloud providers seamlessly.

**Implementation:** Remove vendor suffixes, normalize to .0 patch.

### Auto-Install on Use

**Why:** Better user experience.

**Benefit:** No manual install step needed.

## Security Considerations

### Binary Verification

**Current:** Downloads from official dl.k8s.io

**Future:** Checksum verification planned

### User-Space Only

- No system modifications
- No elevated privileges
- Safe for multi-user systems

### File Permissions

```bash
# Binaries executable
-rwxr-xr-x  kubectl

# Directories readable
drwxr-xr-x  versions/
```

## Performance Characteristics

### Operation Times

| Operation | Time | Notes |
|-----------|------|-------|
| Switch version | &lt;10ms | Symlink update |
| Check current | &lt;5ms | Read symlink |
| List installed | &lt;50ms | Directory scan |
| Install | ~5-10s | Network download |
| Uninstall | &lt;100ms | Directory removal |

### Disk Usage

- Per version: ~50MB
- Overhead: Minimal (~1MB for kuve binary)

## Testing Strategy

### Unit Tests

- Configuration management
- Version normalization
- File operations
- Error handling

### Integration Tests

- Command execution
- File system operations
- Network operations (mocked)

### Test Coverage

Target: >80% coverage for critical paths

## Future Enhancements

Planned improvements:

1. **Checksum Verification**: Verify downloaded binaries
2. **Full Version Listing**: List all available kubectl versions
3. **Global Default**: Set default version
4. **Version Constraints**: Support version ranges
5. **Plugin System**: Extend to other Kubernetes tools

## Contributing

See [Contributing](./contributing) for development guidelines.

## Next Steps

- [Development Guide](./development) - Set up dev environment
- [Contributing](./contributing) - Contribution guidelines
- [CLI Reference](../reference/cli) - Command details
