// Downloads and launches LinPEAS
package main

import (
	"log"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

func main() {
	url := "https://github.com/carlospolop/PEASS-ng/releases/latest/download/linpeas.sh"
	td := os.TempDir()
	bin := "linpeas.sh"
	os.Chdir(td)

	log.Printf("Downloading %s to %s ...", url, td)
	if _, err := grab.Get(".", url); err != nil {
		log.Fatalf("grab: %v", err)
	}

	if err := os.Chmod(bin, 0o707); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	defer func() {
		log.Printf("removing %s", bin)
		os.Remove(bin)
	}()

	args := []string{"-s"}
	iexec.InteractiveTimeout(90*time.Second, "./"+bin, args...)
}
