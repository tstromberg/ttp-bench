package simulate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func GCloudCredentialsTheft() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home dir: %w", err)
	}

	path := filepath.Join(home, ".config/gcloud")

	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("stat failed: %w", err)
	}

	tf, err := ioutil.TempFile("/tmp", "ioc.*.tar")
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

	time.Sleep(1 * time.Second)

	log.Printf("cleaning up ...")
	return nil
}
