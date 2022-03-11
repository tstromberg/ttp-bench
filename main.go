package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

var (
	execTimeout  = 70 * time.Second
	buildTimeout = 45 * time.Second
	timeFormat   = "2006-01-02 15:04:05.999"
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()
	status(fmt.Sprintf("Gathering simulations for %s/%s", runtime.GOOS, runtime.GOARCH))

	choices, err := gatherChoices(ctx)
	if err != nil {
		klog.Fatalf("gather choices: %v", err)
	}

	selected, err := selectChoices(ctx, choices)
	if err != nil {
		klog.Fatalf("show choices: %v", err)
	}

	if len(selected) == 0 {
		msg("また会おうね")
		os.Exit(0)
	}

	status(fmt.Sprintf("Building %d selected simulations", len(selected)))
	if err = buildSimulations(ctx, selected); err != nil {
		klog.Exitf("run failed: %v", err)
	}

	status(fmt.Sprintf("Executing %d selected simulations", len(selected)))
	if err = runSimulations(ctx, selected); err != nil {
		klog.Exitf("run failed: %v", err)
	}
}

type choice struct {
	name string
	desc string
}

func gatherChoices(ctx context.Context) ([]choice, error) {
	dirs, err := os.ReadDir("cmd")
	if err != nil {
		klog.Exitf("readdir failed: %v", err)
	}

	choices := []choice{}

	for _, d := range dirs {
		c := d.Name()
		cmd := exec.CommandContext(ctx, "go", "doc", "./cmd/"+c)
		out, err := cmd.CombinedOutput()
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() == 1 {
				klog.V(2).Infof("%s: %s", c, out)
				continue
			}
			return choices, fmt.Errorf("%s failed: %v\n%s", cmd, err, out)
		}

		choices = append(choices, choice{name: c, desc: strings.TrimSpace(string(out))})
	}

	return choices, nil
}

func buildSimulations(ctx context.Context, checks []choice) error {
	failed := 0

	if err := os.MkdirAll("out", 0o700); err != nil {
		return fmt.Errorf("mkdir: %w", err)
	}

	if err := os.Chdir("out"); err != nil {
		return fmt.Errorf("chdir: %w", err)
	}

	for i, c := range checks {
		ctx, cancel := context.WithTimeout(context.Background(), buildTimeout)
		defer cancel()
		cmd := exec.CommandContext(ctx, "go", "build", "../cmd/"+c.name)
		out, err := cmd.CombinedOutput()
		if err != nil {
			klog.Errorf("#%d: build failed: %v\n%s", i, err, out)
			failed++
			continue
		}
	}

	return nil
}

func runSimulations(ctx context.Context, checks []choice) error {
	failed := 0
	su, err := exec.LookPath("doas")
	if err != nil {
		su, err = exec.LookPath("sudo")
		if err != nil {
			su = "su"
		}
	}

	for i, c := range checks {
		if _, err := os.Stat(c.name); err != nil {
			klog.Errorf("%c not found - skipping")
			failed++
			continue
		}

		title := fmt.Sprintf("[%d of %d] %s at %s", i+1, len(checks), c, time.Now().Format(timeFormat))

		announce(title)
		subtitle(c.desc)

		ctx, cancel := context.WithTimeout(context.Background(), execTimeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, "./"+c.name)
		if strings.HasSuffix(c.name, "-root") {
			notice(fmt.Sprintf("This simulation requires root privileges - will use %s", su))
			cmd = exec.CommandContext(ctx, su, "./"+c.name)
		}

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			klog.Errorf("#%d: check failed: %v", i, err)
			failed++
			continue
		}

		// Make it easier to disambiguate in the logs
		time.Sleep(1 * time.Second)
	}

	if failed > 0 {
		return fmt.Errorf("%d of %d checks failed", failed, len(checks))
	}
	return nil
}
