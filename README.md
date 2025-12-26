# Kuve - Kubernetes Client Switcher

Kuve is a lightweight CLI tool for managing and switching between multiple kubectl versions on your system.

**[📚 Documentation](https://kuve.germainlefebvre.fr)** | **[🐳 Docker Hub](https://hub.docker.com/repository/docker/germainlefebvre4/kuve/general)**

## Getting Started

```bash
# Download the latest release
curl -LO https://github.com/germainlefebvre4/kuve/releases/latest/download/kuve-linux-amd64
chmod +x kuve-linux-amd64
sudo mv kuve-linux-amd64 /usr/local/bin/kuve

# Add to PATH
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# Install and switch kubectl versions
kuve install v1.29.1
kuve switch v1.29.1
```

For other platforms and installation methods, see the [installation guide](https://kuve.germainlefebvre.fr/docs/getting-started/installation).

## Features

- **Version Management** - Install, list, and uninstall specific kubectl versions
- **Instant Switching** - Switch between installed versions with zero overhead
- **Version Files** - Use `.kubernetes-version` files for project-specific versions
- **Cluster Detection** - Auto-detect and match your cluster's Kubernetes version
- **Automatic Installation** - Install missing versions on-the-fly when switching
- **Shell Completion** - Tab completion for Bash, Zsh, Fish, and PowerShell

## Usage Examples

### Basic Version Management

```bash
# Install a specific version
kuve install v1.28.0

# List installed versions
kuve list installed

# Switch to a version
kuve switch v1.28.0

# Show current version
kuve current
```

### Project-Specific Versions

```bash
# Create version file for your project
cd my-k8s-project
kuve init v1.28.0

# Automatically switch when entering the directory
kuve use
```

### Cluster Version Matching

```bash
# Auto-detect and use cluster version
kuve use --from-cluster
```

## How It Works

Kuve stores kubectl binaries in `~/.kuve/versions/` and manages a symbolic link at `~/.kuve/bin/kubectl` that points to the active version. This approach provides instant version switching without performance overhead.

```
~/.kuve/
├── bin/
│   └── kubectl -> ../versions/v1.28.0/kubectl
└── versions/
    ├── v1.28.0/kubectl
    ├── v1.29.1/kubectl
    └── ...
```

## Documentation

Complete documentation including installation guides, usage examples, CLI reference, and troubleshooting is available at [kuve.germainlefebvre.fr](https://kuve.germainlefebvre.fr).

## Docker Hub

Kuve is also available as a Docker image on [Docker Hub](https://hub.docker.com/repository/docker/germainlefebvre4/kuve/general). Use it in containerized environments or CI/CD pipelines.

## Contributing

Contributions are welcome! Please see the [Contributing Guide](https://kuve.germainlefebvre.fr/docs/developers/contributing) for details

## License

MIT License - see LICENSE file for details

## Author

Germain Lefebvre (@germainlefebvre4)
