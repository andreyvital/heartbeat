package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

type HostWasTracked struct {
	Host string
	IP   ip.InternalPublicPair
	When time.Time
}
