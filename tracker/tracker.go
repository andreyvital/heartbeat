package tracker

import (
	"sync"
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

func New(isa IsConsideredAlive) *Tracker {
	return &Tracker{
		make([]interface{}, 0),
		isa,
		&sync.Mutex{},
	}
}

type Tracker struct {
	events            []interface{}
	isConsideredAlive IsConsideredAlive
	*sync.Mutex
}

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

func (t *Tracker) Track(host string, ip ip.InternalPublicPair) {
	t.Lock()
	defer t.Unlock()

	t.events = t.collectGarbage(t.events)

	t.events = append(t.events, HostWasTracked{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}

func (t *Tracker) Ping(host string, ip ip.InternalPublicPair) {
	t.Lock()
	defer t.Unlock()

	t.events = t.collectGarbage(t.events)

	t.events = append(t.events, HostWasPinged{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}

func (t *Tracker) collectGarbage(events []interface{}) (keep []interface{}) {
	for _, event := range events {
		switch event.(type) {
		case HostWasTracked:
			keep = append(keep, event)
		case HostWasPinged:
			e := event.(HostWasPinged)

			if t.isConsideredAlive(e.Host, e.When) {
				keep = append(keep, event)
			}
		}
	}

	return keep
}
