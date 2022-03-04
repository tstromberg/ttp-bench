// simulate chrome spawning a shell (not elegantly)
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// spawn at least two suspicious children
func spawnShell() {
	c := exec.Command("bash", "-c", "id")
	log.Printf("running %s from %s", c, os.Args[0])
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	fmt.Printf("%s\n", bs)
	os.Exit(0)
}

func main() {
	dest := "/opt/google/chrome/chrome"

	// I am chrome!
	if strings.Contains(os.Args[0], filepath.Base(dest)) {
		spawnShell()
	}

	log.Printf("backing up %s ...", dest)
	if err := os.Rename(dest, dest+".iocbak"); err != nil {
		log.Fatalf("rename: %v", err)
	}

	go func() {
		log.Printf("restoring %s ...", dest)
		os.Rename(dest+".bak", dest)
	}()

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(os.Args[0], dest); err != nil {
		log.Fatalf("copy: %v", err)
	}

	if err := os.Chmod(dest, 0o755); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	log.Printf("running %s ...", dest)
	c := exec.Command(dest, "--type=renderer", "--ioc", "--display-capture-permissions-policy-allowed", "--origin-trial-disabled-features=ConditionalFocus", "--change-stack-guard-on-fork=enable --lang=en-US --num-raster-threads=4 --enable-main-frame-before-activation --renderer-client-id=7 --launch-time-ticks=103508166127 --shared-files=v8_context_snapshot_data:100 --field-trial-handle=0,17226240353387230227,15124016377124433560,131072")
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}

	log.Printf("output: %s", bs)
}
