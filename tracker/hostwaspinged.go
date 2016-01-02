package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

type HostWasPinged struct {
	Host string
	IP   ip.InternalPublicPair
	When time.Time
}
