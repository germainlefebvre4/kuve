# Shell Integration Examples

## Bash

Add to your `~/.bashrc`:

```bash
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"

# Optional: Auto-switch kubectl version when entering a directory
kuve_auto_switch() {
    if [ -f .kubernetes-version ]; then
        local version=$(cat .kubernetes-version | tr -d '[:space:]')
        if [ -n "$version" ]; then
            kuve use 2>/dev/null || echo "Failed to switch to kubectl $version"
        fi
    fi
}

# Run on directory change
cd() {
    builtin cd "$@" && kuve_auto_switch
}

# Run on shell startup (for current directory)
kuve_auto_switch
```

## Zsh

Add to your `~/.zshrc`:

```zsh
# Kuve - Kubernetes Client Switcher
export PATH="$HOME/.kuve/bin:$PATH"

# Optional: Auto-switch kubectl version when entering a directory
kuve_auto_switch() {
    if [ -f .kubernetes-version ]; then
        local version=$(cat .kubernetes-version | tr -d '[:space:]')
        if [ -n "$version" ]; then
            kuve use 2>/dev/null || echo "Failed to switch to kubectl $version"
        fi
    fi
}

# Use chpwd hook for automatic switching
autoload -U add-zsh-hook
add-zsh-hook chpwd kuve_auto_switch

# Run on shell startup (for current directory)
kuve_auto_switch
```

## Fish

Add to your `~/.config/fish/config.fish`:

```fish
# Kuve - Kubernetes Client Switcher
set -gx PATH "$HOME/.kuve/bin" $PATH

# Optional: Auto-switch kubectl version when entering a directory
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

## Manual Setup (Without Auto-Switch)

If you prefer to manually switch versions, you only need to add the bin directory to your PATH:

### Bash/Zsh
```bash
export PATH="$HOME/.kuve/bin:$PATH"
```

### Fish
```fish
set -gx PATH "$HOME/.kuve/bin" $PATH
```

## Verification

After setting up your shell integration, restart your shell or source the configuration:

```bash
# Bash
source ~/.bashrc

# Zsh
source ~/.zshrc

# Fish
source ~/.config/fish/config.fish
```

Then verify:

```bash
which kubectl
# Should output: /home/yourusername/.kuve/bin/kubectl

kubectl version --client
# Should show the active version managed by kuve
```

## Tips

1. **Auto-switch feature**: The auto-switch function automatically runs `kuve use` when you enter a directory with a `.kubernetes-version` file.

2. **Performance**: The auto-switch function only runs when necessary (when a `.kubernetes-version` file exists), so it has minimal impact on shell performance.

3. **Error handling**: The auto-switch function silently fails if there's an error, so it won't interrupt your workflow.

4. **Disable auto-switch**: If you prefer manual control, simply remove the auto-switch functions from your shell configuration and keep only the PATH export.
