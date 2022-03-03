package simulate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"k8s.io/klog/v2"
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

	klog.Infof("archiving %s to %s ...", path, tf.Name())
	cmd := exec.Command("tar", "-cvf", tf.Name(), path)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tar failed: %w", err)
	}

	time.Sleep(1 * time.Second)

	klog.Infof("cleaning up ...")
	return nil
}
