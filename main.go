package main

import (
	"context"
	"flag"
	"os"
	"os/exec"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

var (
	execTimeout = 65 * time.Second
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	failed := 0
	checks := []string{}
	dirs, err := os.ReadDir("cmd")
	if err != nil {
		klog.Exitf("readdir failed: %v", err)
	}

	for _, d := range dirs {
		checks = append(checks, d.Name())
	}

	su, err := exec.LookPath("doas")
	if err != nil {
		su, err = exec.LookPath("sudo")
		if err != nil {
			su = "su"
		}
	}

	klog.Infof("found %d checks: %v", len(checks), checks)
	klog.Infof("giving each check %s to execute", execTimeout)

	if err := os.MkdirAll("out", 0o700); err != nil {
		klog.Exitf("mkdir out: %v", err)
	}

	if err := os.Chdir("out"); err != nil {
		klog.Exitf("chdir out: %v", err)
	}

	for i, c := range checks {
		ctx, cancel := context.WithTimeout(context.Background(), execTimeout)
		defer cancel()

		klog.Infof("#%d: testing %s ...", i, c)
		cmd := exec.CommandContext(ctx, "go", "build", "../cmd/"+c)
		out, err := cmd.CombinedOutput()
		if err != nil {
			klog.Errorf("#%d: build failed: %v\n%s", i, err, out)
			failed++
			continue
		}
		klog.Infof("#%d: %s build complete", i, c)

		cmd = exec.CommandContext(ctx, "./"+c)
		if strings.HasSuffix(c, "-root") {
			klog.Infof("root required for %s - will prompt", c)
			cmd = exec.CommandContext(ctx, su, "./"+c)
		}

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			klog.Errorf("#%d: check failed: %v", i, err)
			failed++
			continue
		}
		klog.Infof("#%d: %s check complete", i, c)
	}

	if failed > 0 {
		klog.Exitf("%d of %d checks failed", failed, len(checks))
		os.Exit(1)
	}
	os.Exit(0)
}
