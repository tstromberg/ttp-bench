package main

import (
	"log"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
)

func main() {
	if err := simulate.ResolveRandom(); err != nil {
		log.Fatalf("resolve random: %v", err)
	}
}
