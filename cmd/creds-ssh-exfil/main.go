// Simulates theft of GCP credentials [1552.001, T15060.002]
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func SSHCredentialsTheft() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home dir: %w", err)
	}

	path := filepath.Join(home, ".ssh")

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("stat failed: %w", err)
	}

	tf, err := ioutil.TempFile("/tmp", "ssh_ioc.*.tar")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tf.Name())

	log.Printf("archiving %s to %s ...", path, tf.Name())
	cmd := exec.Command("tar", "-cvf", tf.Name(), path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tar failed: %w", err)
	}

	time.Sleep(5 * time.Second)

	log.Printf("cleaning up ...")
	return nil
}

func main() {
	if err := SSHCredentialsTheft(); err != nil {
		log.Fatalf("unexpected error: %v", err)

	}
}
