# CLAUDE.md - AI Assistant Guide for cfimport

This document provides guidance for AI assistants working on the cfimport repository.

## Project Overview

**cfimport** is a native Linux CLI tool that automates AWS CloudFormation resource imports. It simplifies the process of bringing existing AWS resources under CloudFormation management.

**Language:** Go 1.21+
**License:** MIT
**Target Platform:** Linux (Debian/Ubuntu packages)
**Install:** `sudo apt install cfimport`

## What This Tool Does

cfimport automates the manual steps documented in AWS CloudFormation's resource import workflow:

1. **Scan** - Discover existing AWS resources using IaC generator
2. **Generate** - Create CloudFormation templates from scanned resources
3. **Import** - Execute the import via change sets

### AWS CloudFormation Import Process (Manual vs cfimport)

| Manual Steps | cfimport Command |
|--------------|------------------|
| Navigate to IaC generator in console | `cfimport scan` |
| Start resource scan | `cfimport scan --region us-east-1` |
| Wait for scan, select resources | `cfimport generate --scan-id X` |
| Download template | `cfimport generate -o template.yaml` |
| Create change set with IMPORT type | `cfimport import --stack X` |
| Map resource identifiers | `cfimport import --resources ids.json` |
| Execute change set | Automatic (or `--dry-run`) |

## Repository Structure

```
cf-import/
├── cmd/
│   └── cfimport/
│       └── main.go           # Application entry point
├── internal/
│   ├── cli/                  # CLI commands (cobra)
│   │   ├── root.go           # Root command, global flags (--region, --profile)
│   │   ├── scan.go           # Resource scanning via IaC generator
│   │   ├── generate.go       # Template generation from scans
│   │   ├── import.go         # Stack import operations
│   │   └── list.go           # List resources/types
│   ├── aws/                  # AWS SDK wrappers (TODO)
│   ├── config/               # Configuration handling (TODO)
│   └── template/             # Template manipulation (TODO)
├── pkg/
│   └── types/                # Public API types (TODO)
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
sudo make install     # Install to /usr/local/bin
sudo make uninstall   # Remove
```

### Building Debian Package

```bash
make deb                          # Build .deb package
sudo dpkg -i ../cfimport_*.deb    # Install
```

## CLI Commands Reference

### Global Flags

```bash
--region      AWS region (or AWS_REGION env)
--profile     AWS profile (or AWS_PROFILE env)
--config      Config file path
--verbose     Verbose output
```

### Commands

```bash
# Scan for resources
cfimport scan --region us-east-1
cfimport scan --types AWS::EC2::Instance,AWS::S3::Bucket

# Generate template from scan
cfimport generate --scan-id abc123 --output template.yaml
cfimport generate --resources i-123,bucket-name -o template.yaml

# Import into stack
cfimport import --stack mystack --template template.yaml --resources resources.json
cfimport import --stack newstack --template template.yaml --resources resources.json --create
cfimport import --stack mystack --template template.yaml --resources resources.json --dry-run

# List resources
cfimport list --types              # Show importable resource types
cfimport list --scan-id abc123     # Show resources from scan
cfimport list --stack mystack      # Show stack resources
```

## Development Workflow

### Adding a New Command

1. Create a new file in `internal/cli/` (e.g., `validate.go`)
2. Define the command using cobra:
   ```go
   var validateCmd = &cobra.Command{
       Use:   "validate",
       Short: "Validate a template for import",
       RunE:  runValidate,
   }

   func init() {
       rootCmd.AddCommand(validateCmd)
       // Add flags...
   }
   ```
3. Add tests in `internal/cli/validate_test.go`
4. Update man page in `man/cfimport.1`

### Testing

```bash
make test           # Run tests
make test-coverage  # With coverage
make lint           # Run linter
```

## Configuration

### Hierarchy (highest to lowest priority)

1. Command-line flags (`--region us-east-1`)
2. Environment variables (`AWS_REGION`)
3. Config file (`~/.cfimport.yaml`)
4. AWS config (`~/.aws/config`)

### Environment Variables

| Variable | Description |
|----------|-------------|
| `AWS_REGION` | AWS region |
| `AWS_PROFILE` | AWS profile name |
| `AWS_ACCESS_KEY_ID` | Access key |
| `AWS_SECRET_ACCESS_KEY` | Secret key |

### Config File Format

```yaml
region: us-east-1
profile: default
verbose: false
```

## Key Dependencies

- [aws/aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) - AWS SDK
- [spf13/cobra](https://github.com/spf13/cobra) - CLI framework
- [spf13/viper](https://github.com/spf13/viper) - Configuration

## AWS CloudFormation Import Requirements

When implementing features, remember these CloudFormation import rules:

1. **DeletionPolicy: Retain** - All imported resources must have this
2. **Resource identifiers** - Each resource type has specific identifier(s)
3. **Same region** - Resources must be in the same region as the stack
4. **Not managed** - Resources can't already be in another stack
5. **Supported types** - Not all resource types support import

### Resource Identifier Examples

| Resource Type | Identifier Key |
|---------------|----------------|
| AWS::EC2::Instance | InstanceId |
| AWS::S3::Bucket | BucketName |
| AWS::RDS::DBInstance | DBInstanceIdentifier |
| AWS::Lambda::Function | FunctionName |
| AWS::DynamoDB::Table | TableName |

## Common Tasks for AI Assistants

### Before Making Changes

1. Read existing code in `internal/cli/` for patterns
2. Run `make test` to ensure tests pass
3. Check AWS SDK v2 documentation for API usage

### When Implementing Features

1. Follow existing patterns in CLI commands
2. Use AWS SDK v2 (not v1)
3. Always implement `--dry-run` for destructive operations
4. Add `--wait` flag for async operations
5. Handle pagination for list operations
6. Mask sensitive data in verbose output

### AWS API Calls

```go
// Pattern for AWS SDK v2 calls
cfg, err := config.LoadDefaultConfig(ctx,
    config.WithRegion(region),
)
client := cloudformation.NewFromConfig(cfg)

// Example: Start resource scan
output, err := client.StartResourceScan(ctx, &cloudformation.StartResourceScanInput{})
```

## Commit Conventions

```
feat: add validate command for template checking
fix: handle pagination in resource listing
docs: update man page with validate command
build: upgrade AWS SDK to v1.25.0
```

## Security Considerations

- Never log AWS credentials
- Use IAM roles when possible
- Validate all user input
- Implement proper error messages (no stack traces to users)
- Support MFA and assume role workflows

## Resources

- [AWS CloudFormation Import](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import.html)
- [IaC Generator](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/generate-IaC.html)
- [Supported Resources for Import](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/resource-import-supported-resources.html)
- [AWS SDK Go v2](https://aws.github.io/aws-sdk-go-v2/docs/)
- [Cobra CLI](https://cobra.dev/)
