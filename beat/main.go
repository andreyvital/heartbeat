package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/CentaurWarchief/heartbeat/ip"
)

func beat(host string) {
	body, err := json.Marshal(map[string]string{
		"host":     host,
		"internal": ip.Internal(),
		"public":   ip.Public(),
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("POST", os.Args[1], bytes.NewReader(body))

	if err != nil {
		return
	}

	cli := &http.Client{}

	if _, err := cli.Do(req); err != nil {
		fmt.Println(err)
	}
}

func main() {
	host, err := os.Hostname()

	if err != nil {
		fmt.Println(err)
		return
	}

	beat(host)

	wg := sync.WaitGroup{}
	wg.Add(1)

	stop := make(chan bool, 1)
	tick := time.NewTicker(30 * time.Second)

	go func() {
		for {
			select {
			case <-tick.C:
				beat(host)
			case <-stop:
				return
			}
		}
	}()

	defer func() {
		stop <- true
	}()

	wg.Wait()
}
