// Package legalgate is a tiny, fail-closed "is this policy in force yet?" gate.
//
// When a service announces an adverse policy change (a new data-retention or
// account-deletion rule, say) with a future effective date, the code path that
// enforces the new rule must stay gated until that date arrives — and must fail
// closed (treat the policy as not yet in force) when the configured date is
// malformed, so a typo can never silently enforce a policy early.
package legalgate

import "time"

// PolicyInForce reports whether the policy with the given effective date
// (YYYY-MM-DD, interpreted at KST midnight) is in force at now.
//
// A malformed effectiveDate returns false (fail-closed): the announced policy
// is treated as not yet in force rather than risk enforcing it early.
func PolicyInForce(effectiveDate string, now time.Time) bool {
	kst := time.FixedZone("KST", 9*3600)
	d, err := time.ParseInLocation("2006-01-02", effectiveDate, kst)
	if err != nil {
		return false // fail-closed
	}
	return !now.In(kst).Before(d)
}
