# Development Summary

## Project Overview

**Kuve** (Kubernetes Client Switcher) is a CLI tool written in Go that allows users to easily manage and switch between multiple kubectl versions.

## Implementation Details

### Technology Stack
- **Language**: Go 1.25
- **CLI Framework**: Cobra
- **Testing**: Go standard testing library
- **Build**: Makefile + Go modules

### Project Structure

```
kuve/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command and app initialization
│   ├── install.go         # Install kubectl versions
│   ├── uninstall.go       # Uninstall kubectl versions
│   ├── switch.go          # Switch between versions
│   ├── list.go            # List versions (remote/installed)
│   └── use.go             # Use version from .kubernetes-version file
├── internal/              # Internal packages
│   ├── kubectl/           # kubectl management
│   │   └── installer.go   # Download, install, uninstall logic
│   └── version/           # Version management
│       ├── manager.go     # Version operations
│       └── manager_test.go # Unit tests
├── pkg/                   # Public packages
│   └── config/            # Configuration
│       ├── config.go      # App configuration
│       └── config_test.go # Unit tests
├── .github/               # GitHub workflows
│   ├── workflows/         # CI/CD pipelines
│   │   ├── ci.yml        # Continuous integration
│   │   └── release.yml   # Release automation
│   └── instructions/      # Development instructions
├── main.go               # Entry point
├── Makefile              # Build automation
└── [documentation files]
```

### Core Components

#### 1. Configuration (`pkg/config`)
- Manages application directories (`~/.kuve/`)
- Defines constants (binary names, version file name)
- Creates necessary directory structure

#### 2. Version Manager (`internal/version`)
- Fetches stable kubectl version from Kubernetes releases
- Lists installed versions
- Reads `.kubernetes-version` files
- Searches directory tree for version files
- Detects Kubernetes cluster version from current context
- Executes kubectl to query server version

#### 3. Kubectl Installer (`internal/kubectl`)
- Downloads kubectl binaries from official sources
- Installs to versioned directories
- Creates/updates symbolic links
- Validates installations
- Handles uninstallation with safety checks

#### 4. CLI Commands (`cmd/`)
- **root**: Base command with version info
- **install**: Download and install specific version
- **uninstall**: Remove installed version (with safeguards)
- **switch**: Change active version via symlink
- **list installed**: Show all installed versions
- **list remote**: Show available remote versions
- **use**: Auto-switch using `.kubernetes-version` file or cluster detection
- **init**: Create `.kubernetes-version` file
- **current**: Display active version

### Key Features Implemented

✅ **Install kubectl versions**
- Downloads from `dl.k8s.io`
- Supports version normalization (with/without 'v' prefix)
- Stores in `~/.kuve/versions/<version>/`

✅ **Switch versions**
- Uses symbolic links for instant switching
- Points `~/.kuve/bin/kubectl` to active version
- No PATH manipulation required after setup

✅ **List versions**
- Shows installed versions with current marker (*)
- Fetches latest stable version from official source
- Sorted output

✅ **Uninstall versions**
- Safety check: prevents uninstalling active version
- Cleans up version directory completely

✅ **Version files**
- `.kubernetes-version` file support
- Auto-install missing versions when using version file
- Searches parent directories for version file

✅ **Cluster detection**
- Detects Kubernetes version from current cluster context
- Uses `kubectl version` with JSON output
- Fallback to short output format
- Auto-installs matching kubectl version
- Works with `kuve use --from-cluster` flag

✅ **Shell integration**
- PATH-based approach
- Optional auto-switching on directory change
- Examples for Bash, Zsh, and Fish

### Testing

- **Unit tests** for critical components:
  - Configuration management
  - Version file reading
  - Installed version detection
  - Directory operations

- **Test coverage**: All tests passing
  - `pkg/config`: Configuration logic
  - `internal/version`: Version management

### Build System

**Makefile targets:**
- `build`: Build binary
- `install`: Build and install to system
- `test`: Run unit tests
- `test-coverage`: Generate coverage report
- `clean`: Remove build artifacts
- `fmt`: Format code
- `lint`: Run linter (requires golangci-lint)
- `deps`: Download dependencies
- `build-all`: Multi-platform builds
- `help`: Show available targets

### CI/CD

**GitHub Actions workflows:**
1. **CI Pipeline** (`.github/workflows/ci.yml`)
   - Runs on push and PR
   - Tests on multiple Go versions (1.23, 1.24, 1.25)
   - Runs tests, builds, and linting
   - Uploads coverage to Codecov

2. **Release Pipeline** (`.github/workflows/release.yml`)
   - Triggers on version tags (v*)
   - Builds for multiple platforms (Linux/macOS, amd64/arm64)
   - Creates GitHub releases with binaries
   - Generates checksums

### Documentation

- **README.md**: Main documentation with usage examples
- **QUICKSTART.md**: Getting started guide
- **CONTRIBUTING.md**: Contribution guidelines
- **SHELL_INTEGRATION.md**: Shell setup and auto-switching
- **LICENSE**: MIT License
- **DEVELOPMENT.md**: This file

### Code Statistics

- **Total Go code**: ~1020 lines
- **Number of commands**: 8 main commands
- **Test files**: 2 test files with multiple test cases
- **External dependencies**: Minimal (Cobra CLI only)

### Design Decisions

1. **Symlink-based switching**: Fast, atomic, no shell function overhead
2. **Self-contained directory**: All data in `~/.kuve/`
3. **Version normalization**: Accept both `v1.28.0` and `1.28.0`
4. **Safety checks**: Prevent uninstalling active version
5. **Auto-install on use**: Install missing versions automatically when using version file
6. **Progressive directory search**: Find `.kubernetes-version` in parent directories

### Future Enhancements

Potential improvements (not implemented):
- Full remote version listing (requires GitHub API integration)
- Checksum verification for downloaded binaries
- Global default version configuration
- Version constraints/ranges in version files
- Plugin architecture for other Kubernetes tools
- Shell hooks for auto-switching without manual setup
- Version update notifications

## Getting Started with Development

```bash
# Clone repository
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve

# Download dependencies
make deps

# Run tests
make test

# Build
make build

# Try it out
./kuve --help
```

## Code Quality

- ✅ All tests passing
- ✅ No `go vet` warnings
- ✅ Code formatted with `go fmt`
- ✅ Follows Go best practices
- ✅ Error handling throughout
- ✅ Clear separation of concerns
- ✅ Meaningful variable/function names

## Maintainability

- **Modular structure**: Clear package organization
- **Testable code**: Logic separated from CLI
- **Documentation**: Inline comments and external docs
- **Type safety**: Leverages Go's type system
- **Error propagation**: Consistent error handling pattern

---

Built with ❤️ using Go and Cobra
