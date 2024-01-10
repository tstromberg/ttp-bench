// Simulates C&C discovery via DNS over HTTPS (ala Godlua)
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.DNSOverHTTPS(); err != nil {
		log.Fatalf("dns over https: %v", err)
	}
}
