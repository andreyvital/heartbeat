package net

type Interface struct {
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Loopback bool   `json:"loopback"`
}
