package tracker

import "time"

type IsConsideredAlive func(host string, seen time.Time) bool
