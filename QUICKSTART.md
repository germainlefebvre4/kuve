# Quick Start Guide

Get started with Kuve in just a few minutes!

## 1. Build and Install

```bash
# Clone the repository (if not already done)
git clone https://github.com/germainlefebvre4/kuve.git
cd kuve

# Build and install
make install
```

Or build without installing:
```bash
make build
./kuve --help
```

## 2. Set Up Your Shell

Add Kuve's bin directory to your PATH:

```bash
# For Bash
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# For Zsh
echo 'export PATH="$HOME/.kuve/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

## 3. Install Your First kubectl Version

```bash
# Install the latest stable version
kuve install v1.29.1

# Or install a specific version
kuve install v1.28.0
```

## 4. Switch to a Version

```bash
# Switch to the installed version
kuve switch v1.29.1

# Verify it's active
kuve current
kubectl version --client
```

## 5. Try Version Files

```bash
# Go to your project directory
cd ~/my-k8s-project

# Create a version file
kuve init v1.28.0

# The file is created
cat .kubernetes-version
# Output: v1.28.0

# Use the version from the file
kuve use
```

## 6. Try Cluster Detection

```bash
# Make sure you have a Kubernetes cluster configured
kubectl cluster-info

# Auto-detect and use the cluster version
kuve use --from-cluster

# Verify it worked
kuve current
```

## Common Commands

```bash
# List installed versions
kuve list installed

# List remote versions (shows latest stable)
kuve list remote

# Show current version
kuve current

# Uninstall a version
kuve uninstall v1.27.0

# Get help
kuve --help
kuve <command> --help
```

## What's Next?

- Check out the [README.md](README.md) for detailed documentation
- See [SHELL_INTEGRATION.md](SHELL_INTEGRATION.md) for automatic version switching
- Read [CONTRIBUTING.md](CONTRIBUTING.md) if you want to contribute

## Troubleshooting

### kubectl command not found

Make sure `~/.kuve/bin` is in your PATH:
```bash
echo $PATH | grep .kuve
```

If not, add it to your shell configuration and reload.

### Version not switching

1. Check that the version is installed:
   ```bash
   kuve list installed
   ```

2. Verify the symlink exists:
   ```bash
   ls -la ~/.kuve/bin/kubectl
   ```

3. Make sure you're using kuve's kubectl:
   ```bash
   which kubectl
   # Should output: /home/yourusername/.kuve/bin/kubectl
   ```

### Need Help?

- Run `kuve --help` for command help
- Check the full [README.md](README.md) for more details
- Open an issue on GitHub if you find a bug

Happy Kubernetes managing! 🎉
