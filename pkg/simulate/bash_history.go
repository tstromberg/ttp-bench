package simulate

import (
	"fmt"
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
	klog.Infof("backing up %s ...", path)
	if err := cp.Copy(path, path+".bak"); err != nil {
		return fmt.Errorf("copy: %w", err)
	}

	time.Sleep(1 * time.Second)
	klog.Infof("Truncating %s ...", path)

	if err := os.Truncate(path, 0); err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	s, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	klog.Infof("stat: %+v", s)
	time.Sleep(1 * time.Second)

	klog.Infof("restoring %s ...", path)
	return cp.Copy(path+".bak", path)
}
