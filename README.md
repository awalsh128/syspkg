# SysPkg

[![Test and Coverage](https://github.com/bluet/syspkg/actions/workflows/test-and-coverage.yml/badge.svg)](https://github.com/bluet/syspkg/actions/workflows/test-and-coverage.yml)
[![Build](https://github.com/bluet/syspkg/actions/workflows/build.yml/badge.svg)](https://github.com/bluet/syspkg/actions/workflows/build.yml)
[![Lint and Format](https://github.com/bluet/syspkg/actions/workflows/lint-and-format.yml/badge.svg)](https://github.com/bluet/syspkg/actions/workflows/lint-and-format.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bluet/syspkg)](https://goreportcard.com/report/github.com/bluet/syspkg)
[![Go Reference](https://pkg.go.dev/badge/github.com/bluet/syspkg.svg)](https://pkg.go.dev/github.com/bluet/syspkg)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](https://github.com/bluet/syspkg/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bluet/syspkg)](https://github.com/bluet/syspkg)
[![GitHub release](https://img.shields.io/github/v/release/bluet/syspkg)](https://github.com/bluet/syspkg/releases)

SysPkg is a unified CLI tool and Golang library for managing system packages across different package managers. Currently, it supports APT, YUM, Snap, and Flatpak, with plans for more. It simplifies package management by providing a consistent interface and API through an abstraction layer that focuses on package manager tools rather than specific operating systems.

## Key Features

- **Cross-Package Manager Compatibility**: Normalized status reporting (e.g., APT's config-files state maps to available)
- **Consistent API**: Same interface across all supported package managers
- **Tool-Focused**: Works wherever package manager tools work (containers, cross-platform, etc.)
- **Production Ready**: Comprehensive testing across multiple OS distributions
- **Performance Optimized**: Efficient parsing with compiled regexes and robust error handling
- **Cross-Platform**: Handles different line endings (CRLF/LF) and whitespace variations

## Features

- A unified package management interface for various package managers
- Supports popular package managers such as APT, YUM, Snap, Flatpak, and more
- Easy-to-use API for package installation, removal, search, listing, and system upgrades
- Expandable architecture to support more package managers in the future

## API Documentation

See the [Go Reference](https://pkg.go.dev/github.com/bluet/syspkg) for the full API documentation.

## Getting Started

### Prerequisites

- Go 1.23 or later

### Installation (as CLI tool)

Install the CLI tool using the `go install` command:

```bash
go install github.com/bluet/syspkg/cmd/syspkg@latest
```

### Installation (as Go library)

Install the library using the `go get` command:

```bash
go get github.com/bluet/syspkg
```

## Usage

### CLI Tool

SysPkg provides a unified CLI tool for managing system packages across different package managers. It simplifies the process of working with various package managers by providing a consistent interface through an abstraction layer.

Here's an example demonstrating how to use SysPkg as a CLI tool:

```bash
# Install a package using APT
syspkg --apt install vim

# Remove a package using APT
syspkg --apt remove vim

# Search for a package using Snap
syspkg --snap search vim

# Show all upgradable packages using Flatpak
syspkg --flatpak show upgradable

# Install a package using YUM (on RHEL/CentOS/Rocky/AlmaLinux)
syspkg --yum install vim

# Show package information
syspkg --apt show package vim

# List installed packages
syspkg --snap show installed

# List upgradable packages
syspkg --flatpak show upgradable
```

Or, you can do operations without knowing the package manager:

```bash
# Install a package using all available package managers
syspkg install vim

# Remove a package using all available package manager
syspkg remove vim

# Search for a package using all available package manager
syspkg search vim

# Upgrade all packages using all available package manager
syspkg upgrade

# Refresh package lists
syspkg refresh

# Show package information
syspkg show package vim

# List installed packages
syspkg show installed

# List upgradable packages
syspkg show upgradable
```

For more examples and real use cases, see the [cmd/syspkg/](cmd/syspkg/) directory.

### Go Library

Here's an example demonstrating how to use SysPkg as a Go library:

```go
package main

import (
 "fmt"
 "github.com/bluet/syspkg"
)

func main() {
 // Initialize SysPkg with all available package managers on current system
 includeOptions := syspkg.IncludeOptions{
  AllAvailable: true,
 }
 syspkgManager, err := syspkg.New(includeOptions)
 if err != nil {
  fmt.Printf("Error initializing SysPkg: %v\n", err)
  return
 }

 // Get APT package manager (if available)
 aptManager, err := syspkgManager.GetPackageManager("apt")
 if err != nil {
  fmt.Printf("APT package manager not available: %v\n", err)
  return
 }

 // List installed packages using APT
 installedPackages, err := aptManager.ListInstalled(nil)
 if err != nil {
  fmt.Printf("Error listing installed packages: %v\n", err)
  return
 }

 fmt.Println("Installed packages:")
 for _, pkg := range installedPackages {
  fmt.Printf("- %s (%s)\n", pkg.Name, pkg.Version)
 }
}
```

For more examples and real use cases, see the [cmd/syspkg/](cmd/syspkg/) directory.

## Supported Package Managers

| Package Manager | Install | Remove | Search | Upgrade | List Installed | List Upgradable | Get Package Info | AutoRemove | Clean | Refresh |
| --------------- | ------- | ------ | ------ | ------- | -------------- | --------------- | ---------------- | ---------- | ----- | ------- |
| APT             | ✅      | ✅    | ✅     | ✅     | ✅             | ✅             | ✅               | ✅         | ✅    | ✅      |
| YUM             | ✅      | ✅    | ✅     | ✅     | ✅             | ✅             | ✅               | ✅         | ✅    | ✅      |
| SNAP            | ✅      | ✅    | ✅     | ✅     | ✅             | ✅             | ✅               | ✅         | ✅    | ✅      |
| Flatpak         | ✅      | ✅    | ✅     | ✅     | ✅             | ✅             | ✅               | ✅         | ✅    | ✅      |
| DNF             | 🚧      | 🚧    | 🚧     | 🚧     | 🚧             | 🚧             | 🚧               | 🚧         | 🚧    | 🚧      |
| APK (Alpine)    | 🚧      | 🚧    | 🚧     | 🚧     | 🚧             | 🚧             | 🚧               | 🚧         | 🚧    | 🚧      |
| Zypper (openSUSE) | 🚧   | 🚧    | 🚧     | 🚧     | 🚧             | 🚧             | 🚧               | 🚧         | 🚧    | 🚧      |

**Legend:** ✅ Implemented, 🚧 Planned, ❌ Not supported

**Philosophy:** SysPkg focuses on supporting package manager tools wherever they work, regardless of the underlying operating system. If you have apt+dpkg working on macOS via Homebrew, or in a container, SysPkg will support it.

Please open an issue (or PR ❤️) if you'd like to see support for any unlisted specific package manager.

## Development

### Documentation

#### For Users
- **[README.md](README.md)** - Getting started, usage examples, and features
- **[Go Reference](https://pkg.go.dev/github.com/bluet/syspkg)** - Complete API documentation

#### For Contributors
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Development workflow and comprehensive testing guide
- **[docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)** - Technical architecture, design patterns, and interfaces
- **[docs/EXIT_CODES.md](docs/EXIT_CODES.md)** - Package manager exit code behavior (critical for implementation)
- **[manager/{pm}/EXIT_CODES.md](manager/)** - Package manager specific exit code documentation

#### For AI Assistants
- **[CLAUDE.md](CLAUDE.md)** - Development guidelines and AI-specific instructions

#### Infrastructure
- **[testing/](testing/)** - Test fixtures and Docker testing infrastructure
- **[.pre-commit-config.yaml](.pre-commit-config.yaml)** - Secure pre-commit hooks
- **[.github/workflows/](.github/workflows/)** - CI/CD pipelines

### Recent Fixes
- ✅ **APT Exit Code Bug**: Fixed in commit 3751f45 - now properly propagates errors (Issue #21)
- ✅ **Snap Exit Code Bug**: Fixed in commit 3751f45 - now properly handles usage errors (Issue #22)
- ✅ **Flatpak Exit Code Bug**: Fixed in commit 3751f45 - now properly handles general errors (Issue #24)

### Active Development
- **CommandBuilder Architecture**: Ongoing refactoring for Issue #20 to improve testability
- **Security Enhancements**: Input validation for package names (Issue #23)

### CI/CD Status

| Workflow | Status | Description |
| -------- | ------ | ----------- |
| **Test and Coverage** | ✅ | Go 1.23/1.24 testing with coverage reporting |
| **Lint and Format** | ✅ | golangci-lint, gofmt, go vet quality checks |
| **Build** | ✅ | Multi-version build verification |
| **Multi-OS Tests** | ✅ | Docker-based testing across Ubuntu, Rocky Linux, Alpine |
| **Release Binaries** | ✅ | Cross-platform binary releases |

- ✅ **Pre-commit hooks**: Automated code quality and security checks
- ✅ **Go mod verification**: Dependency integrity validation
- ✅ **Multi-OS compatibility**: Docker testing with Go 1.23.4 across distributions
- ✅ **Fixture-based testing**: Real package manager output validation

## Contributing

We welcome contributions to SysPkg!

### For Users
- **Bug reports**: Open an issue with details about the problem
- **Feature requests**: Please let us know what package managers or features you'd like to see

### For Developers
- **Quick start**: See [CONTRIBUTING.md](CONTRIBUTING.md) for a comprehensive development guide
- **Architecture**: See [CLAUDE.md](CLAUDE.md) for detailed technical documentation
- **Testing strategy**: Multi-OS Docker testing with environment-aware test execution

**Development workflow:**
```bash
git clone https://github.com/bluet/syspkg.git
cd syspkg
make test          # ✅ Smart testing - works on any OS (30s)
make check         # ✅ Code quality checks (15s)

# Working on package managers? See CONTRIBUTING.md for:
# make test-docker-rocky   # 🐳 Test YUM on Rocky Linux (5min)
# make test-docker-fedora  # 🐳 Test DNF on Fedora (5min)
# make test-docker-all     # 🐳 Test all OS (15min)
```

**🎯 Quick decision:** Always start with `make test` - it automatically detects your OS and tests what's available!

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
