---
sidebar_position: 1
---

# Installation

This guide covers installing Kuve on your system.

## Prerequisites

### Required
- **Operating System**: Linux or macOS
- **Go**: Version 1.25 or later (for building from source)
- **Git**: For cloning the repository

### Optional
- **Make**: For using Makefile commands
- **kubectl**: Will be managed by Kuve after installation

## Installation Methods

### From Source (Recommended)

This is the recommended method for most users:

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

After installation, you need to configure your shell to use Kuve.

### Add to PATH

The `~/.kuve/bin` directory must be in your PATH to use kubectl managed by Kuve:

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

:::warning Important
This needs to be added to your shell configuration file to persist across sessions.
:::

### Configure Your Shell

Choose your shell below and add the configuration:

#### Bash

Add to `~/.bashrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"
```

Apply changes:
```bash
source ~/.bashrc
```

#### Zsh

Add to `~/.zshrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"
```

Apply changes:
```bash
source ~/.zshrc
```

#### Fish

Add to `~/.config/fish/config.fish`:

```fish
# Kuve - Kubernetes Client Switcher
set -gx PATH "$HOME/.kuve/bin" $PATH
```

Apply changes:
```fish
source ~/.config/fish/config.fish
```

## Shell Completion (Optional)

Kuve supports shell completion for enhanced user experience:

### Bash

```bash
# System-wide (requires sudo)
kuve completion bash | sudo tee /etc/bash_completion.d/kuve

# Current session only
source <(kuve completion bash)

# Add to ~/.bashrc for persistence
echo 'source <(kuve completion bash)' >> ~/.bashrc
```

### Zsh

```bash
# Generate completion file
kuve completion zsh > "${fpath[1]}/_kuve"

# Reload completions
autoload -U compinit && compinit
```

### Fish

```bash
# Generate completion file
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

### PowerShell

```powershell
kuve completion powershell | Out-String | Invoke-Expression
```

## Verification

After installation and configuration, verify that Kuve is working correctly:

### Check Kuve Installation

```bash
# Check if kuve is accessible
which kuve
# Expected output: /home/username/.kuve/bin/kuve

# Check version
kuve --version
# Expected output: kuve version dev

# View help
kuve --help
```

### Test Directory Structure

```bash
# Verify Kuve directories were created
ls -la ~/.kuve/
# Expected: bin/ and versions/ directories
```

## Directory Structure

After installation, Kuve creates the following structure:

```
~/.kuve/
├── bin/              # Contains kuve binary and kubectl symlink
│   ├── kuve          # Kuve executable
│   └── kubectl       # Symlink to active kubectl version
└── versions/         # Installed kubectl versions
    ├── v1.28.0/
    │   └── kubectl
    └── v1.29.0/
        └── kubectl
```

## Upgrading Kuve

To upgrade Kuve to a newer version:

```bash
cd kuve
git pull origin main
make install
```

## Troubleshooting

### Command Not Found

If you get `kuve: command not found`:

1. Verify installation:
   ```bash
   ls -l ~/.kuve/bin/kuve
   ```

2. Check PATH:
   ```bash
   echo $PATH | grep .kuve
   ```

3. Add to PATH if missing and restart shell

### Permission Denied

If you get permission errors:

```bash
chmod +x ~/.kuve/bin/kuve
```

### Build Errors

If build fails:

1. Verify Go version:
   ```bash
   go version  # Should be 1.25+
   ```

2. Update dependencies:
   ```bash
   go mod tidy
   go mod download
   ```

## Next Steps

Now that Kuve is installed:

1. [**Quick Start**](./quickstart) - Get up and running in minutes
2. [**Shell Setup**](./shell-setup) - Configure advanced shell features
3. [**Basic Usage**](../user-guide/basic-usage) - Learn the essential commands
