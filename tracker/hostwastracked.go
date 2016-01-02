package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

// HostWasTracked represents an event when the host was tracked
type HostWasTracked struct {
	Host string
	IP   ip.InternalPublicPair
	When time.Time
}
