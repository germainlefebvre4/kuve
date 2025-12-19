# Installation Guide

This guide covers all the methods to install Kuve on your system.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation Methods](#installation-methods)
  - [From Source](#from-source)
  - [Using Make](#using-make)
  - [Manual Build](#manual-build)
- [Post-Installation Setup](#post-installation-setup)
- [Shell Configuration](#shell-configuration)
- [Verification](#verification)

## Prerequisites

### Required

- **Operating System**: Linux or macOS
- **Go**: Version 1.25 or later (for building from source)
- **Git**: For cloning the repository

### Optional

- **Make**: For using Makefile commands
- **kubectl**: Will be managed by Kuve after installation

## Installation Methods

### From Source

This is the recommended method for most users.

```bash
# Clone the repository
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve

# Build and install using make
make install
```

This will:
1. Build the `kuve` binary
2. Copy it to `~/.kuve/bin/kuve`
3. Create necessary directory structure

### Using Make

If you have already cloned the repository:

```bash
cd kuve

# Build only (creates binary in current directory)
make build

# Build and install
make install

# Clean build artifacts
make clean
```

### Manual Build

If you prefer to build manually without Make:

```bash
# Build the binary
go build -o kuve main.go

# Install to a directory in your PATH
sudo mv kuve /usr/local/bin/

# Or install to user directory
mkdir -p ~/.kuve/bin
mv kuve ~/.kuve/bin/
```

## Post-Installation Setup

After installation, you need to add Kuve's bin directory to your PATH.

### Add to PATH

The `~/.kuve/bin` directory must be in your PATH to use kubectl managed by Kuve:

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

**Important**: This needs to be added to your shell configuration file to persist across sessions.

## Shell Configuration

Configure your shell to automatically set up Kuve on startup.

### Bash

Add to `~/.bashrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"
```

Apply changes:
```bash
source ~/.bashrc
```

### Zsh

Add to `~/.zshrc`:

```zsh
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"
```

Apply changes:
```zsh
source ~/.zshrc
```

### Fish

Add to `~/.config/fish/config.fish`:

```fish
# Kuve - Kubernetes Client Switcher
set -gx PATH "$HOME/.kuve/bin" $PATH
```

Apply changes:
```fish
source ~/.config/fish/config.fish
```

### Shell Completion (Optional)

Kuve supports shell completion for enhanced user experience:

#### Bash

```bash
# System-wide (requires sudo)
kuve completion bash | sudo tee /etc/bash_completion.d/kuve

# Current session only
source <(kuve completion bash)

# Add to ~/.bashrc for persistence
echo 'source <(kuve completion bash)' >> ~/.bashrc
```

#### Zsh

```zsh
# Generate completion file
kuve completion zsh > "${fpath[1]}/_kuve"

# Reload completions
autoload -U compinit && compinit
```

#### Fish

```fish
# Generate completion file
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

#### PowerShell

```powershell
kuve completion powershell | Out-String | Invoke-Expression
```

## Verification

After installation and configuration, verify that Kuve is working correctly:

### Check Kuve Installation

```bash
# Check if kuve is accessible
which kuve
# Expected output: /home/username/.kuve/bin/kuve (or similar)

# Check version
kuve --version
# Expected output: kuve version dev (or specific version)

# View help
kuve --help
```

### Test Basic Functionality

```bash
# List remote versions (should show latest stable version)
kuve list remote

# List installed versions (should be empty initially)
kuve list installed
```

### Install Your First kubectl Version

```bash
# Install a kubectl version
kuve install v1.29.1

# Verify installation
kuve list installed

# Switch to it
kuve switch v1.29.1

# Verify kubectl works
kubectl version --client
```

## Directory Structure

After installation, Kuve creates the following directory structure:

```
~/.kuve/
├── bin/
│   ├── kuve              # Kuve binary
│   └── kubectl           # Symlink to active kubectl version
└── versions/
    └── v1.29.1/          # Example installed version
        └── kubectl       # kubectl binary
```

## Troubleshooting Installation

### kuve: command not found

**Problem**: The shell cannot find the `kuve` command.

**Solution**: Ensure `~/.kuve/bin` is in your PATH and you've reloaded your shell configuration:
```bash
echo $PATH | grep -q ".kuve/bin" && echo "Found" || echo "Not found"
source ~/.bashrc  # or ~/.zshrc for Zsh
```

### Permission Denied

**Problem**: Cannot execute `kuve` or `kubectl`.

**Solution**: Make sure the binaries have execute permissions:
```bash
chmod +x ~/.kuve/bin/kuve
```

### Go Version Mismatch

**Problem**: Build fails due to Go version.

**Solution**: Install Go 1.25 or later:
```bash
# Check Go version
go version

# Update Go if needed (Linux)
wget https://go.dev/dl/go1.25.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.25.0.linux-amd64.tar.gz
```

## Updating Kuve

To update Kuve to the latest version:

```bash
cd kuve
git pull origin main
make install
```

## Uninstalling Kuve

To completely remove Kuve from your system:

```bash
# Remove Kuve directory
rm -rf ~/.kuve

# Remove PATH entry from shell configuration
# Edit ~/.bashrc, ~/.zshrc, or ~/.config/fish/config.fish
# Remove the line: export PATH="$HOME/.kuve/bin:$PATH"
```

## Next Steps

After successful installation:

1. Read the [Quick Start Guide](../QUICKSTART.md)
2. Learn about [Usage](./usage.md)
3. Set up [Shell Integration](../SHELL_INTEGRATION.md) for auto-switching
4. Explore [CLI Reference](./cli-reference.md) for all commands
