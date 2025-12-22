---
sidebar_position: 3
---

# Development Guide

Complete guide for setting up your development environment and working on Kuve.

## Prerequisites

### Required Software

| Tool | Version | Purpose |
|------|---------|---------|
| **Go** | 1.25+ | Development language |
| **Git** | Any | Version control |
| **Make** | Any | Build automation |

### Optional Tools

| Tool | Purpose |
|------|---------|
| **golangci-lint** | Code linting |
| **gopls** | Language server (for IDE) |
| **delve** | Debugging |

## Setup Development Environment

### Install Go

- Install Go for your OS from the [official site](https://go.dev/dl/)
- Or use a version manager like `gvm` or `asdf`

### Fork and Clone

```bash
# Fork on GitHub first
# Then clone your fork
git clone https://github.com/YOUR_USERNAME/kuve.git
cd kuve

# Add upstream remote
git remote add upstream https://github.com/germainlefebvre4/kuve.git

# Verify remotes
git remote -v
```

### Install Dependencies

```bash
# Download Go modules
go mod download

# Verify
go mod verify
```

### Build Project

```bash
# Build binary
make build

# Verify build
./kuve --version
```

### Install Development Tools

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install delve (debugger)
go install github.com/go-delve/delve/cmd/dlv@latest

# Verify
golangci-lint --version
dlv version
```

## Development Workflow

### Daily Workflow

```bash
# 1. Update your fork
git fetch upstream
git checkout main
git merge upstream/main

# 2. Create feature branch
git checkout -b feature/my-feature

# 3. Make changes
# Edit files...

# 4. Test locally
make test

# 5. Build and install
make install

# 6. Manual testing
kuve list installed
kuve install v1.28.0

# 7. Commit changes
git add .
git commit -m "feat: add my feature"

# 8. Push to fork
git push origin feature/my-feature

# 9. Create Pull Request on GitHub
```

### Making Changes

#### Edit Code

```bash
# Use your favorite editor
code .              # VS Code
vim internal/version/manager.go
```

#### Format Code

```bash
# Auto-format all Go files
make fmt

# Or manually
go fmt ./...
```

#### Run Tests

```bash
# All tests
make test

# Specific package
go test ./internal/version

# With verbose output
go test -v ./...

# With coverage
make test-coverage
```

#### Run Linter

```bash
# Using make
make lint

# Or directly
golangci-lint run
```

## Project Structure

### Key Directories

```
kuve/
├── cmd/              # CLI commands
├── internal/         # Internal packages
│   ├── kubectl/      # kubectl management
│   └── version/      # Version operations
├── pkg/              # Public packages
│   └── config/       # Configuration
├── docs/             # Documentation
├── .github/          # GitHub Actions
├── main.go           # Entry point
├── Makefile          # Build automation
└── go.mod            # Go modules
```

### Adding New Code

#### New Command

```bash
# Create new command file
touch cmd/newcommand.go
```

```go
package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
    Use:   "new [args]",
    Short: "Short description",
    Long: `Long description of what the command does.`,
    Args: cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
        fmt.Printf("New command with arg: %s\n", args[0])
    },
}
```

#### New Internal Package

```bash
# Create package directory
mkdir -p internal/mypackage

# Create Go file
touch internal/mypackage/myfile.go
```

```go
package mypackage

// MyFunction does something useful.
func MyFunction(input string) (string, error) {
    // Implementation
    return result, nil
}
```

#### Tests

```bash
# Create test file
touch internal/mypackage/myfile_test.go
```

```go
package mypackage

import "testing"

func TestMyFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    "test",
            expected: "result",
            wantErr:  false,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := MyFunction(tt.input)
            
            if (err != nil) != tt.wantErr {
                t.Errorf("MyFunction() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            
            if result != tt.expected {
                t.Errorf("MyFunction() = %v, want %v", result, tt.expected)
            }
        })
    }
}
```

## Build System

### Makefile Targets

```bash
# Build binary
make build

# Install to system
make install

# Run tests
make test

# Test with coverage
make test-coverage

