package simulate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
	"k8s.io/klog/v2"
)

func TruncateBashHistory() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("home dir: %w", err)
	}

	path := filepath.Join(home, ".bash_history")
	log.Printf("backing up %s ...", path)
	if err := cp.Copy(path, path+".bak"); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	defer func() {
		log.Printf("restoring %s ...", path)
		if err := cp.Copy(path+".bak", path); err != nil {
			klog.Errorf("unable to restore %s: %v", path, err)
		}
	}()

	time.Sleep(1 * time.Second)
	log.Printf("Truncating %s ...", path)

	if err := os.Truncate(path, 0); err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	s, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	log.Printf("stat: %+v", s)
	time.Sleep(1 * time.Second)
	return nil
}
