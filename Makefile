# cfimport Makefile
# Build and package the Cloudflare Import Tool

BINARY := cfimport
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0-dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

GOFLAGS := -trimpath
LDFLAGS := -ldflags "-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

PREFIX := /usr/local
DESTDIR :=

.PHONY: all build clean install uninstall test lint fmt vet deps deb

all: build

# Build the binary
build:
	go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY) ./cmd/cfimport

# Build for release (all platforms)
build-all:
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY)-linux-amd64 ./cmd/cfimport
	GOOS=linux GOARCH=arm64 go build $(GOFLAGS) $(LDFLAGS) -o dist/$(BINARY)-linux-arm64 ./cmd/cfimport

# Install to system
install: build
	install -d $(DESTDIR)$(PREFIX)/bin
	install -m 0755 $(BINARY) $(DESTDIR)$(PREFIX)/bin/$(BINARY)
	install -d $(DESTDIR)$(PREFIX)/share/man/man1
	install -m 0644 man/cfimport.1 $(DESTDIR)$(PREFIX)/share/man/man1/cfimport.1

# Uninstall from system
uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(BINARY)
	rm -f $(DESTDIR)$(PREFIX)/share/man/man1/cfimport.1

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Download dependencies
deps:
	go mod download
	go mod tidy

# Build Debian package
deb:
	dpkg-buildpackage -us -uc -b

# Clean build artifacts
clean:
	rm -f $(BINARY)
	rm -rf dist/
	rm -f coverage.out coverage.html
	go clean -cache

# Development: build and run
run: build
	./$(BINARY)

# Show help
help:
	@echo "cfimport Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make              Build the binary"
	@echo "  make build        Build the binary"
	@echo "  make build-all    Build for all platforms"
	@echo "  make install      Install to system (may need sudo)"
	@echo "  make uninstall    Remove from system"
	@echo "  make test         Run tests"
	@echo "  make lint         Run linter"
	@echo "  make fmt          Format code"
	@echo "  make deps         Download dependencies"
	@echo "  make deb          Build Debian package"
	@echo "  make clean        Clean build artifacts"
	@echo "  make help         Show this help"
