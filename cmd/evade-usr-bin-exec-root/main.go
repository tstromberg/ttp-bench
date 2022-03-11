//go:build linux

// Simulates malicious program installing itself into /usr/bin [T1036.005]
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	cp "github.com/otiai10/copy"
)

func main() {
	dest := "/usr/bin/modload"

	if strings.Contains(os.Args[0], filepath.Base(dest)) {
		log.Printf("running netstat from %s", os.Args[0])
		c := exec.Command("netstat", "-an")
		bs, err := c.CombinedOutput()
		if err != nil {
			log.Fatalf("run failed: %v", err)
		}
		fmt.Printf("%s\n", bs)
		os.Exit(0)
	}

	ms, err := os.Stat(os.Args[0])
	if err != nil {
		log.Fatalf("unable to stat myself: %v", err)
	}

	ds, err := os.Stat(dest)
	if err == nil && ds.Size() != ms.Size() {
		log.Fatalf("found unexpected file in %s", dest)
	}

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(os.Args[0], dest); err != nil {
		log.Fatalf("copy: %v", err)
	}

	defer func() {
		log.Printf("removing implant from %s ...", dest)
		os.Remove(dest)
	}()

	if err := os.Chmod(dest, 0o700); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	log.Printf("running %s ...", dest)
	c := exec.Command(dest, "ioc")
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}

	log.Printf("output: %s", bs)
}
