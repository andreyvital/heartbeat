package tracker

import (
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

func New(isa IsConsideredAlive) *Tracker {
	return &Tracker{
		events:            make([]interface{}, 0),
		isConsideredAlive: isa,
	}
}

type Tracker struct {
	events            []interface{}
	isConsideredAlive IsConsideredAlive
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
	t.events = append(t.events, HostWasTracked{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}

func (t *Tracker) Ping(host string, ip ip.InternalPublicPair) {
	t.events = append(t.events, HostWasPinged{
		Host: host,
		IP:   ip,
		When: time.Now(),
	})
}
