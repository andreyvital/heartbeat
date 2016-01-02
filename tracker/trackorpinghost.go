package tracker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CentaurWarchief/heartbeat/ip"
)

type trackOrPingRequestPayload struct {
	Host     string `json:"host"`
	Internal string `json:"internal"`
	Public   string `json:"public"`
}

func TrackOrPingHost(t *Tracker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload trackOrPingRequestPayload

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer r.Body.Close()

		if err := json.Unmarshal(body, &payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pair := ip.NewInternalPublicPair(payload.Internal, payload.Public)

		w.WriteHeader(http.StatusOK)

		if t.IsHostBeingTracked(payload.Host) {
			t.Ping(payload.Host, pair)
			return
		}

		t.Track(payload.Host, pair)
	}
}
