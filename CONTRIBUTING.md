# Contributing to Kuve

Thank you for considering contributing to Kuve! This document provides guidelines for contributing to the project.

## Development Setup

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/kuve.git
   cd kuve
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Build the project:
   ```bash
   make build
   ```

## Development Workflow

1. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes following the coding standards

3. Run tests:
   ```bash
   make test
   ```

4. Format your code:
   ```bash
   make fmt
   ```

5. Commit your changes with a descriptive message:
   ```bash
   git commit -m "Add feature: description"
   ```

6. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

7. Create a Pull Request

## Coding Standards

- Follow Go best practices and idioms
- Write clean, readable, and maintainable code
- Use meaningful variable and function names
- Keep functions small and focused on a single task
- Handle errors gracefully with informative messages
- Add unit tests for new functionality
- Document exported functions and types
- Use `go fmt` to format code

## Testing

- Write unit tests for critical components
- Aim for good test coverage
- Run tests before submitting a PR:
  ```bash
  make test
  ```

## Pull Request Guidelines

- Provide a clear description of the changes
- Reference any related issues
- Ensure all tests pass
- Update documentation if needed
- Keep PRs focused on a single feature or fix

## Reporting Issues

When reporting issues, please include:

- A clear description of the problem
- Steps to reproduce
- Expected vs actual behavior
- Your environment (OS, Go version, etc.)
- Any relevant logs or error messages

## Code of Conduct

- Be respectful and constructive
- Welcome newcomers and help them learn
- Focus on what is best for the community
- Show empathy towards other community members

Thank you for contributing to Kuve!
