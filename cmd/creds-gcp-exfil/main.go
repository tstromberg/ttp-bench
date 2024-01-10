// Simulates theft of GCP credentials [1552.001, T15060.002]
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.GCloudCredentialsTheft(); err != nil {
		log.Fatalf("unexpected error: %v", err)

	}
}
