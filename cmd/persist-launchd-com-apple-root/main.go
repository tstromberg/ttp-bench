//go:build darwin

// Simulates persistance via a fake unsigned Apple launchd service
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	cp "github.com/otiai10/copy"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

var plist = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>__SVC__</string>
	<key>Program</key>
	<string>__PATH__</string>
	<key>RunAtLoad</key>
	<true/>
	<key>KeepAlive</key>
	<true/>
</dict>
</plist>`

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func main() {
	basedir := "/Library/WebServer/.TTPBench"
	install := false
	if !exists(basedir) {
		if err := os.MkdirAll(basedir, 0o755); err != nil {
			log.Fatalf("mkdir: %v", err)
		}
		install = true
	}

	dest := filepath.Join(basedir, "Safari")
	if !exists(dest) {
		log.Printf("populating %s ...", dest)
		if err := cp.Copy(os.Args[0], dest); err != nil {
			log.Fatalf("copy: %v", err)
		}
		install = true
	}

	svc := "com.apple.ttp-bench"
	path := fmt.Sprintf("/Library/LaunchAgents/%s.plist", svc)
	if !exists(path) {
		contents := []byte(strings.ReplaceAll(strings.ReplaceAll(plist, "__SVC__", svc), "__PATH__", dest))
		log.Printf("writing %s to %s ...", contents, path)
		err := os.WriteFile(path, contents, 0o555)
		if err != nil {
			log.Fatalf("writefile: %v", err)
		}
		install = true
	}

	if install {
		if err := iexec.WithTimeout(10*time.Second, "/bin/launchctl", "enable", fmt.Sprintf("system/%s", svc)); err != nil {
			log.Fatalf("bootstrap: %v", err)
		}

		if err := iexec.WithTimeout(10*time.Second, "/bin/launchctl", "bootstrap", "system", path); err != nil {
			log.Fatalf("bootstrap: %v", err)
		}
		defer func() {
			if err := iexec.WithTimeout(10*time.Second, "/bin/launchctl", "bootout", "system", path); err != nil {
				log.Printf("launchd stop: %v", err)
			}
			os.Remove(path)
			os.Remove(dest)
			os.Remove(basedir)
		}()
	}

	wait := 60 * time.Second
	log.Printf("resting for %s ...", wait)
	time.Sleep(wait)
}
