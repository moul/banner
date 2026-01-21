# External Integrations

**Analysis Date:** 2026-01-21

## APIs & External Services

**None:**
- This is a pure library with no external API integrations
- No HTTP clients or API SDKs used
- No webhook endpoints

## Data Storage

**Databases:**
- Not applicable - No database connectivity

**File Storage:**
- Local filesystem only - Embedded font data in source code (`font.go`)

**Caching:**
- Not applicable - No caching layer

## Authentication & Identity

**Auth Provider:**
- Not applicable - No authentication required

## Monitoring & Observability

**Error Tracking:**
- None - Standard Go error returns

**Logs:**
- Standard output (fmt package) - No structured logging framework
- Example: `fmt.Println()` in tests

**Metrics:**
- Not applicable - No metrics collection

## CI/CD & Deployment

**Hosting:**
- GitHub (source repository)
- GitHub Releases (artifact distribution)

**CI Pipeline:**
- GitHub Actions (`.github/workflows/`)
- Pipeline stages:
  - `golangci-lint` - Static analysis (v1.28)
  - Cross-platform testing (Windows, macOS, Linux)
  - Go versions tested: 1.11, 1.12, 1.13, 1.14
  - Semantic Release - Automated versioning and release on master

**Docker Registry:**
- Docker Hub (`moul/banner`) - Containerized binary distribution

## Environment Configuration

**Required env vars:**
- None - No environment variables needed for library operation

**Build-time env vars:**
- `GO111MODULE=on` - Enable Go modules in Dockerfile
- Docker build args: `BUILD_DATE`, `VCS_REF`, `VERSION`

**Secrets location:**
- None - No secrets management needed

## Webhooks & Callbacks

**Incoming:**
- Not applicable - Library has no webhook endpoints

**Outgoing:**
- GitHub Actions triggers on:
  - Push to master branch
  - Pull requests
  - Version tags (v*)

## External Dependencies Status

**go.sum:**
- Empty - No production dependencies required
- Pure standard library implementation

## Distribution Channels

**Package Distribution:**
- Go: `go get -u moul.io/banner`
- GitHub Releases: Binary downloads
- Docker Hub: `docker run moul/banner`

---

*Integration audit: 2026-01-21*
