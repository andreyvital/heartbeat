package ip_test

import (
	"encoding/json"
	"testing"

	"github.com/CentaurWarchief/heartbeat/ip"
	"github.com/stretchr/testify/assert"
)

func TestNewInternalPublicPair(t *testing.T) {
	pair := ip.NewInternalPublicPair("192.168.55.55", "131.131.131.131")

	assert.NotNil(t, pair)
	assert.Equal(t, "192.168.55.55", pair.Internal)
	assert.Equal(t, "131.131.131.131", pair.Public)
}

func TestPublicIPIsNotExposed(t *testing.T) {
	pair := ip.NewInternalPublicPair("192.168.55.55", "131.132.133.134")

	assert.NotNil(t, pair)

	bytes, err := json.Marshal(pair)

	assert.Nil(t, err)
	assert.Equal(t, `{"internal":"192.168.55.55"}`, string(bytes))
}
