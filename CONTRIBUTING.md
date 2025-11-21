# Contributing to go-arch

Thank you for your interest in contributing to **go-arch**! 🚀

This document explains how you can contribute, the development workflow, coding standards, and how to submit issues or pull requests.

---

## 🧱 Table of Contents

* [Getting Started](#getting-started)
* [How to Contribute](#how-to-contribute)

  * [Reporting Bugs](#reporting-bugs)
  * [Suggesting Features](#suggesting-features)
  * [Submitting Pull Requests](#submitting-pull-requests)
* [Development Setup](#development-setup)
* [Coding Guidelines](#coding-guidelines)
* [Commit Messages](#commit-messages)
* [Branching Model](#branching-model)
* [Code Review Process](#code-review-process)
* [License](#license)

---

## Getting Started

Before contributing, please:

* Review the project's goals and structure.
* Make sure an issue does not already exist.
* Follow the guidelines in this file.

If you're unsure about anything, feel free to open an issue and ask.

---

## How to Contribute

### Reporting Bugs

If you find a bug, create a GitHub issue with the following:

* A clear description of the problem
* Steps to reproduce
* Expected vs actual behavior
* Go version and OS information

Use the label **bug**.

### Suggesting Features

Feature requests are welcome! Please include:

* The problem the feature solves
* A high-level description
* Optional: a small example

Use the label **enhancement**.

### Submitting Pull Requests

1. Fork the repository
2. Create a new branch: `git checkout -b feature/my-feature`
3. Make changes and commit them following the commit message rules
4. Push the branch and open a Pull Request
5. Ensure CI passes and reviewers approve

---

## Development Setup

### Requirements

* Go 1.22+
* Make (optional but recommended)

### Running the Project

```bash
go run ./...
```

### Running Tests

```bash
go test ./... -cover
```

### Running Linters (if configured)

```bash
make lint
```

---

## Coding Guidelines

* Follow standard Go formatting using `gofmt` or `go fmt`.
* Keep functions small and focused.
* Avoid unnecessary abstractions.
* Add comments where needed, especially for exported functions.
* Write unit tests for new code.

---

## Commit Messages

Follow conventional commit guidelines:

```
feat: add new interface for service
fix: correct nil pointer dereference
refactor: simplify config loading
test: add missing unit tests
docs: update README and examples
```

---

## Branching Model

* `main` → always stable
* Feature branches → `feature/...`
* Bug fixes → `fix/...`
* Documentation → `docs/...`

---

## Code Review Process

When you open a PR:

* Automated tests must pass
* Code should be clean and idiomatic
* Reviewers may request changes
* After approval, the PR will be merged

---

## License

By contributing, you agree that your contributions will be licensed under the repository's license.

---

Thanks again for contributing! 🙌
