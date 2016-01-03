package tracker_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
	"github.com/CentaurWarchief/heartbeat/tracker"
	"github.com/stretchr/testify/assert"
)

func TestTrackerTrack(t *testing.T) {
	tracker := tracker.New(func(host string, seen time.Time) bool {
		return true
	})

	assert.Equal(t, 0, tracker.CountOfTracked())

	tracker.Track("localhost-localdomain", ip.NewInternalPublicPair(
		"192.168.55.55",
		"131.131.131.131",
	))

	assert.Equal(t, 1, tracker.CountOfTracked())

	tracker.Track("localhost-localdomain", ip.InternalPublicPair{
		Internal: "192.168.55.55",
	})

	assert.Equal(t, 2, tracker.CountOfTracked())

	tracker.Track("localhost-localdomain", ip.InternalPublicPair{
		Internal: "192.168.55.55",
	})

	assert.Equal(t, 3, tracker.CountOfTracked())
}

func TestPossiblyAlive(t *testing.T) {
	tracker := tracker.New(func(host string, seen time.Time) bool {
		return true
	})

	tracker.Track("localhost-localdomain", ip.InternalPublicPair{
		Internal: "192.168.55.55",
	})

	alive := tracker.ToPossiblyAlive()

	assert.Len(t, alive, 1)
	assert.Equal(t, "192.168.55.55", alive["localhost-localdomain"].Internal)

	tracker.Ping("localhost-localdomain", ip.InternalPublicPair{
		Internal: "192.168.55.55",
		Public:   "131.131.131.131",
	})

	alive = tracker.ToPossiblyAlive()

	assert.Len(t, alive, 1)
	assert.Equal(t, "192.168.55.55", alive["localhost-localdomain"].Internal)
	assert.Equal(t, "131.131.131.131", alive["localhost-localdomain"].Public)
}

func TestItWillHaveGarbageCollected(t *testing.T) {
	tracker := tracker.New(func(host string, seen time.Time) bool {
		return host != "localhost"
	})

	tracker.Track("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost-localdomain", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost-localdomain", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})
	tracker.Ping("localhost", ip.InternalPublicPair{Internal: "192.168.55.55"})

	assert.Equal(
		t,
		4,
		reflect.Indirect(reflect.ValueOf(*tracker)).FieldByName("events").Len(),
	)
}
