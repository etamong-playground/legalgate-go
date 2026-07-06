# legalgate-go

> **About** — One of several shared libraries/tools behind a personal homelab "fleet" of
> small apps. Published to show the **design decisions** behind these cross-cutting concerns.
> Authored and maintained with [Claude Code](https://www.anthropic.com/claude-code) (Anthropic's
> agentic CLI), not hand-written.
>
> **This is a public repository** — keep internal infrastructure details (hostnames, secret
> paths, private URLs, internal issue/MR references) out of code, comments, and commit messages.

A tiny, **fail-closed** "is this policy in force yet?" gate for Go services.

When a service announces an adverse policy change — a new data-retention rule, an
account-deletion policy — with a *future* effective date, the code path that enforces the new
rule has to stay behind a gate until that date arrives. Two properties matter:

- **Time-zone honesty.** The effective date is a calendar date (`YYYY-MM-DD`) that means
  midnight in the operator's timezone (KST), not UTC. Comparing a bare date against
  `time.Now()` in the wrong zone shifts the cutover by up to a day.
- **Fail-closed.** A malformed or missing effective date must be treated as *not yet in force*
  — never enforce an announced-but-not-yet-effective policy because of a config typo.

Collapsing that into one audited function keeps every service in the fleet gating grace-vs-enforce
transitions identically.

## Usage

```go
import (
	"time"

	"github.com/etamong-playground/legalgate-go"
)

if legalgate.PolicyInForce(cfg.DeletionPolicyEffectiveDate, time.Now()) {
	// the new policy is in force — enforce it (e.g. de-identify)
} else {
	// still inside the announced grace window — keep the old behavior
}
```

`effectiveDate` is `YYYY-MM-DD`, interpreted at **KST midnight**. A value that fails to parse
returns `false` (fail-closed).

## Acknowledgements

No third-party dependencies — standard library `time` only.

## License

MIT — see [LICENSE](LICENSE).
