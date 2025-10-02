# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```

***Questions for PM Timothy***:
1) Should letter grades use exact cutoffs on the unrounded final numeric grade (89.999 = B, 90.0 = A), and what rounding is required for display vs. logic?
2) If any category (essays, for example) has no grades, do we treat its average as empty/ignored or as 0?
3) Do all the weights have to add up to 100%? How should we report if they don’t?
