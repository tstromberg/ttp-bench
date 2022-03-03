package main

import (
	"flag"
	"os"
	"os/exec"

	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	checks := []string{
		"cnc-dns-over-https",
		"cnc-resolve-random",
		"creds-browser-cookies",
		"creds-gcp-exfil",
		"evade-bash-history",
	}
	failed := 0
	for i, c := range checks {
		klog.Infof("#%d: testing %s ...", i, c)
		cmd := exec.Command("go", "run", "./cmd/"+c)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			klog.Errorf("#%d: check failed: %v", i, err)
			failed++
		} else {
			klog.Infof("#%d: %s check complete", i, c)
		}
	}

	if failed > 0 {
		klog.Exitf("%d of %d checks failed", failed, len(checks))
		os.Exit(1)
	}
	os.Exit(0)
}
