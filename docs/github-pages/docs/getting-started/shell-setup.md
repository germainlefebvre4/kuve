---
sidebar_position: 3
---

# Shell Setup

Configure your shell for the best Kuve experience.

## Basic Setup

At minimum, you need to add Kuve's bin directory to your PATH. This is covered in the [Installation](./installation) guide, but here's a quick reference:

### Bash

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

### Zsh

```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

### Fish

```fish
set -gx PATH "$HOME/.kuve/bin" $PATH
```

## Shell Completion

Enable tab completion for Kuve commands:

### Bash Completion

```bash
# System-wide
sudo kuve completion bash > /etc/bash_completion.d/kuve

# Current user
echo 'source <(kuve completion bash)' >> ~/.bashrc
source ~/.bashrc
```

### Zsh Completion

```zsh
# Generate completion
kuve completion zsh > "${fpath[1]}/_kuve"

# Reload completions
autoload -U compinit && compinit
```

### Fish Completion

```fish
# Generate completion
kuve completion fish > ~/.config/fish/completions/kuve.fish
```

## Verification

Test your shell setup:

### Test PATH

```bash
# Verify kuve is in PATH
which kuve
# Expected: /home/username/.kuve/bin/kuve

# Verify kubectl symlink works
which kubectl
# Expected: /home/username/.kuve/bin/kubectl
```

### Test Completion

```bash
# Type and press TAB
kuve <TAB>
# Should show: install, uninstall, switch, list, use, init, etc.
```

## Troubleshooting

### Completion Not Working

1. **Verify completion is installed**:
   ```bash
   # Bash
   ls /etc/bash_completion.d/kuve

   # Zsh
   ls ${fpath[1]}/_kuve

   # Fish
   ls ~/.config/fish/completions/kuve.fish
   ```

2. **Reload shell**:
   ```bash
   exec $SHELL
   ```

## Next Steps

- [**Basic Usage**](../user-guide/basic-usage) - Learn essential Kuve commands
- [**Version Files**](../user-guide/version-files) - Master project-specific versions
- [**Workflows**](../user-guide/workflows) - Common usage patterns
