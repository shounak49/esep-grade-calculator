## 5.1 Can you achieve 100% code coverage without testing all cases? Why?
**Yes.** Statement/line coverage only shows that lines executed, but it does not show that behavior was verified or that important edges were exercised. You can hit 100% by calling each function while still missing:
- Boundary cutoffs
- Error/empty-path behavior
- Alternative branches and conditions

**High coverage ≠ high confidence.** It is necessary to a degree but it is sometimes not enough.

## 5.2 Comment out the assertions—what happens to coverage?
Coverage usually remains **the same** (still very high). The code still runs so lines are “covered,” even though nothing checks the results. This demonstrates why coverage must be paired with strong assertions and good test design.

## 5.3 How would you address this in a team setting?
- **Mutation testing:** require a minimum mutation score so tests must fail when behavior changes.
- **Branch/condition coverage:** target more than statements; verify both sides of important conditions.
- **Assertion quality:** enforce at least one meaningful assert per test; avoid no-op tests.
- **Design for testability:** keep logic pure/small; separate I/O from computation.