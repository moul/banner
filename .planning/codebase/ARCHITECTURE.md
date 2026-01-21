# Architecture

**Analysis Date:** 2026-01-21

## Pattern Overview

**Overall:** Single-package library with data-driven rendering

**Key Characteristics:**
- Pure Go library with no external dependencies
- Flat single-package structure (`banner` package)
- Renders ASCII art text by composing pre-defined letter glyphs
- Data-driven: font data embedded as compiled maps
- Functional transformation pipeline: input text → letter glyphs → assembled lines → trimmed output

## Layers

**Public API Layer:**
- Purpose: Expose ASCII banner generation to consumers
- Location: `banner.go`
- Contains: `Inline(input string) string` function
- Depends on: Font data (small font map)
- Used by: External packages and applications

**Font Data Layer:**
- Purpose: Store ASCII character definitions and font rendering logic
- Location: `font.go`
- Contains: `font` type (map[rune]letter), `letter` type with rendering methods
- Depends on: Nothing (self-contained)
- Used by: Public API layer

**Data Structures:**
- `font`: A map from rune (character) to letter definition, with lookup fallback to '?' character
- `letter`: A string type representing the ASCII art for a single character, with methods to convert to lines

## Data Flow

**Text Rendering Pipeline:**

1. **Input Validation:** `Inline(input string)` receives raw text
2. **Trim Input:** Strip whitespace from input string
3. **Handle Empty:** Return empty string if input becomes empty
4. **Get First Letter:** Retrieve ASCII art for first character using `small.Get(rune(input[0]))`
5. **Initialize Lines:** Extract lines from first letter glyph
6. **Iterate Characters:** For remaining characters in input:
   - If space: append two spaces to each line
   - If other: fetch letter glyph, append to each line
7. **Trim Whitespace:** Remove trailing spaces from each line
8. **Remove Empty Lines:** Remove trailing empty line if present, remove leading empty line if present
9. **Join Lines:** Concatenate all lines with newlines to create final output

**State Management:**
- Stateless function: no mutable state, no side effects
- Lines array is built up progressively during character iteration
- All transformations are pure (same input always produces same output)

## Key Abstractions

**Font Map (font type):**
- Purpose: Provides character lookup with fallback behavior
- Examples: `small` variable in `font.go` contains all ASCII letter definitions
- Pattern: Map-based registry with fallback to placeholder ('?') character for unknown glyphs
- Implementation: `Get(key rune) letter` method checks map, returns default '?' if not found

**Letter Type (letter type):**
- Purpose: Represents single ASCII art character with rendering logic
- Pattern: String with embedded line separator (@\n) that can be split into lines
- Methods:
  - `String()`: Renders letter as newline-separated string
  - `lines()`: Parses letter string into array of strings (splits on @\n)
- Data Format: Starts and ends with markers, contains lines separated by @\n

**Horizontal Composition:**
- Pattern: Build output line-by-line across all input characters
- Approach: Maintain array of output lines, concatenate character renderings to each line
- Spacing: Double space between characters (via space handling)

## Entry Points

**Inline Function:**
- Location: `banner.go`, line 7
- Triggers: Called by external code passing text string
- Responsibilities:
  - Validates and transforms input
  - Orchestrates character lookup and line composition
  - Cleans up output (trims empty lines, trailing spaces)
  - Returns single string with complete banner

## Error Handling

**Strategy:** No error returns - graceful degradation via fallback

**Patterns:**
- Unknown characters: Lookup falls back to '?' (question mark) character using `font.Get()` method
- Empty input: Returns empty string without error
- Edge cases: Leading/trailing empty lines automatically removed to maintain clean output
- No validation errors: Function accepts any string

## Cross-Cutting Concerns

**Logging:** None - pure library with no logging

**Validation:** Input trimming and empty-check performed at start of `Inline()` function

**Character Support:** Limited to defined glyphs (a-z, A-Z, ., -, _, ?, and space); others map to '?'

---

*Architecture analysis: 2026-01-21*
