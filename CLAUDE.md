# CLAUDE.md - AI Assistant Guide for cfimport

This document provides guidance for AI assistants working on the cfimport repository.

## Project Overview

**cfimport** is a native Linux CLI tool for importing and exporting Cloudflare configurations and resources. It is designed to be installed via `sudo apt install cfimport` as a Debian package.

**Language:** Go 1.21+
**License:** MIT
**Target Platform:** Linux (Debian/Ubuntu packages)

## Repository Structure

```
cf-import/
├── cmd/
│   └── cfimport/
│       └── main.go           # Application entry point
├── internal/
│   ├── cli/                  # CLI commands (cobra)
│   │   ├── root.go           # Root command and global flags
│   │   ├── dns.go            # DNS import command
│   │   ├── firewall.go       # Firewall import command
│   │   └── export.go         # Export command
│   ├── cloudflare/           # Cloudflare API client (TODO)
│   ├── config/               # Configuration handling (TODO)
│   └── importer/             # Import logic (TODO)
├── pkg/
│   └── api/                  # Public API types (TODO)
├── debian/                   # Debian packaging files
│   ├── control               # Package metadata
│   ├── rules                 # Build rules
│   ├── changelog             # Version history
│   ├── copyright             # License info
│   └── source/format         # Package format
├── man/
│   └── cfimport.1            # Man page
├── Makefile                  # Build automation
├── go.mod                    # Go module definition
├── LICENSE                   # MIT License
└── CLAUDE.md                 # This file
```

## Quick Start

### Building

```bash
# Download dependencies
make deps

# Build the binary
make build

# Run the tool
./cfimport --help
```

### Installing Locally

```bash
# Install to /usr/local/bin (requires sudo)
sudo make install

# Uninstall
sudo make uninstall
```

### Building Debian Package

```bash
# Build .deb package
make deb

# Install the package
sudo dpkg -i ../cfimport_*.deb
```

## Development Workflow

### Adding a New Command

1. Create a new file in `internal/cli/` (e.g., `pages.go`)
2. Define the command using cobra:
   ```go
   var pagesCmd = &cobra.Command{
       Use:   "pages",
       Short: "Import Cloudflare Pages settings",
       RunE:  runPagesImport,
   }

   func init() {
       rootCmd.AddCommand(pagesCmd)
       // Add flags...
   }
   ```
3. Add tests in `internal/cli/pages_test.go`
4. Update man page in `man/cfimport.1`

### Code Organization

- **cmd/**: Entry points only - minimal code
- **internal/**: Private packages - core implementation
- **pkg/**: Public packages - APIs for external use

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint
```

## Commit Conventions

Follow conventional commits:
- `feat:` - New features or commands
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `refactor:` - Code refactoring
- `test:` - Test additions
- `build:` - Build system changes
- `chore:` - Maintenance tasks

Examples:
```
feat: add pages import command
fix: handle empty zone response gracefully
docs: update man page with new options
build: update Go version to 1.22
```

## Configuration

### Hierarchy (highest to lowest priority)

1. Command-line flags (`--api-token`)
2. Environment variables (`CF_API_TOKEN`)
3. Config file (`~/.cfimport.yaml` or `/etc/cfimport/config.yaml`)

### Environment Variables

| Variable | Description |
|----------|-------------|
| `CF_API_TOKEN` | Cloudflare API token (preferred) |
| `CF_API_KEY` | Cloudflare API key (legacy) |
| `CF_API_EMAIL` | Account email for API key auth |

### Config File Format

```yaml
api_token: "your-api-token"
default_zone: "example.com"
verbose: false
```

**Never commit credentials to the repository.**

## Key Dependencies

- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [spf13/viper](https://github.com/spf13/viper) - Configuration management

## Common Tasks for AI Assistants

### Before Making Changes

1. Read existing code to understand patterns
2. Run `make test` to ensure tests pass
3. Check `go.mod` for dependency versions

### When Implementing Features

1. Follow existing code patterns in `internal/cli/`
2. Add appropriate flags and help text
3. Implement dry-run mode where applicable
4. Add error handling with clear messages
5. Update man page documentation
6. Write tests for new functionality

### Version Updates

1. Update `debian/changelog` with new entry
2. Tag release with `git tag v0.x.x`
3. Build packages with `make deb`

## Cloudflare API Guidelines

- Always use API tokens over API keys when possible
- Implement rate limiting handling
- Support pagination for list operations
- Provide dry-run mode for destructive operations
- Log API calls in verbose mode (without secrets)

## Security Considerations

- Never log or display API tokens/keys
- Validate all user input before API calls
- Use HTTPS for all API communication
- Sanitize file paths to prevent directory traversal
- Don't execute user-provided data as commands

## Debian Package Notes

- Binary installs to `/usr/bin/cfimport`
- Man page installs to `/usr/share/man/man1/cfimport.1`
- System config goes to `/etc/cfimport/`
- User config goes to `~/.cfimport.yaml`

## Resources

- Cloudflare API: https://developers.cloudflare.com/api/
- Cobra CLI: https://cobra.dev/
- Debian Packaging: https://www.debian.org/doc/manuals/maint-guide/
