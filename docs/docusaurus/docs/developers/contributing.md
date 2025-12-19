---
sidebar_position: 2
---

# Contributing

Thank you for considering contributing to Kuve! This guide will help you get started.

## Code of Conduct

- Be respectful and constructive
- Welcome newcomers and help them learn
- Focus on what is best for the community
- Show empathy towards other community members

## How to Contribute

### Reporting Bugs

Before creating a bug report:

1. **Check existing issues**: Search for similar issues
2. **Verify it's a bug**: Not expected behavior
3. **Test latest version**: Bug might be fixed

**Create a bug report with:**

- **Clear title**: Describe the issue
- **Steps to reproduce**: Exact steps to trigger the bug
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Environment**:
  - Kuve version: `kuve --version`
  - OS: `uname -a`
  - Go version: `go version` (if building)
- **Logs**: Include error messages (use `--verbose`)

**Example:**

```markdown
## Bug: Cannot uninstall version after cluster detection

**Steps to reproduce:**
1. Run `kuve use --from-cluster`
2. Try `kuve uninstall v1.28.0`

**Expected:** Version should uninstall
**Actual:** Error: "cannot uninstall active version"

**Environment:**
- Kuve version: dev
- OS: Linux 5.15.0
- Go version: 1.25.0
```

### Suggesting Features

When suggesting features:

1. **Check existing requests**: Avoid duplicates
2. **Explain the use case**: Why is this needed?
3. **Describe the solution**: How should it work?
4. **Consider alternatives**: Other approaches?

**Template:**

```markdown
## Feature: Support for version constraints

**Use Case:**
As a developer, I want to specify version ranges (e.g., ">=1.28.0, &lt;1.30.0") 
in .kubernetes-version files to ensure compatibility.

**Proposed Solution:**
Allow syntax like:
```
>=1.28.0, \<1.30.0
```

**Alternatives:**
- Exact version only (current approach)
- Minimum version only
```

### Contributing Code

#### Development Setup

```bash
# Fork the repository on GitHub
# Then clone your fork
git clone https://github.com/YOUR_USERNAME/kuve.git
cd kuve

# Add upstream remote
git remote add upstream https://github.com/germainlefebvre4/kuve.git

# Install dependencies
go mod download

# Build
make build

# Run tests
make test
```

#### Create a Branch

```bash
# Update your fork
git fetch upstream
git checkout main
git merge upstream/main

# Create feature branch
git checkout -b feature/your-feature-name
```

**Branch naming:**

- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation
- `test/` - Test improvements
- `refactor/` - Code refactoring

#### Make Changes

