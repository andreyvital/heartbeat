package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

// New creates a new tracker using the specified function to determine
// if a host is still alive
func New(isa IsConsideredAlive) *Tracker {
	return &Tracker{
		events:            make([]interface{}, 0),
		isConsideredAlive: isa,
	}
}

// Tracker provides a simple API to track hosts
type Tracker struct {
	events            []interface{}
	isConsideredAlive IsConsideredAlive
}

// IsHostBeingTracked tells if a host is being tracked
func (t *Tracker) IsHostBeingTracked(host string) bool {
	for _, event := range t.events {
		switch event.(type) {
		case HostWasTracked:
			if event.(HostWasTracked).Host == host {
				return true
			}
		}
	}

	return false
}

// ToPossiblyAlive returns a map of hosts which are possibly alive
func (t *Tracker) ToPossiblyAlive() map[string]ip.InternalPublicPair {
	m := make(map[string]ip.InternalPublicPair, 0)

	for _, event := range t.events {
		switch event.(type) {
		case HostWasTracked:
			e := event.(HostWasTracked)

			if t.isConsideredAlive(e.Host, e.When) {
				m[e.Host] = e.IP
			}
		case HostWasPinged:
			e := event.(HostWasPinged)

			if t.isConsideredAlive(e.Host, e.When) {
				m[e.Host] = e.IP
			}
		}
	}

	return m
}

// CountOfTracked returns the count of hosts being tracked
func (t *Tracker) CountOfTracked() int {
	count := 0

	for _, event := range t.events {
		switch event.(type) {
		case HostWasTracked:
			count++
		}
	}

	return count
}

// Track starts tracking a host with the following IP pair
func (t *Tracker) Track(host string, ip ip.InternalPublicPair) {
	t.events = append(t.events, HostWasTracked{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}

// Ping registers a ping event from the host with given IP pair
func (t *Tracker) Ping(host string, ip ip.InternalPublicPair) {
	t.events = append(t.events, HostWasPinged{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}
