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

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(os.Args[0], dest); err != nil {
		log.Fatalf("copy: %v", err)
	}

	if err := os.Chmod(dest, 0o700); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	log.Printf("running %s ...", dest)
	c := exec.Command("/usr/bin/modload", "ioc")
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}

	log.Printf("output: %s", bs)
}
