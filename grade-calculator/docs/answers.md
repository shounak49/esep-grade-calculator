## Part 3 — Updating Code and Tests

### 3.1 Diagnosis
- Wrong assertion: TestGetGradeF expected "F" while inputs (100,95,91) return "A". This was changed with 50/35/15 weights.
- Code changes:Code bugs from starter: essays averaged from wrong slice; average loop summed index not value; potential divide-by-zero; numeric truncated.

### 3.2 Fixes
- Implemented float64 numeric pipeline with weights 50/35/15; corrected essays and averaging; safe handling
- Updated TestGetGradeF inputs to yield F; A/B tests unchanged.
