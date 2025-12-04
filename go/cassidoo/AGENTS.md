# Agent Guidelines for Cassidoo Challenge Repository

## Commands
- Run all tests: `go test ./...`
- Run single test: `cd <challenge_dir> && go test -run <TestName>`
- Run from root: `go test ./<challenge_dir> -run <TestName>`
- Format code: `gofmt -w .`

## Workflow
- **AI Solution Creation**: When asked to create an AI solution, copy `solution_test.go` to `solution_ai_test.go` and append `AI` to all tested function names (e.g., `SortHtmlColors` → `SortHtmlColorsAI`)
- **File Access Restrictions**: Never read `solution.go` unless explicitly asked to compare solutions
- **Allowed Files**: Can read `QUESTION.md`, `types.go`, `utils.go`, test files, and any other supporting files

## Structure
- Each challenge in its own directory: `YYYY_MM_DD_challenge_name/`
- Files: `QUESTION.md`, `solution.go`, `solution_test.go`, optionally `solution_ai.go`, `solution_ai_test.go`, `types.go`, `utils.go`
- Module: `cassidoo` with Go 1.25.1

## Code Style
- Package name: camelCase matching challenge name (e.g., `sortHtmlColors`)
- Imports: Standard library first, grouped with empty line before external imports
- Formatting: Use `gofmt` (some files may not be formatted yet)
- Naming: Exported functions use PascalCase, private use camelCase
- Types: Define in separate `types.go` if multiple types needed
- Error handling: Return errors explicitly (though many solutions don't use errors)
- Comments: Add inline comments for algorithm explanations
- Testing: Use custom assertions from `cassidoo/utils` package (`AssertDeepEquals`, `AssertIsPositive`, `AssertIsNegative`)
- Test structure: Use `t.Run()` for subtests with descriptive names
