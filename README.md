# Kuve - Kubernetes Client Switcher

Kuve is a CLI tool to easily manage and switch between multiple kubectl versions on your system.

## ⚡ Quick Start

```bash
# Build and install
make install

# Add to PATH
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Install a kubectl version
kuve install v1.29.1

# Switch to it
kuve switch v1.29.1
```

See [QUICKSTART.md](QUICKSTART.md) for a detailed getting started guide.

## Features

- 📦 **Install** specific kubectl versions
- 🔄 **Switch** between installed kubectl versions
- 📋 **List** available and installed versions
- 🗑️ **Uninstall** kubectl versions you no longer need
- 📄 **Version files** - Use `.kubernetes-version` file for project-specific versions
- 🔍 **Cluster detection** - Auto-detect Kubernetes version from current cluster context
- 🚀 **Automatic installation** - Install versions on-the-fly when switching

## Installation

### From Source

```bash
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve
go build -o kuve main.go
sudo mv kuve /usr/local/bin/
```

### Prerequisites

- Go 1.25 or later (for building from source)
- Linux or macOS

## Usage

### Setup

After installation, make sure `~/.kuve/bin` is in your PATH:

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

Add this line to your `~/.bashrc`, `~/.zshrc`, or equivalent shell configuration file.

### Shell Completion (Optional)

Kuve supports shell completion for Bash, Zsh, Fish, and PowerShell:

```bash
# Bash
kuve completion bash > /etc/bash_completion.d/kuve  # System-wide
# or
source <(kuve completion bash)  # Current session

# Zsh
kuve completion zsh > "${fpath[1]}/_kuve"

# Fish
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

### Commands

#### Install a kubectl version

```bash
kuve install v1.28.0
# or without the 'v' prefix
kuve install 1.28.0
```

#### List installed versions

```bash
kuve list installed
```

Output:
```
Installed kubectl versions:
  v1.26.3
* v1.28.0
  v1.29.1

* = current version (v1.28.0)
```

#### List available remote versions

```bash
kuve list remote
```

#### Switch to a different version

```bash
kuve switch v1.28.0
# or use the alias
kuve use v1.28.0
```

#### Show current version

```bash
kuve current
```

#### Uninstall a version

```bash
kuve uninstall v1.28.0
```

**Note:** You cannot uninstall the currently active version. Switch to another version first.

### Working with .kubernetes-version files

#### Create a version file

Create a `.kubernetes-version` file in your project directory:

```bash
# Use the current active version
kuve init

# Or specify a version
kuve init v1.28.0
```

#### Use version from file

Switch to the version specified in `.kubernetes-version` file:

```bash
kuve use
```

This command will:
1. Search for `.kubernetes-version` in the current and parent directories
2. Install the version if it's not already installed
3. Switch to that version

#### Use version from cluster

Auto-detect and use the Kubernetes version from your current cluster:

```bash
kuve use --from-cluster
# or shorthand
kuve use -c
```

This command will:
1. Connect to your current Kubernetes cluster (using current context)
2. Detect the cluster's Kubernetes version
3. Install the matching kubectl version if needed
4. Switch to that version

### Example Workflow

```bash
# Install a few versions
kuve install v1.28.0
kuve install v1.29.1

# Switch to a specific version
kuve switch v1.28.0

# Verify the current version
kuve current
kubectl version --client

# Create a version file for your project
cd my-k8s-project
kuve init v1.28.0

# Later, when you return to the project
cd my-k8s-project
kuve use  # Automatically switches to v1.28.0

# Or detect version from your cluster
kuve use --from-cluster  # Detects and uses cluster version
```

## How It Works

Kuve manages kubectl versions by:

1. Storing different kubectl versions in `~/.kuve/versions/<version>/`
2. Creating a symbolic link at `~/.kuve/bin/kubectl` that points to the active version
3. Using the symlink to switch between versions instantly

## Directory Structure

```
~/.kuve/
├── bin/
│   └── kubectl -> ../versions/v1.28.0/kubectl
└── versions/
    ├── v1.28.0/
    │   └── kubectl
    ├── v1.29.1/
    │   └── kubectl
    └── ...
```

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build -o kuve main.go
```

### Project Structure

```
kuve/
├── cmd/              # CLI commands
├── internal/         # Internal packages
│   ├── kubectl/      # kubectl installation logic
│   └── version/      # version management
├── pkg/              # Public packages
│   └── config/       # configuration
├── main.go           # Entry point
└── go.mod            # Go module file
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. See the [Contributing Guide](CONTRIBUTING.md) for details.

## Documentation

### Getting Started
- [Quick Start Guide](QUICKSTART.md) - Get started in minutes
- [Installation Guide](docs/installation.md) - Complete installation instructions
- [Usage Guide](docs/usage.md) - Comprehensive usage documentation

### Features
- [Cluster Detection](CLUSTER_DETECTION.md) - Auto-detect cluster versions
- [Shell Integration](SHELL_INTEGRATION.md) - Auto-switching setup
- [Version Files](docs/usage.md#version-files) - Project-specific versions

### Reference
- [CLI Reference](docs/cli-reference.md) - All commands and options
- [Configuration](docs/configuration.md) - Configuration and customization
- [Troubleshooting](docs/troubleshooting.md) - Common issues and solutions

### Development
- [Architecture](docs/architecture.md) - System design and architecture
- [Development Guide](DEVELOPMENT.md) - Developer documentation
- [Contributing Guide](CONTRIBUTING.md) - How to contribute

📚 **[Full Documentation Index](docs/README.md)** - Browse all documentation

## License

MIT License - see LICENSE file for details

## Author

Germain Lefebvre (@germainlefebvre4)
