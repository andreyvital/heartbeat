package tracker

import (
	"encoding/json"
	"log"
	"net/http"
)

func PossiblyAlive(t *Tracker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(t.ToPossiblyAlive()); err != nil {
			log.Println(err)
		}
	}
}
