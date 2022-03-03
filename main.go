package main

import (
	"flag"
	"os"

	"github.com/tstromberg/ioc-bench/pkg/simulate"
	"k8s.io/klog/v2"
)

type errFunc func() error

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	checks := []errFunc{
		simulate.GCloudCredentialsTheft,
		simulate.CookieTheft,
		simulate.TruncateBashHistory,
		simulate.DNSOverHTTPS,
		simulate.ResolveRandom,
	}

	failed := 0
	for i, c := range checks {
		if err := c(); err != nil {
			klog.Errorf("#%d: check failed: %v", i, err)
			failed++
		} else {
			klog.Infof("#%d: check complete", i)
		}
	}

	if failed > 0 {
		klog.Exitf("%d of %d checks failed", failed, len(checks))
		os.Exit(1)
	}
	os.Exit(0)
}
