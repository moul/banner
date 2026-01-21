# Technology Stack

**Analysis Date:** 2026-01-21

## Languages

**Primary:**
- Go 1.13+ - Core library code for ASCII-art banner generation

## Runtime

**Environment:**
- Go (golang) - Cross-platform runtime
- Alpine Linux 3.11 - Production Docker runtime

**Package Manager:**
- Go Modules (go mod) - Dependency management
- Lockfile: Present (`go.sum` - currently empty, no external dependencies)

## Frameworks

**Core:**
- Standard Library only - No external frameworks required

**Testing:**
- Go's built-in testing package - Standard Go test framework
- Example-based tests (ExampleXxx pattern) - Documentation tests

**Build/Dev:**
- make - Build automation
- golangci-lint - Code linting and analysis
- Docker - Containerization

## Key Dependencies

**Critical:**
- None - This is a pure Go standard library project with zero external dependencies

**Infrastructure:**
- None required - No database, cache, or infrastructure dependencies

## Configuration

**Environment:**
- No runtime configuration required
- Project uses go.mod for version specification

**Build:**
- `Makefile` - Build orchestration, includes rules.mk
- `rules.mk` - Common Makefile rules for Go projects
- `.golangci.yml` - Linter configuration
- `Dockerfile` - Multi-stage Docker build (golang:1.14-alpine as builder, alpine:3.11 runtime)
- `.editorconfig` - Editor formatting standards (tabs for Go, spaces for others)

## Platform Requirements

**Development:**
- Go 1.11+ (CI tests against 1.11, 1.12, 1.13, 1.14)
- make
- git

**Production:**
- Alpine Linux 3.11 (or any system where binary can run)
- Optional: Docker

## Build Artifacts

**Binary:**
- Output: `/go/bin/banner` (Dockerfile target)
- Size: ~38KB with default build options, ~12KB with upx + strip
- Language: Pure compiled Go binary

---

*Stack analysis: 2026-01-21*
