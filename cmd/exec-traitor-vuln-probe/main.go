//go:build linux

// Simulates probing system for privilege escalation vulns
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.Traitor(); err != nil {
		log.Fatalf("exploit failed: %v", err)
	}
}
