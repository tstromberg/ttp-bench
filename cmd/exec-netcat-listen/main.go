//go:build !windows

// Launches netcat to listen on a port [T1059.004]
package main

import (
	"log"
	"time"

	"github.com/tstromberg/ioc-bench/pkg/iexec"
)

func main() {
	if err := iexec.WithTimeout(30*time.Second, "nc", "-l", "12345"); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}