Follow the [Coding Standards](#coding-standards) below.

#### Test Your Changes

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Test manually
make install
kuve --version
```

#### Commit Changes

**Commit message format:**

```
<type>: <short summary>

<optional body>

<optional footer>
```

**Types:**

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation
- `test`: Tests
- `refactor`: Code refactoring
- `style`: Formatting, no code change
- `chore`: Maintenance

**Examples:**

```bash
git commit -m "feat: add support for version constraints"

git commit -m "fix: prevent race condition in version switching

- Added mutex to protect symlink updates
- Added test case for concurrent switches
- Fixes #123"

git commit -m "docs: update installation instructions for macOS"
```

#### Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name
```

Then create a Pull Request on GitHub.

### Pull Request Guidelines

**PR Title:** Same format as commit messages

**PR Description should include:**

- **What**: What does this PR do?
- **Why**: Why is this change needed?
- **How**: How does it work?
- **Testing**: How was it tested?
- **Related Issues**: Fixes #123, Closes #456

**Example:**

```markdown
## Add version constraint support

**What:**
Adds support for version ranges in .kubernetes-version files.

**Why:**
Users need flexibility to specify compatible version ranges rather than exact versions.

**How:**
- Added constraint parsing to version manager
- Implemented semver-style range matching
- Updated `kuve use` to resolve constraints

**Testing:**
- Added unit tests for constraint parsing
- Tested with various range formats
- Manual testing with real .kubernetes-version files

**Related Issues:**
- Fixes #123
- Addresses feature request in #45
```

**Checklist before submitting:**

- [ ] Tests pass locally
- [ ] Code follows style guide
- [ ] Documentation updated (if needed)
- [ ] Commit messages are clear
- [ ] Branch is up to date with main
- [ ] PR description is complete

## Coding Standards

### Go Best Practices

Follow standard Go conventions:

```go
// Good: Exported function with doc comment
// InstallVersion downloads and installs a specific kubectl version.
// It returns an error if the download fails or version already exists.
func InstallVersion(version string) error {
    // Implementation
}

// Bad: No doc comment, unclear name
func iv(v string) error {
    // Implementation
}
```

### Style Guidelines

1. **Formatting**: Use `go fmt`
   ```bash
   make fmt
   ```

2. **Naming**:
   - Use clear, descriptive names
   - Follow Go naming conventions
   - Exported names start with uppercase
   - Unexported names start with lowercase

3. **Error Handling**:
   ```go
   // Good: Descriptive error messages
   if err != nil {
       return fmt.Errorf("failed to download kubectl v%s: %w", version, err)
   }
   
   // Bad: Generic error
   if err != nil {
       return err
   }
   ```

4. **Functions**:
   - Keep functions small and focused
   - Single responsibility principle
   - Maximum ~50 lines per function

5. **Comments**:
   ```go
   // Good: Explain why, not what
   // Normalize removes vendor suffixes to match official kubectl releases.
   // GKE, EKS, and AKS all add suffixes that don't exist in dl.k8s.io.
   
   // Bad: Redundant comment
   // This function normalizes the version
   ```

### Testing Standards

Write tests for:
- New features
- Bug fixes
- Edge cases

```go
func TestNormalizeVersion(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {
            name:     "GKE version",
            input:    "v1.28.3-gke.1234",
            expected: "v1.28.0",
        },
        {
            name:     "EKS version",
            input:    "v1.27.5-eks-abc",
            expected: "v1.27.0",
        },
        // More test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := NormalizeVersion(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### Documentation Standards

Update documentation when:
- Adding new features
- Changing behavior
- Adding commands/flags
- Fixing bugs (if user-facing)

**Documentation files:**
- `README.md` - Main documentation
- `docs/` - Detailed guides
- Code comments - Exported functions
- CLI help - Command descriptions

## Development Workflow

### Local Testing

```bash
# Build
make build

# Install locally
make install

# Test manually
kuve list installed
kuve install v1.28.0
kuve switch v1.28.0
```

### Running Tests

```bash
# All tests
make test

# Specific package
go test ./pkg/config

# With coverage
make test-coverage
```

### Code Quality

```bash
# Format code
make fmt

# Run linter
make lint

# Check for issues
go vet ./...
```

## Project-Specific Guidelines

### Adding a New Command

1. Create command file in `cmd/`:
   ```go
   // cmd/newcommand.go
   package cmd
   
   func init() {
       rootCmd.AddCommand(newCmd)
   }
   
   var newCmd = &cobra.Command{
       Use:   "new",
       Short: "Short description",
       Long:  `Long description`,
       Run: func(cmd *cobra.Command, args []string) {
           // Implementation
       },
   }
   ```

2. Add tests

3. Update documentation:
   - `README.md`
   - `docs/cli-reference.md`
   - CLI help text

### Modifying Version Logic

When changing version handling:

1. Update `internal/version/manager.go`
2. Add tests in `manager_test.go`
3. Consider backward compatibility
4. Update normalization docs if needed

### Adding Configuration

For new config options:

1. Update `pkg/config/config.go`
2. Add constants/structs
3. Update initialization
4. Document in `docs/configuration.md`

## Getting Help

- **Questions**: [GitHub Discussions](https://github.com/germainlefebvre4/kuve/discussions)
- **Issues**: [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues)
- **Development**: See [Development Guide](./development)

## Recognition

Contributors are recognized in:
- GitHub contributors page
- Release notes
- Project README

Thank you for contributing to Kuve! 🎉

## Next Steps

- [Development Guide](./development) - Set up environment
- [Architecture](./architecture) - Understand the system
- [GitHub Issues](https://github.com/germainlefebvre4/kuve/issues) - Find issues to work on
