## Part 4 — Adding Tests

### 4.1 Coverage after fixes
- Command: `go test -coverprofile="coverage.out" ./...`
- Coverage result: 
PS C:\Users\sjosh\esep-grade-calculator\grade-calculator> go test -coverprofile="coverage.out" ./...
ok      esep/grade-calculator   0.449s  coverage: 100.0% of statements

### 4.2 New tests added to reach 100% coverage
- Added explicit **C** and **D** cases (exact cutoffs at 70 and 60).
- Boundary checks at **60/70/80/90**.
- Multiple items per category to verify the **averaging** logic.
- **Missing categories count as 0**: only assignments, only exams, only essays, and **no grades** cases.
- Covered the `GradeType.String()` mapping.

### 4.3 Refactor thoughts
- Keep calculation as **pure functions**; compare raw numeric, no rounding in logic.
- **Separate data collection** (`AddGrade`) from **computation** (`calculateNumericalGrade`).
- Make the **missing-category policy** configurable.
- Prefer **small, single-purpose helpers** with clear return types. This will help with coverage.
