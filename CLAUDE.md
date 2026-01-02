# CLAUDE.md - AI Assistant Guide for cf-import

This document provides guidance for AI assistants working on the cf-import repository.

## Project Overview

**cf-import** is a project for importing Cloudflare configurations, resources, or data. This repository is in early development and follows MIT licensing.

## Repository Structure

```
cf-import/
├── LICENSE          # MIT License
├── CLAUDE.md        # This file - AI assistant guidance
└── .git/            # Git version control
```

As the project evolves, this structure section should be updated to reflect new directories and files.

## Development Workflow

### Branch Naming

- Feature branches: `claude/<feature-name>-<session-id>` or `feature/<description>`
- Bug fixes: `fix/<description>`
- Documentation: `docs/<description>`

### Commit Conventions

Follow conventional commit format:
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `refactor:` - Code refactoring
- `test:` - Test additions or modifications
- `chore:` - Maintenance tasks

Example: `feat: add cloudflare zone import functionality`

### Git Operations

1. Always work on feature branches, not main
2. Use `git push -u origin <branch-name>` for initial push
3. Keep commits atomic and focused on single changes
4. Write descriptive commit messages explaining "why" not just "what"

## Code Conventions

### General Principles

1. **Keep it simple** - Avoid over-engineering; implement only what's needed
2. **Security first** - Never expose API keys, tokens, or sensitive data
3. **Error handling** - Provide clear error messages for debugging
4. **Documentation** - Document public APIs and complex logic

### File Organization

- Keep related functionality grouped together
- Use clear, descriptive file and function names
- Prefer flat structures over deep nesting when possible

## Configuration

### Environment Variables

When Cloudflare API integration is added, expect these environment variables:
- `CF_API_TOKEN` - Cloudflare API token (preferred)
- `CF_API_KEY` - Cloudflare API key (legacy)
- `CF_API_EMAIL` - Associated email for API key auth

**Never commit credentials to the repository.**

## Testing

When tests are added to the project:
- Run tests before committing changes
- Ensure all tests pass before pushing
- Add tests for new functionality

## Common Tasks for AI Assistants

### Before Making Changes

1. Read existing code to understand patterns
2. Check for existing similar functionality
3. Understand the project's coding style

### When Implementing Features

1. Start with the minimal viable implementation
2. Add error handling for expected failure cases
3. Consider edge cases but don't over-engineer
4. Test changes locally when possible

### Security Considerations

- Validate all external input
- Use secure API practices (HTTPS, token auth)
- Never log sensitive data
- Follow Cloudflare API best practices

## Getting Help

- Cloudflare API Documentation: https://developers.cloudflare.com/api/
- Repository Issues: Check GitHub issues for context on known problems

## Notes for Future Development

This section should be updated as the project matures:
- [ ] Define primary programming language
- [ ] Add dependency management approach
- [ ] Document build and deployment processes
- [ ] Add CI/CD pipeline details
