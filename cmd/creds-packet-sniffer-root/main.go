// Simulates theft of credentials via network sniffing [T1040]
package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.PacketSniffer(); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
}
