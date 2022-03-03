package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.DNSOverHTTPS(); err != nil {
		log.Fatalf("dns over https: %v", err)
	}
}
