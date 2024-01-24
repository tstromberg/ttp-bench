//go:build !windows

// Launches a temporary reverse shell using Python
package main

import (
	"log"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	if err := simulate.PythonReverseShell(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
