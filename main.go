package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CentaurWarchief/heartbeat/tracker"
	"github.com/gorilla/mux"
)

func main() {
	t := tracker.New(func(host string, seen time.Time) bool {
		return time.Now().Sub(seen) < 30*time.Second
	})

	r := mux.NewRouter().StrictSlash(true)

	r.Methods("GET").Path("/").HandlerFunc(tracker.PossiblyAlive(t))
	r.Methods("POST").Path("/").HandlerFunc(tracker.TrackOrPingHost(t))

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	if err := http.ListenAndServe(":1025", r); err != nil {
		fmt.Println(err.Error())
	}
}
