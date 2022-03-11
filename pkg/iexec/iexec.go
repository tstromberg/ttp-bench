package iexec

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func WithTimeout(timeout time.Duration, program string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, program, args...)
	log.Printf("running %s ... (timeout=%s)", cmd, timeout)
	bs, err := cmd.CombinedOutput()
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("hit my %s time limit, have a wonderful day! ...", timeout)
			return nil
		}
		return fmt.Errorf("cmd: %v\n%s", err, bs)
	}

	log.Printf("output: %s", bs)
	return nil
}

func InteractiveTimeout(timeout time.Duration, program string, args ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, program, args...)
	log.Printf("running %s ... (timeout=%s)", cmd, timeout)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Printf("hit my %s time limit, have a wonderful day! ...", timeout)
			return nil
		}
		return err
	}

	return nil
}
