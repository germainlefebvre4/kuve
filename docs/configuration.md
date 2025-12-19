# Configuration

Kuve configuration, directory structure, and customization options.

## Table of Contents

- [Directory Structure](#directory-structure)
- [Configuration Files](#configuration-files)
- [Environment Variables](#environment-variables)
- [Version Files](#version-files)
- [Customization](#customization)
- [Advanced Configuration](#advanced-configuration)

## Directory Structure

Kuve uses a well-defined directory structure in the user's home directory.

### Default Layout

```
$HOME/
└── .kuve/                              # Base directory
    ├── bin/                            # Binary directory
    │   ├── kuve                        # Kuve binary
    │   └── kubectl -> ../versions/v1.28.0/kubectl  # Symlink to active version
    └── versions/                       # Installed kubectl versions
        ├── v1.26.3/
        │   └── kubectl                 # kubectl v1.26.3 binary
        ├── v1.28.0/
        │   └── kubectl                 # kubectl v1.28.0 binary
        └── v1.29.1/
            └── kubectl                 # kubectl v1.29.1 binary
```

### Directory Paths

| Directory | Path | Purpose |
|-----------|------|---------|
| **Base** | `~/.kuve/` | Root directory for all Kuve data |
| **Bin** | `~/.kuve/bin/` | Executable binaries (kuve, kubectl symlink) |
| **Versions** | `~/.kuve/versions/` | Installed kubectl versions |
| **Version Dir** | `~/.kuve/versions/v1.28.0/` | Specific version directory |

### Disk Usage

Each kubectl version uses approximately:
- **Linux/amd64**: ~50 MB
- **macOS/amd64**: ~52 MB
- **macOS/arm64**: ~50 MB

**Example** with 3 versions installed:
```
~/.kuve/                Total: ~155 MB
├── bin/                      ~51 MB (includes kuve + symlink)
└── versions/
    ├── v1.26.3/              ~50 MB
    ├── v1.28.0/              ~52 MB
    └── v1.29.1/              ~52 MB
```

### Directory Creation

Directories are automatically created on first use:

```bash
# First install triggers directory creation
kuve install v1.28.0
# Creates: ~/.kuve/bin/ and ~/.kuve/versions/
```

Manual creation (if needed):

```bash
mkdir -p ~/.kuve/bin
mkdir -p ~/.kuve/versions
```

## Configuration Files

### System-Level Configuration

Kuve currently uses **code-based configuration** defined in `pkg/config/config.go`:

```go
const (
    KuveDir    = ".kuve"
    BinDir     = "bin"
    VersionsDir = "versions"
    KubectlBinary = "kubectl"
    VersionFile = ".kubernetes-version"
)
```

**Note**: There is no user-editable configuration file yet. Configuration is hardcoded.

### Project-Level Configuration

Each project can have a `.kubernetes-version` file:

**Format:**
```
v1.28.0
```

**Location:** Project root directory

**Example:**
```bash
my-project/
├── .kubernetes-version    # Contains: v1.28.0
├── deployments/
│   └── app.yaml
└── README.md
```

**Usage:**
```bash
cd my-project
kuve use  # Automatically uses v1.28.0
```

## Environment Variables

### Used by Kuve

Kuve respects standard environment variables:

| Variable | Usage | Example |
|----------|-------|---------|
| `HOME` | Locates `~/.kuve` directory | `/home/username` |
| `PATH` | Must include `~/.kuve/bin` | `$HOME/.kuve/bin:...` |
| `HTTP_PROXY` | Proxy for downloads | `http://proxy.example.com:8080` |
| `HTTPS_PROXY` | Secure proxy for downloads | `https://proxy.example.com:8443` |

### Setting Environment Variables

#### Bash/Zsh

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$HOME/.kuve/bin:$PATH"
export HTTP_PROXY="http://proxy.example.com:8080"
export HTTPS_PROXY="https://proxy.example.com:8443"
```

#### Fish

```fish
# Add to ~/.config/fish/config.fish
set -gx PATH "$HOME/.kuve/bin" $PATH
set -gx HTTP_PROXY "http://proxy.example.com:8080"
set -gx HTTPS_PROXY "https://proxy.example.com:8443"
```

### Custom Base Directory (Future)

Currently not supported. The base directory is hardcoded to `~/.kuve`.

**Potential future enhancement:**
```bash
export KUVE_HOME="$HOME/.local/share/kuve"
kuve install v1.28.0  # Would install to $KUVE_HOME
```

## Version Files

### .kubernetes-version File

Project-specific version specification.

#### Format

Simple text file containing only the version:

```
v1.28.0
```

Or without 'v' prefix:

```
1.28.0
```

#### Location

- **Current directory**: Checked first
- **Parent directories**: Searched recursively up to `$HOME`

#### Creation

```bash
# Create with specific version
kuve init v1.28.0

# Create with current active version
kuve init
```

#### Version Control

**Recommended**: Commit to version control

```bash
git add .kubernetes-version
git commit -m "Pin kubectl version to v1.28.0"
git push
```

**Benefits:**
- Team uses same kubectl version
- Documents version requirements
- Reproducible environments

#### Multiple Projects

Different projects can use different versions:

```
~/projects/
├── legacy-app/
│   ├── .kubernetes-version    # v1.26.0
│   └── ...
├── current-app/
│   ├── .kubernetes-version    # v1.28.0
│   └── ...
└── next-app/
    ├── .kubernetes-version    # v1.29.0
    └── ...
```

Navigate between projects:
```bash
cd ~/projects/legacy-app
kuve use  # Switches to v1.26.0

cd ~/projects/current-app
kuve use  # Switches to v1.28.0
```

### Version File Hierarchy

Kuve searches from current directory upward:

```
/home/user/projects/team-a/microservice-x/feature-branch/
├── .kubernetes-version (v1.29.0) ← Used (most specific)
└── ...

/home/user/projects/team-a/microservice-x/
├── .kubernetes-version (v1.28.0) ← Ignored (not used)
└── ...

/home/user/projects/team-a/
├── .kubernetes-version (v1.27.0) ← Ignored (not used)
└── ...

/home/user/projects/
└── .kubernetes-version (v1.26.0)  ← Ignored (not used)
```

**Rule**: First `.kubernetes-version` file found is used.

## Customization

### Shell Integration

Customize shell behavior with auto-switching.

#### Basic PATH Setup

**Bash** (`~/.bashrc`):
```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

**Zsh** (`~/.zshrc`):
```zsh
export PATH="$HOME/.kuve/bin:$PATH"
```

**Fish** (`~/.config/fish/config.fish`):
```fish
set -gx PATH "$HOME/.kuve/bin" $PATH
```

#### Auto-Switching

Automatically switch kubectl version when entering a directory.

**Bash** (`~/.bashrc`):
```bash
# Kuve auto-switch
kuve_auto_switch() {
    if [ -f .kubernetes-version ]; then
        local version=$(cat .kubernetes-version | tr -d '[:space:]')
        if [ -n "$version" ]; then
            kuve use 2>/dev/null || echo "Failed to switch to kubectl $version"
        fi
    fi
}

# Hook into cd
cd() {
    builtin cd "$@" && kuve_auto_switch
}

# Run on shell startup
kuve_auto_switch
```

**Zsh** (`~/.zshrc`):
```zsh
# Kuve auto-switch
kuve_auto_switch() {
    if [ -f .kubernetes-version ]; then
        local version=$(cat .kubernetes-version | tr -d '[:space:]')
        if [ -n "$version" ]; then
            kuve use 2>/dev/null || echo "Failed to switch to kubectl $version"
        fi
    fi
}

# Use chpwd hook
autoload -U add-zsh-hook
add-zsh-hook chpwd kuve_auto_switch

# Run on shell startup
kuve_auto_switch
```

**Fish** (`~/.config/fish/config.fish`):
```fish
# Kuve auto-switch
function kuve_auto_switch --on-variable PWD
    if test -f .kubernetes-version
        set version (cat .kubernetes-version | tr -d '[:space:]')
        if test -n "$version"
            kuve use 2>/dev/null; or echo "Failed to switch to kubectl $version"
        end
    end
end

# Run on shell startup
kuve_auto_switch
```

### Prompt Customization

Show current kubectl version in shell prompt.

**Bash** (`~/.bashrc`):
```bash
# Add to PS1
kuve_prompt() {
    local version=$(kuve current 2>/dev/null | awk '{print $NF}')
    if [ -n "$version" ]; then
        echo " [kubectl:$version]"
    fi
}

export PS1='\u@\h:\w$(kuve_prompt)\$ '
```

**Zsh** (`~/.zshrc`):
```zsh
# Add to prompt
kuve_prompt() {
    local version=$(kuve current 2>/dev/null | awk '{print $NF}')
    if [ -n "$version" ]; then
        echo " [kubectl:$version]"
    fi
}

RPROMPT='$(kuve_prompt)'
```

**Result:**
```
user@host:~/project [kubectl:v1.28.0]$
```

### Aliases

Convenient shortcuts for common operations.

**Bash/Zsh** (`~/.bashrc` or `~/.zshrc`):
```bash
# Kuve aliases
alias kv='kuve'
alias kvi='kuve install'
alias kvs='kuve switch'
alias kvl='kuve list installed'
alias kvc='kuve current'
alias kvu='kuve use'

# Combined operations
alias kvup='kuve use && kubectl cluster-info'
alias kvcluster='kuve use --from-cluster'
```

**Fish** (`~/.config/fish/config.fish`):
```fish
# Kuve aliases
abbr kv 'kuve'
abbr kvi 'kuve install'
abbr kvs 'kuve switch'
abbr kvl 'kuve list installed'
abbr kvc 'kuve current'
abbr kvu 'kuve use'

# Combined operations
abbr kvup 'kuve use && kubectl cluster-info'
abbr kvcluster 'kuve use --from-cluster'
```

## Advanced Configuration

### Custom Download Mirror (Future)

Not currently supported. Downloads are from `dl.k8s.io` only.

**Potential future configuration:**
```yaml
# ~/.kuve/config.yaml (future)
download:
  mirror: "https://custom-mirror.example.com/kubectl"
  timeout: 300
  retries: 3
```

### Version Aliases (Future)

Not currently supported.

**Potential future configuration:**
```yaml
# ~/.kuve/config.yaml (future)
aliases:
  stable: v1.28.0
  latest: v1.29.1
  legacy: v1.26.3
```

Usage:
```bash
kuve switch stable  # Switches to v1.28.0
```

### Auto-Cleanup (Future)

Not currently supported.

**Potential future configuration:**
```yaml
# ~/.kuve/config.yaml (future)
cleanup:
  enabled: true
  keep_versions: 3  # Keep only 3 most recent versions
  keep_current: true  # Always keep current version
```

### Verification Settings (Future)

Not currently supported. No checksum verification yet.

**Potential future configuration:**
```yaml
# ~/.kuve/config.yaml (future)
security:
  verify_checksum: true
  verify_signature: false
  allow_unsigned: true
```

## Configuration Best Practices

### 1. Version Control

Always commit `.kubernetes-version` files:

```bash
# .gitignore - Do NOT ignore version files
# .kubernetes-version should be committed

git add .kubernetes-version
git commit -m "Pin kubectl version"
```

### 2. Team Standardization

Document kubectl version in README:

```markdown
## Prerequisites

- kubectl v1.28.0 (managed by Kuve)

### Setup

1. Install Kuve: https://github.com/germainlefebvre4/kuve
2. Install kubectl: `kuve use`
```

### 3. Multiple Environments

Use different version files for different branches:

```bash
# main branch
echo "v1.28.0" > .kubernetes-version

# development branch
git checkout develop
echo "v1.29.0" > .kubernetes-version
```

### 4. PATH Priority

Ensure Kuve's kubectl is used:

```bash
# Add at END of shell config to override other PATH entries
export PATH="$HOME/.kuve/bin:$PATH"
```

### 5. Regular Cleanup

Remove unused versions:

```bash
# List installed versions
kuve list installed

# Remove old versions
kuve uninstall v1.25.0
kuve uninstall v1.26.0
```

## Migration from Other Tools

### From kubectl Version Manager (kvm)

Similar directory structure:

```bash
# Backup existing installations
cp -r ~/.kvm ~/.kvm.backup

# Install kuve
make install

# Recreate installations with kuve
kuve install v1.28.0
kuve switch v1.28.0
```

### From Manual kubectl Installation

```bash
# Note current version
kubectl version --client

# Install with kuve
kuve install v1.28.0
kuve switch v1.28.0

# Remove old kubectl (optional)
sudo rm /usr/local/bin/kubectl
```

### From asdf-kubectl

```bash
# Note installed versions
asdf list kubectl

# Install with kuve
kuve install v1.28.0
kuve install v1.29.0

# Switch with kuve
kuve switch v1.28.0

# Optionally remove asdf plugin
asdf plugin remove kubectl
```

## Related Documentation

- [Installation Guide](./installation.md) - Setup instructions
- [Usage Guide](./usage.md) - How to use Kuve
- [Shell Integration Guide](../SHELL_INTEGRATION.md) - Auto-switching setup
- [Troubleshooting](./troubleshooting.md) - Common issues