# Format code
make fmt

# Run linter
make lint

# Clean build artifacts
make clean

# Build for all platforms
make build-all

# Download dependencies
make deps

# Show help
make help
```

### Manual Commands

```bash
# Build
go build -o kuve main.go

# Run tests
go test ./...

# Install
go install

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o kuve-linux-amd64

# Cross-compile for macOS
GOOS=darwin GOARCH=amd64 go build -o kuve-darwin-amd64
```

## Testing

### Unit Tests

```bash
# Run all tests
go test ./...

# Run specific package
go test ./internal/version

# Run specific test
go test -run TestNormalizeVersion ./internal/version

# Verbose output
go test -v ./...

# With coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Integration Tests

```bash
# Build and install
make install

# Test real usage
kuve install v1.28.0
kuve switch v1.28.0
kuve current
kuve list installed
kuve uninstall v1.28.0
```

### Test Coverage

```bash
# Generate coverage
make test-coverage

# View in browser
go tool cover -html=coverage.out
```

## Debugging

### Using Delve

```bash
# Start debugger
dlv debug

# Debug specific test
dlv test ./internal/version -- -test.run TestNormalizeVersion

# Breakpoint commands
(dlv) break main.main
(dlv) continue
(dlv) next
(dlv) step
(dlv) print variable
(dlv) quit
```

### Print Debugging

```go
import "fmt"

func MyFunction() {
    fmt.Printf("Debug: value = %+v\n", value)
}
```

### Logging

```go
import "log"

func MyFunction() {
    log.Printf("Processing version: %s", version)
}
```

## Common Development Tasks

### Add a New Flag

```go
var myFlag string

func init() {
    myCmd.Flags().StringVarP(&myFlag, "my-flag", "m", "", "Description")
}
```

### Add Configuration Option

```go
// In pkg/config/config.go
const NewOption = "value"

// In NewConfig()
config.NewOption = "default"
```

### Modify Version Detection

```go
// In internal/version/manager.go
func (vm *VersionManager) DetectVersion() (string, error) {
    // Your implementation
}
```

## Best Practices

### Code Organization

- Keep functions small (&lt;50 lines)
- Single responsibility per function
- Clear, descriptive names
- Add comments for exported functions

### Error Handling

```go
// Good: Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to download kubectl v%s: %w", version, err)
}

// Bad: Return raw error
if err != nil {
    return err
}
```

### Testing

- Test happy path
- Test error conditions
- Test edge cases
- Use table-driven tests

### Documentation

- Doc comments for exported items
- Update README for new features
- Update CLI help text
- Add examples in docs/

## Troubleshooting Development Issues

### Build Fails

```bash
# Clean and rebuild
make clean
go mod tidy
make build
```

### Tests Fail

```bash
# Check test output
go test -v ./...

# Run specific failing test
go test -v -run TestName ./package
```

### Import Issues

```bash
# Update dependencies
go mod tidy
go mod download
```

### IDE Issues

```bash
# Regenerate gopls cache
rm -rf ~/.cache/gopls
```

## Release Process

(For maintainers)

```bash
# 1. Update version
git tag v0.2.0

# 2. Push tag
git push origin v0.2.0

# 3. GitHub Actions will:
#    - Build binaries
#    - Create release
#    - Upload artifacts
```

## Resources

### Go Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go by Example](https://gobyexample.com/)

### Cobra Resources

- [Cobra Documentation](https://github.com/spf13/cobra)
- [Cobra User Guide](https://github.com/spf13/cobra/blob/main/user_guide.md)

### Testing Resources

- [Go Testing](https://go.dev/doc/tutorial/add-a-test)
- [Table-Driven Tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)

## Getting Help

- **Questions**: [GitHub Discussions](https://github.com/germainlefebvre4/kuve/discussions)
- **Issues**: [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues)
- **Contributing**: See [Contributing Guide](./contributing)

## Next Steps

- [Contributing](./contributing) - Contribution guidelines
- [Architecture](./architecture) - System architecture
- [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) - Find issues to work on
