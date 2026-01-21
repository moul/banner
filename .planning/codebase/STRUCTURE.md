# Codebase Structure

**Analysis Date:** 2026-01-21

## Directory Layout

```
banner/
├── banner.go           # Main public API: Inline() function
├── banner_test.go      # Tests for banner functionality
├── doc.go              # Package documentation and copyright
├── font.go             # Font data and letter rendering logic
├── Dockerfile          # Docker image definition
├── Makefile            # Build and development targets
├── README.md           # Project documentation
├── go.mod              # Go module definition
├── go.sum              # Go dependency checksums
├── .golangci.yml       # Linter configuration
├── .editorconfig        # Editor formatting settings
├── .gitignore          # Git ignore rules
├── .gitattributes      # Git attribute rules
├── .dockerignore        # Docker ignore rules
├── .releaserc.js        # Release configuration
├── rules.mk            # Shared build rules
├── .github/            # GitHub workflows and configuration
├── .githooks/          # Git hooks
├── AUTHORS             # Project authors
├── COPYRIGHT           # Copyright information
├── LICENSE-APACHE      # Apache 2.0 license
└── LICENSE-MIT         # MIT license
```

## Directory Purposes

**Root Directory:**
- Purpose: Flat single-package Go library
- Contains: Source code, tests, configuration, build files, and documentation
- Key files: `banner.go`, `font.go`, `banner_test.go`

## Key File Locations

**Entry Points:**
- `banner.go`: Exports public API function `Inline(input string) string`

**Configuration:**
- `go.mod`: Go module definition (moul.io/banner)
- `.golangci.yml`: Linting configuration
- `.editorconfig`: Editor formatting standards

**Core Logic:**
- `banner.go`: Main rendering logic (lines 7-41)
  - `Inline()`: Public function that orchestrates banner generation
  - `font` type: Type alias for rune-to-letter map
  - `font.Get()`: Lookup method with '?' fallback (lines 45-51)
  - `letter` type: String type for single character glyph
  - `letter.String()`: Renders letter as string
  - `letter.lines()`: Parses letter into line array

- `font.go`: Font data and structure definitions (lines 1-262)
  - `small` variable: Map of rune → letter definitions
  - Contains: ASCII art for a-z, A-Z (via lowercase), ., -, _, ?, and space characters

**Testing:**
- `banner_test.go`: Test package using package_test convention
  - Example tests: `ExampleInline()`, `ExampleInline_lowercase()`, `ExampleInline_uppercase()` (lines 10-72)
  - Unit tests: `TestInline()` (lines 74-129)
  - Test data: Inline test cases with input and expected output strings
  - Assertion pattern: Direct string comparison with `if expected != output`

**Documentation:**
- `doc.go`: Package documentation with ASCII art banner and license info
- `README.md`: Usage, installation, and project overview

## Naming Conventions

**Files:**
- Source files: Lowercase with underscores (`banner.go`, `banner_test.go`, `font.go`)
- Test files: Suffix with `_test.go` (standard Go convention)

**Functions:**
- Public: PascalCase (`Inline`, `Get`, `String`)
- Private: None in this codebase (all at package scope)

**Types:**
- PascalCase: `font`, `letter` types exported at package level
- Convention: Single-letter type names for simple wrapper types

**Variables:**
- Package-level data: Lowercase (`small` - the embedded font map)

**Constants:**
- None defined in this codebase (character data embedded in map)

## Where to Add New Code

**New Function (e.g., Multi-line render):**
- Primary code: `banner.go` (alongside `Inline()`)
- Tests: `banner_test.go` (add to `TestInline()` or create `TestNewFunction()`)
- Pattern: Follow existing style - receive input, return string, handle edge cases

**New Font:**
- Implementation: `font.go` (add new map variable like `small`)
- Export: Add public variable if should be user-accessible
- Usage: Would require new public function in `banner.go` to use different font
- Structure: Match format of `small` map: rune keys, letter values with @\n delimiters

**Helper Functions:**
- Location: `banner.go` or dedicated file if substantial
- Scope: Keep internal to package (lowercase names) unless exposing new capability
- Pattern: Should support text-to-banner pipeline

## Special Directories

**.github/:**
- Purpose: GitHub-specific configuration
- Generated: No
- Committed: Yes
- Contents: CI/CD workflows, pull request templates, action configurations

**.githooks/:**
- Purpose: Git hooks for local development
- Generated: No
- Committed: Yes
- Contents: Custom hooks to run during git operations

**build outputs (not in repo):**
- Purpose: Compiled binaries and artifacts
- Generated: Yes (during build)
- Committed: No (.gitignore excludes)

## File Relationships

**Dependency Graph:**
```
banner.go (main API)
  ├── font.go (font data and types)
  │   └── (no dependencies)
  └── (no other internal dependencies)

banner_test.go (tests)
  ├── imports "moul.io/banner"
  └── tests banner.go
```

**Module Boundary:**
- Single package `banner` (no subpackages)
- Entire functionality accessible via `moul.io/banner` import
- No internal/ directory (all code exported at package level)

## Import Organization

**In banner.go:**
```go
import (
	"strings"
)
```
- Only standard library import needed (strings for manipulation)

**In banner_test.go:**
```go
import (
	"fmt"
	"testing"
	"moul.io/banner"
)
```
- Standard library: `fmt`, `testing`
- Local package: `moul.io/banner`

---

*Structure analysis: 2026-01-21*
