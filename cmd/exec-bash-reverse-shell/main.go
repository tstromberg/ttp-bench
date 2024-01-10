//go:build !windows

// Launches a temporary reverse shell using bash
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.BashReverseShell(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
