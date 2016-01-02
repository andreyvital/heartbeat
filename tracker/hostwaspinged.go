package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

// HostWasPinged represents a ping event from a host
type HostWasPinged struct {
	Host string
	IP   ip.InternalPublicPair
	When time.Time
}
