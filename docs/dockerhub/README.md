# Quick reference

* **Maintained by**:<br>
  [Germain LEFEBVRE](https://github.com/germainlefebvre4)

* **Where to get help**:<br>
  [Github Discussions](https://github.com/germainlefebvre4/kuve/discussions)

# Supported tags and respective Dockerfile links

* [`latest`, `v0`, `v0.2`, `v0.2.1`](https://github.com/germainlefebvre4/kuve/blob/v0.2.1/Dockerfile)

# Quick reference (cont.)

* **Where to file issues**:<br>
  https://github.com/germainlefebvre4/kuve/issues⁠

* Supported architectures: ([more info⁠]())<br>
  amd64, arm64

* **Source of this description**:<br>
  [kuve repo's `docs/dockerhub/` directory](https://github.com/germainlefebvre4/kuve/tree/main/docs/dockerhub/) ([history](https://github.com/docker-library/docs/commits/master/nginx))

# What is Kuve?

Kuve is a CLI tool to easily switch versions of kubectl.

## Features

* **Easily switch kubectl versions:** Quickly change between different versions of kubectl to match your cluster requirements.
* **Automatic version detection:** Kuve can detect the appropriate kubectl version based on your Kubernetes cluster.
* **Version management:** Install, list, and uninstall different kubectl versions with simple commands.
* **Configuration file support:** Use a `.kubernetes-version` file to specify the desired kubectl version for your projects.
* **Lightweight and portable:** Kuve is a small CLI tool that can be easily integrated into your development workflow.

# How to use this image

Kuve helper.

```raw
Kuve is a CLI tool to easily switch versions of kubectl.

It helps you manage multiple kubectl versions on your system,
allowing you to install, switch, and use different versions
based on your needs or project requirements.

Usage:
  kuve [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  current     Show the current kubectl version
  help        Help about any command
  init        Create a .kubernetes-version file
  install     Install a specific kubectl version
  list        List kubectl versions
  switch      Switch to a specific kubectl version
  uninstall   Uninstall a specific kubectl version
  use         Use kubectl version from .kubernetes-version file or cluster

Flags:
  -h, --help      help for kuve
  -v, --verbose   verbose output
      --version   version for kuve

Use "kuve [command] --help" for more information about a command.
```

# License

View [license information⁠](https://github.com/germainlefebvre4/kuve/blob/main/LICENSE) for the software contained in this image.

As with all Docker images, these likely also contain other software which may be under other licenses (such as Bash, etc from the base distribution, along with any direct or indirect dependencies of the primary software being contained).

As for any pre-built image usage, it is the image user's responsibility to ensure that any use of this image complies with any relevant licenses for all software contained within.
