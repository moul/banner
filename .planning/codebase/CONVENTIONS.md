# Coding Conventions

**Analysis Date:** 2026-01-21

## Naming Patterns

**Files:**
- Lowercase with underscores for test files: `*_test.go`
- Package documentation file: `doc.go`
- Descriptive names for feature files: `banner.go`, `font.go`

**Functions:**
- PascalCase for exported functions: `Inline()`, `Get()`, `String()`
- camelCase for unexported functions: `lines()`

**Variables:**
- camelCase for variable names: `input`, `lines`, `height`, `letter`
- Short names for common loop variables: `r` for rune, `i` for index
- Descriptive names for struct fields and map keys

**Types:**
- PascalCase for exported types: `N/A` (no exported types in this codebase)
- lowercase for unexported types: `font`, `letter`

## Code Style

**Formatting:**
- Tab indentation (configured in `.editorconfig`)
- UTF-8 charset
- Unix line endings (LF)
- Final newline in all files
- Trim trailing whitespace

**Linting:**
- Tool: `golangci-lint`
- Config: `.golangci.yml`
- Run with: `make lint` (via `go.lint` target in `rules.mk`)
- Key enabled linters: `gofmt`, `goimports`, `govet`, `errcheck`, `funlen`, `gocyclo`, `staticcheck`, `unused`, `unparam`, `nakedret`
- Deadline: 1 minute
- Tests are skipped during linting
- Files `*_test.go`, `*.pb.go`, `*.gen.go` are excluded

## Import Organization

**Order:**
1. Standard library imports
2. Third-party imports
3. Local package imports

**Examples from codebase:**
- Standard library: `"strings"`
- Local imports: `"moul.io/banner"`

**Path Aliases:**
- Uses full import path: `moul.io/banner`

## Error Handling

**Patterns:**
- No explicit error handling in public functions (functions handle edge cases with return values)
- Return empty string for invalid input in `Inline()` function
- Return default character from font map if character not found: `f['?']` fallback in `Get()` method

**Example from `banner.go`:**
```go
func (f font) Get(key rune) letter {
	letter, found := f[key]
	if found {
		return letter
	}
	return f['?']  // Fallback to '?' character
}
```

## Logging

**Framework:** Standard library only (`fmt` for output in examples/tests)

**Patterns:**
- No logging in the library code itself
- Tests use `t.Log()` and `t.Errorf()` for output and error reporting
- Example functions use `fmt.Println()` for documentation output

## Comments

**When to Comment:**
- Document the purpose of types and functions
- Explain design decisions (e.g., font selection rationale in `font.go`)
- Include copyright and license headers in `doc.go`

**JSDoc/TSDoc:**
- Not applicable (Go uses different doc comment style)
- Package documentation in `doc.go` with `// Package banner` convention

**Go Doc Comments:**
- Public identifiers documented with comments starting with the identifier name
- Example from `banner.go`: Functions like `Inline()` have single-line descriptions
- Unexported functions documented with lowercase comments

## Function Design

**Size:**
- Functions are kept small and focused (most functions under 20 lines)
- Example: `Inline()` is 34 lines with clear logic flow

**Parameters:**
- Minimal parameters (typically 1-2)
- Examples: `Inline(input string)`, `Get(key rune)`

**Return Values:**
- Single return value for most functions
- Methods return either the value or a zero value (empty string)
- No error returns in public functions (handles edge cases gracefully)

**Example from `banner.go`:**
```go
func Inline(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return ""
	}
	// ... implementation
	return strings.Join(lines, "\n")
}
```

## Module Design

**Exports:**
- Only `Inline()` function is exported for the public API
- All types and helper functions are unexported (`font`, `letter`)
- Private types handle internal data structures

**Barrel Files:**
- Single package module (`banner` package)
- No barrel file pattern used
- Entry point: `doc.go` with package declaration and documentation

## Struct and Type Design

**Unexported Types as Maps:**
- `font` type is a map of `rune` to `letter` values
- `letter` type is a string alias with methods
- Internal implementation details kept private

**Method Receivers:**
- Receiver types are value receivers for small types
- Example: `func (l letter) lines() []string`
- Example: `func (f font) Get(key rune) letter`

---

*Convention analysis: 2026-01-21*
