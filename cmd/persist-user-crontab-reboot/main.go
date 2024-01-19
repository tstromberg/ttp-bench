//go:build !windows

// Simulates a command inserting itself into the user crontab for persistence
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

func main() {
	_, err := exec.LookPath("crontab")
	if err != nil {
		log.Printf("crontab command not found, skipping")
		os.Exit(0)
	}

	cd, err := os.UserConfigDir()
	if err != nil {
		log.Panicf("unable to find config dir: %v", err)
	}

	dest := filepath.Join(cd, "ioc-persist")
	log.Printf("populating %s ...", dest)
	if err := cp.Copy(os.Args[0], dest); err != nil {
		log.Fatalf("copy: %v", err)
	}

	install := fmt.Sprintf(`( crontab -l | egrep -v "%s"; echo @reboot "%s" ) | crontab`, dest, dest)
	if err := iexec.WithTimeout(10*time.Second, "sh", "-c", install); err != nil {
		log.Fatalf("crontab install failed: %v", err)
	}

	defer func() {
		remove := fmt.Sprintf(`crontab -l | egrep -v "%s" | crontab`, dest)
		if err := iexec.WithTimeout(10*time.Second, "sh", "-c", remove); err != nil {
			log.Fatalf("crontab remove failed: %v", err)
		}
		os.Remove(dest)
	}()

	wait := 60 * time.Second
	log.Printf("resting for %s ...", wait)
	time.Sleep(wait)
}
