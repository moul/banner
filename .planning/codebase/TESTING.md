# Testing Patterns

**Analysis Date:** 2026-01-21

## Test Framework

**Runner:**
- Go's built-in `testing` package
- Config: `rules.mk` provides test command with `go.unittest` target
- Go version: 1.13+

**Assertion Library:**
- Manual assertion (no external assertion library)
- Uses `testing.T` methods: `t.Run()`, `t.Log()`, `t.Errorf()`

**Run Commands:**
```bash
make unittest              # Run all tests with coverage
make test                  # Run tests with linting
go test ./...              # Standard go test
go test -v ./...           # Verbose output
```

## Test File Organization

**Location:**
- Co-located with implementation
- Test file in same directory as code: `banner_test.go` alongside `banner.go`

**Naming:**
- Standard Go convention: `*_test.go` suffix
- Package: `banner_test` (separate test package for integration-style testing)

**Structure:**
```
/media/psf/p/moul.io/banner/
├── banner.go              # Main implementation
├── banner_test.go         # Tests for Inline function
├── font.go                # Font data
└── doc.go                 # Package documentation
```

## Test Structure

**Suite Organization:**
```go
func TestInline(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Test cases
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			// Test logic
		})
	}
}
```

**Patterns Observed:**

1. **Table-Driven Tests:** All test cases defined in a slice of structs at the beginning
   - Each test case has `input` and `expected` fields
   - Iteration with `for _, test := range tests`
   - Sub-tests with `t.Run()` for individual test cases

2. **Setup/Teardown:** None required for this package (no state management)

3. **Assertion Pattern:**
```go
output := banner.Inline(test.input)
expected := test.expected[1 : len(test.expected)-1]
if expected != output {
	t.Log("output: \n" + output)
	t.Log("expected: \n" + expected)
	t.Errorf("output differs")
}
```

## Example Tests

**Documentation Examples:**
- `ExampleInline()` - Basic usage example
- `ExampleInline_lowercase()` - Demonstrates lowercase letters
- `ExampleInline_uppercase()` - Demonstrates uppercase letters
- Format: Standard Go example functions with `// Output:` comment block
- Examples are verified by the test runner

**Example Structure:**
```go
func ExampleInline() {
	fmt.Println("start of banner")
	fmt.Println(banner.Inline("hey world."))
	fmt.Println("end of banner")
	// Output:
	// start of banner
	// [expected output lines]
	// end of banner
}
```

## Mocking

**Framework:** Not used (no external dependencies to mock)

**Patterns:**
- No mocking needed - this is a pure string transformation library
- All dependencies are from the standard library (`strings`)

**What to Mock:**
- Not applicable for this codebase

**What NOT to Mock:**
- Standard library functions

## Fixtures and Factories

**Test Data:**
- Inline test data in test table structures
- Example edge cases: empty strings, single characters, special characters (`@`, `!`, `?`)

**Test Cases in `TestInline()`:**
```go
tests := []struct {
	input    string
	expected string
}{
	{"jjj", `[expected banner]`},
	{"j j", `[expected banner with space]`},
	{"j", `[single character]`},
	{"@?!", `[special characters]`},
	{"ccc", `[repeated character]`},
	{" ", `[space only]`},
	{"", `[empty string]`},
}
```

**Location:**
- Test data defined inline in test functions
- No separate fixture files

## Coverage

**Requirements:** No explicit coverage target set

**View Coverage:**
```bash
make go.coverfunc         # Show coverage by function
```

**Coverage Details:**
- `rules.mk` provides atomic coverage mode with `-covermode=atomic`
- Coverage file generated to `./coverage.txt`
- Atomic mode for race-condition-safe reporting
- `-race` flag enabled during test runs

**Analyze Coverage:**
- Run `go test -cover ./...` to see coverage percentage
- Use `go tool cover -html=coverage.txt` to view HTML report

## Test Types

**Unit Tests:**
- `TestInline()` - Tests the primary `Inline()` function
- Scope: Single function with various input combinations
- Approach: Table-driven tests with edge cases and normal cases

**Integration Tests:**
- Example tests (`ExampleInline`, `ExampleInline_lowercase`, `ExampleInline_uppercase`)
- Verifies behavior from user perspective
- Tests are part of documentation

**E2E Tests:**
- Not used (library has simple, focused API)

## Common Patterns

**Passing Tests:**
```go
t.Run(test.input, func(t *testing.T) {
	output := banner.Inline(test.input)
	expected := test.expected[1 : len(test.expected)-1]
	if expected != output {
		t.Log("output: \n" + output)
		t.Log("expected: \n" + expected)
		t.Errorf("output differs")
	}
})
```

**Edge Case Testing:**
- Empty string input: `{"", ""}`
- Whitespace only: `{" ", ""}`
- Single character: `{"j", "[banner]"}`
- Multiple characters with spaces: `{"j j", "[banner]"}`
- Special characters: `{"@?!", "[banner]"}`

**Comparison Pattern:**
- String equality comparison with `==`
- No fuzzy matching or partial verification
- Full output comparison against expected result

## Test Exclusion

**Files Skipped During Linting:**
- `testing.go` (if present) - configuration only

**Test Files Not Excluded:**
- `*_test.go` files are included in test runs
- Linting is skipped for test files during `golangci-lint` execution

---

*Testing analysis: 2026-01-21*
