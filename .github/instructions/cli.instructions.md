# CLI instructions

## Technologies

- Golang v1.25
- Cobra CLI
- GVM for Go version management

## Uses cases

- Handle different kubernetes distributions:
  - k3s:
    - v1.2.3+k3s1 > v1.2.3
  - microk8s:
    - v1.2.3+microk8s1 > v1.2.3
  - k0s:
  - aws eks:
    - v1.2.3-eks-4567abc > v1.2.3
  - gcp gke:
    - v1.2.3-gke.1 > v1.2.3
    - v1.33.5-gke.1308000 > v1.33.5
  - azure aks:
    - v1.2.3-aks-4567abc > v1.2.3

## Coding Standards

- Follow Go best practices and idioms
- Write clean, readable, and maintainable code
- Use meaningful variable and function names
- Keep functions small and focused on a single task
- Handle errors gracefully and provide informative error messages
- Write unit tests for critical components
- Use Go modules for dependency management
- Document code with comments where necessary
