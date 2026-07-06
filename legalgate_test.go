package legalgate

import (
	"testing"
	"time"
)

func TestPolicyInForce(t *testing.T) {
	kst := time.FixedZone("KST", 9*3600)
	// A reference "now" at 2026-07-06 12:00 KST.
	now := time.Date(2026, 7, 6, 12, 0, 0, 0, kst)

	cases := []struct {
		name          string
		effectiveDate string
		now           time.Time
		want          bool
	}{
		{"past date in force", "2026-07-01", now, true},
		{"same day in force", "2026-07-06", now, true},
		{"future date not in force", "2026-08-01", now, false},
		{"malformed date fails closed", "not-a-date", now, false},
		{"empty date fails closed", "", now, false},
		{"wrong format fails closed", "2026/07/01", now, false},

		// Boundary: the gate opens at KST midnight of the effective date.
		{"one second before KST midnight", "2026-07-06",
			time.Date(2026, 7, 5, 23, 59, 59, 0, kst), false},
		{"exactly KST midnight", "2026-07-06",
			time.Date(2026, 7, 6, 0, 0, 0, 0, kst), true},

		// now given in a different zone is normalized to KST before comparing.
		{"utc now just before KST midnight", "2026-07-06",
			time.Date(2026, 7, 5, 14, 59, 59, 0, time.UTC), false}, // 23:59:59 KST
		{"utc now at KST midnight", "2026-07-06",
			time.Date(2026, 7, 5, 15, 0, 0, 0, time.UTC), true}, // 00:00:00 KST
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := PolicyInForce(c.effectiveDate, c.now); got != c.want {
				t.Errorf("PolicyInForce(%q, %v) = %v, want %v", c.effectiveDate, c.now, got, c.want)
			}
		})
	}
}
