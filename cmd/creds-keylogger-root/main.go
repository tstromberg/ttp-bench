//go:build linux

// Simulate theft of credentials via key logging [T1056]
package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.Keylogger(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
