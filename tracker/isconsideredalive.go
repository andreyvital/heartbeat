package tracker

import "time"

// IsConsideredAlive is a interface to check whether the host is considered
// alive by the given last seen time
type IsConsideredAlive func(host string, seen time.Time) bool
