//go:build +darwin

// Simulates a command inserting itself into the system launchd as a fake unsigned Apple service
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

plist := ```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>io.kandji.apple.</string>
	<key>Program</key>
	<string>__PATH__</string>
	<key>RunAtLoad</key>
	<true/>
	<key>KeepAlive</key>
	<true/>
</dict>
</plist>
```

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func main() {
	_, err := exec.LookPath("launchctl")
	if err != nil {
		log.Printf("crontab command not found, skipping")
		os.Exit(0)
	}
	
	if !exists(dest) { 
		dest := "/Library/WebServer/Documents/.share"
		log.Printf("populating %s ...", dest)
		if err := cp.Copy(os.Args[0], dest); err != nil {
			log.Fatalf("copy: %v", err)
		}
	}

	path := "/Library/LaunchAgents/com.apple.ioc-bench.plist"
	if !exists(path) {
		log.Printf("writing to %s ...", path)
		err := os.WriteFile(plist, []byte(strings.ReplaceAll(plist, "__PATH__", dest))
		if err != nil {
			log.Fatalf("writefile: %v")
		}
	}

	if err := iexec.WithTimeout(10*time.Second, "launchtl", "enable", path); err != nil {
		log.Fatalf("enable: %v", err)
	}


	if err := iexec.WithTimeout(10*time.Second, "launchtl", "start", path); err != nil {
		log.Fatalf("enable: %v", err)
	}

	defer func() {
		if err := iexec.WithTimeout(10*time.Second, "launchtl", "stop", path); err != nil {
			log.Errorf("stop: %v", err)
		}
		
		if err := iexec.WithTimeout(10*time.Second, "launchtl", "disable", path); err != nil {
			log.Errorf("disable: %v", err)
		}
	
		os.Remove(path)
		os.Remove(dest)
	}()

	wait := 60 * time.Second
	log.Printf("resting for %s ...", wait)
	time.Sleep(wait)
}
