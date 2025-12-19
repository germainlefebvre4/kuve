# Documentation Index

Welcome to Kuve documentation! This guide will help you find the information you need.

## Getting Started

New to Kuve? Start here:

1. **[Installation Guide](./installation.md)** - Install Kuve on your system
2. **[Quick Start](../QUICKSTART.md)** - Get up and running in minutes
3. **[Usage Guide](./usage.md)** - Learn how to use Kuve effectively

## Core Documentation

### User Guides

| Document | Description | When to Read |
|----------|-------------|--------------|
| [Installation](./installation.md) | Complete installation instructions | First time setup |
| [Usage](./usage.md) | Comprehensive usage guide | Learn all features |
| [CLI Reference](./cli-reference.md) | Complete command reference | Look up specific commands |
| [Configuration](./configuration.md) | Configuration and customization | Customize your setup |
| [Troubleshooting](./troubleshooting.md) | Common issues and solutions | When you encounter problems |

### Technical Documentation

| Document | Description | Audience |
|----------|-------------|----------|
| [Architecture](./architecture.md) | System design and architecture | Developers, Contributors |
| [Development](../DEVELOPMENT.md) | Development guide | Contributors |
| [Contributing](../CONTRIBUTING.md) | How to contribute | Contributors |

### Feature-Specific Guides

| Document | Description |
|----------|-------------|
| [Cluster Detection](../CLUSTER_DETECTION.md) | Auto-detect cluster versions |
| [Shell Integration](../SHELL_INTEGRATION.md) | Auto-switching setup |

## Quick Links

### Common Tasks

- [Install kubectl version](./usage.md#installing-versions)
- [Switch versions](./usage.md#switching-versions)
- [Use version files](./usage.md#version-files)
- [Detect cluster version](./usage.md#cluster-detection)
- [Set up auto-switching](../SHELL_INTEGRATION.md)

### Reference

- [All commands](./cli-reference.md#commands)
- [Version formats](./cli-reference.md#version-formats)
- [Exit codes](./cli-reference.md#exit-codes)
- [Directory structure](./configuration.md#directory-structure)

### Troubleshooting

- [Installation issues](./troubleshooting.md#installation-issues)
- [PATH issues](./troubleshooting.md#path-and-command-issues)
- [Version management issues](./troubleshooting.md#version-management-issues)
- [Network issues](./troubleshooting.md#network-and-download-issues)

## Documentation by Role

### End Users

Recommended reading order:

1. [Installation Guide](./installation.md)
2. [Quick Start](../QUICKSTART.md)
3. [Usage Guide](./usage.md)
4. [CLI Reference](./cli-reference.md) (as needed)
5. [Troubleshooting](./troubleshooting.md) (when needed)

### System Administrators

Focus on these documents:

1. [Installation Guide](./installation.md)
2. [Configuration](./configuration.md)
3. [Shell Integration](../SHELL_INTEGRATION.md)
4. [Troubleshooting](./troubleshooting.md)

### Developers/Contributors

Essential reading:

1. [Architecture](./architecture.md)
2. [Development Guide](../DEVELOPMENT.md)
3. [Contributing Guide](../CONTRIBUTING.md)
4. [Usage Guide](./usage.md) (for testing)

## Documentation Structure

```
docs/
├── README.md                  # This file (documentation index)
├── installation.md            # Installation instructions
├── usage.md                   # Comprehensive usage guide
├── cli-reference.md           # Complete CLI reference
├── configuration.md           # Configuration options
├── troubleshooting.md         # Common issues and solutions
└── architecture.md            # Technical architecture

Root level documentation:
├── README.md                  # Project overview
├── QUICKSTART.md              # Quick start guide
├── CLUSTER_DETECTION.md       # Cluster detection feature
├── SHELL_INTEGRATION.md       # Shell integration examples
├── DEVELOPMENT.md             # Development guide
└── CONTRIBUTING.md            # Contributing guidelines
```

## Getting Help

### In Documentation

1. Use the search function in GitHub
2. Check the [Troubleshooting Guide](./troubleshooting.md)
3. Review [CLI Reference](./cli-reference.md) for command details

### Command Line Help

```bash
# General help
kuve --help

# Command-specific help
kuve install --help
kuve use --help

# Enable verbose mode for debugging
kuve --verbose install v1.28.0
```

### Community Support

- **GitHub Issues**: https://github.com/germainlefebvre4/kuve/issues
- **Discussions**: https://github.com/germainlefebvre4/kuve/discussions

### Reporting Documentation Issues

Found a problem in the documentation?

1. Check if it's already reported
2. Create an issue with:
   - Document name
   - Section/paragraph
   - What's wrong or unclear
   - Suggested improvement

## Contributing to Documentation

Documentation improvements are welcome! See [Contributing Guide](../CONTRIBUTING.md).

### Documentation Guidelines

- Clear and concise language
- Include examples
- Use consistent formatting
- Test all commands before documenting
- Keep up-to-date with code changes

## Version Information

This documentation is for:
- **Kuve**: Latest version
- **kubectl**: v1.26.0 and later

Last updated: December 2025

## Related Resources

### Official Kubernetes Documentation

- [kubectl Reference](https://kubernetes.io/docs/reference/kubectl/)
- [kubectl Installation](https://kubernetes.io/docs/tasks/tools/)
- [kubectl Releases](https://github.com/kubernetes/kubernetes/releases)

### Version Management

- [Semantic Versioning](https://semver.org/)
- [kubectl Version Skew Policy](https://kubernetes.io/docs/setup/release/version-skew-policy/)

## Feedback

We value your feedback! If you have suggestions for improving this documentation:

1. Open an issue on GitHub
2. Submit a pull request with improvements
3. Share your experience in discussions

Thank you for using Kuve!
