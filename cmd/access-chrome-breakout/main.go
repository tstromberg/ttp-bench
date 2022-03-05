// simulate chrome spawning a shell (not elegantly)
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
)

// spawn at least two suspicious children
func spawnShell() error {
	c := exec.Command("bash", "-c", "id")
	log.Printf("running %s from %s", c, os.Args[0])
	bs, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run failed: %v\n%s", err, bs)
	}
	fmt.Printf("%s\n", bs)
	time.Sleep(5 * time.Second)
	return nil
}

func setupFakeChrome(dest string) error {
	s, err := os.Stat(os.Args[0])
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	cs, err := os.Stat(dest)
	csize := int64(0)
	if err == nil {
		csize = cs.Size()
	}

	if s.Size() == csize {
		log.Printf("%s already appears to be fake", dest)
	} else {
		log.Printf("backing up %s ...", dest)
		if err := os.Rename(dest, dest+".iocbak"); err != nil {
			return fmt.Errorf("rename: %w", err)
		}
	}

	defer func() {
		log.Printf("restoring %s ...", dest)
		os.Rename(dest+".bak", dest)
	}()

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(os.Args[0], dest); err != nil {
		return fmt.Errorf("copy: %v", err)
	}

	if err := os.Chmod(dest, 0o755); err != nil {
		return fmt.Errorf("chmod failed: %v", err)
	}

	user := os.Getenv("DOAS_USER")
	if user == "" {
		user = os.Getenv("SUDO_USER")
	}
	if user == "" {
		user = "nobody"
	}

	fakeChrome := dest + " --type=renderer --ioc --display-capture-permissions-policy-allowed --origin-trial-disabled-features=ConditionalFocus --change-stack-guard-on-fork=enable --lang=en-US --num-raster-threads=4 --enable-main-frame-before-activation --renderer-client-id=7 --launch-time-ticks=103508166127 --shared-files=v8_context_snapshot_data:100"
	c := exec.Command("/usr/bin/su", user, "-c", fakeChrome)
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run failed: %v\n%s", err, bs)
	}
	log.Printf("output: %s", bs)
	return nil
}

func main() {
	dest := "/opt/google/chrome/chrome"

	// I am chrome!
	if filepath.Base(os.Args[0]) == filepath.Base(dest) {
		if err := spawnShell(); err != nil {
			log.Fatalf("spawn: %v", err)
		}
		os.Exit(0)
	}
	if err := setupFakeChrome(dest); err != nil {
		log.Fatalf("setup fake chrome: %v", err)
	}
}
