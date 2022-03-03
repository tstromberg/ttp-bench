package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.GCloudCredentialsTheft(); err != nil {
		log.Fatalf("unexpected error: %v", err)

	}
}
