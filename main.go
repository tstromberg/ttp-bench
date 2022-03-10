package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

var (
	execTimeout  = 70 * time.Second
	buildTimeout = 45 * time.Second
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := context.Background()

	p := createSpinner("Inspecting available simulations ...")
	choices, err := gatherChoices(ctx)
	if err != nil {
		klog.Fatalf("gather choices: %v", err)
	}
	p.Quit()

	selected, err := showChoices(ctx, choices)
	if err != nil {
		klog.Fatalf("show choices: %v", err)
	}

	p = createSpinner(fmt.Sprintf("Building %d simulations ...", len(selected)))
	if err = buildSimulations(ctx, selected); err != nil {
		klog.Exitf("run failed: %v", err)
	}
	p.Quit()

	p = createSpinner(fmt.Sprintf("Executing %d simulations ...", len(selected)))
	// Quit because we announce the simulations differently
	time.Sleep(5 * time.Millisecond)
	p.Quit()

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

func buildSimulations(ctx context.Context, checks []string) error {
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
		cmd := exec.CommandContext(ctx, "go", "build", "../cmd/"+c)
		out, err := cmd.CombinedOutput()
		if err != nil {
			klog.Errorf("#%d: build failed: %v\n%s", i, err, out)
			failed++
			continue
		}
	}

	return nil
}

func runSimulations(ctx context.Context, checks []string) error {
	failed := 0
	su, err := exec.LookPath("doas")
	if err != nil {
		su, err = exec.LookPath("sudo")
		if err != nil {
			su = "su"
		}
	}

	for i, c := range checks {
		if _, err := os.Stat(c); err != nil {
			klog.Errorf("%c not found - skipping")
			failed++
			continue
		}

		title := fmt.Sprintf("Launching %s at %s", c, time.Now().Format(time.RFC3339Nano))
		if strings.HasSuffix(c, "-root") {
			title = title + " (will prompt for root password)"
		}
		announce(title)

		ctx, cancel := context.WithTimeout(context.Background(), execTimeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, "./"+c)
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

		// Make it easier to disambiguate in the logs
		time.Sleep(1 * time.Second)
	}

	if failed > 0 {
		return fmt.Errorf("%d of %d checks failed", failed, len(checks))
	}
	return nil
}
