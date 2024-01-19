//go:build !windows

// Simulates attack cleanup via bash_history truncation [T1070.003]
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.TruncateShellHistory(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
