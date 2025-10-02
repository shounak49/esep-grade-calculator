## Part 2 — Understanding a New Project

### 2.1 Tech stack used in this repo
- Language: Go (Golang), using Go modules via go.mod
- Build/Run: go build, go test
- Test framework: Go’s standard testing package
- Format/Lint: go fmt; optional go vet
- Package name: esepunittests

### 2.2 Brief summary of each file (4 total)
- go.mod — Declares module; enables Go modules.
- README.md — Setup/usage notes
- grade_calculator.go — Library code for adding grades and computing weighted letter grades (A/B/C/D/F) using 50/35/15.
- grade_calculator_test.go — Unit tests for expected letters given sample inputs (A, B, F cases).

### 2.3 Are there tests in the repository?
Yes.

1) What do they do?
They build a GradeCalculator, add grades across categories, call GetFinalGrade(), and assert the letter grade (“A”, “B”, “F”).

2) How do you run them?
go test ./...
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt

### 2.4 Resources used to understand the code
- https://go.dev/doc/modules
- https://go.dev/doc/tutorial/add-a-test
- https://go.dev/blog/cover
