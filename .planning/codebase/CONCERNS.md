# Codebase Concerns

**Analysis Date:** 2026-01-21

## Known Bugs

**Example test failure for uppercase characters:**
- Symptoms: `ExampleInline_uppercase` test fails with actual output showing `?` characters (question marks) instead of uppercase letter glyphs
- Files: `banner_test.go`, `banner.go`, `font.go`
- Trigger: Call `banner.Inline()` with uppercase letters (A-Z). The function outputs `?` character glyphs instead
- Root Cause: Font definition in `font.go` only contains lowercase letters (a-z), not uppercase (A-Z). The `Inline()` function at line 13 calls `small.Get(rune(input[0]))` which will fall back to the `?` (question mark) glyph for any character not in the font map (line 50 in `banner.go`)
- Current Status: Test added in commit `ddd5487` but marked as expected to fail. Commit `d354fbd` merged this into master despite the failure
- Impact: Users cannot render uppercase letters; library only handles lowercase a-z, plus `.`, `-`, `_`, and `?`

## Incomplete Character Support

**Limited font coverage:**
- What's Missing: Uppercase letters (A-Z) and other ASCII characters
- Files: `font.go` (lines 50-261)
- Current Support: Only a-z (26 chars) + `.` + `-` + `_` + `?` (4 special) = 31 total characters
- Workaround: Users can input lowercase text or special characters only
- README Limitation: Line 22 correctly documents "incomplete alphabet (a-zA-Z.-_?)" but feature is actually only a-z.-_?
- Impact: Library cannot be used for full ASCII text rendering, limiting real-world usage significantly

## Unsafe String Indexing

**Potential panic on short input strings:**
- Issue: Line 60 in `banner.go` performs `string(l[1 : len(l)-2])` without bounds checking
- Files: `banner.go` (line 60)
- Risk: If a letter string in font map is shorter than 3 characters (needs at least index 0, 1, and len(l)-2), this will panic
- Current Mitigation: Font definitions in `font.go` all follow consistent pattern with proper length
- Fragility: Adding new fonts without proper validation would be error-prone
- Safe Modification: Always validate that letter strings meet minimum length requirements before indexing

## Test Coverage Gaps

**Critical path tests missing:**
- What's not tested: Behavior with unsupported characters; edge cases with single-character input; consecutive spaces; very long input strings
- Files: `banner_test.go`
- Current Tests: Only example tests and basic `TestInline` with 6 test cases (lines 74-129)
- Risk: Regressions in edge cases could go unnoticed
- Gap Details:
  - No test for uppercase input validation (test exists but fails at line 49)
  - No test for mixed case input
  - No test for numeric characters (0-9)
  - No test for very long strings (performance/memory)
  - No negative tests for invalid characters

## Architecture Concerns

**Hard-coded font data:**
- Problem: Font definitions baked into source code as Go constants in `font.go`
- Files: `font.go` (261 lines of hard-coded font data)
- Impact: Adding new fonts requires source code modification; no runtime font loading capability; no plugin/extension system
- Scaling Issue: Only one font supported; library design prevents easy font extensibility

## Go Version Targeting

**Outdated Go version requirement:**
- Current: `go.mod` specifies `go 1.13` (line 3)
- Issue: Go 1.13 was released August 2019; now end-of-life for ~5 years
- Files: `go.mod`
- Impact: No access to modern Go features; security fixes in Go runtime may not apply; developers expect minimum Go 1.18+
- Recommendation: Update to at least `go 1.19` or higher to align with industry standards

## Docker Image Concerns

**Outdated base images:**
- Issue: Dockerfile uses `golang:1.14-alpine` (line 7) and `alpine:3.11` (line 17)
- Files: `Dockerfile`
- Version Age: Go 1.14 (March 2020) and Alpine 3.11 (November 2019) are both 5+ years old
- Risk: Known vulnerabilities in old Go and Alpine versions; no security patches available
- Impact: Docker images built with this Dockerfile contain unpatched security issues
- Current Status: Renovate has attempted updates (commits show docker tags updated), but base images in Dockerfile remain old

## Deprecated Linters

**Golangci-lint configuration outdated:**
- Issue: `.golangci.yml` enables deprecated linters that have been removed from golangci-lint
- Files: `.golangci.yml` (lines 20-55)
- Examples: `golint` (line 37) and `interfacer` (line 41) are removed in golangci-lint v1.49.0+
- Impact: Linting will fail on modern versions of golangci-lint; CI/CD will break
- Current Mitigation: CI may be pinned to older golangci-lint version (not checked in codebase)

## String Trimming Logic

**Fragile whitespace handling:**
- Issue: Lines 31-39 in `banner.go` trim whitespace from output
- Files: `banner.go` (lines 31-39)
- Concern: Logic trims trailing spaces from each line (31-33) and then removes entirely-empty first and last lines (34-39)
- Fragility: This assumes consistent formatting from font definition; any font with intentional leading/trailing spaces would be corrupted
- Test Coverage: Not explicitly tested for fonts with intentional spacing

## Character Mapping Case Sensitivity

**No case normalization:**
- Issue: `Inline()` function uses rune values directly without normalization
- Files: `banner.go` (lines 13, 23)
- Problem: Uppercase input falls through to default `?` character because font map only has lowercase keys
- Workaround: Explicit case conversion required by caller, not handled by library
- Better Approach: Library could offer case-insensitive option or auto-lowercase internally
