// New unsigned binary listening from a hidden directory
package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	_ "embed"
)

//go:embed valyrian_debug.zip
var bs []byte

func main() {
	venv, err := os.MkdirTemp(os.TempDir(), "ttp-py-venv")
	if err != nil {
		log.Fatalf("create temp: %v", err)
	}

	if len(bs) == 0 {
		log.Fatalf("embedded 0 byte file")
	}

	c := exec.Command("python", "-m", "venv", venv)
	log.Printf("running %s ...", c)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v\n%s", err, out)
	}

	dest := filepath.Join(venv, "egg.zip")
	log.Printf("writing %d bytes to %s", len(bs), dest)
	if err := os.WriteFile(dest, bs, 0o500); err != nil {
		log.Fatalf("write: %v", err)
	}

	c = exec.Command(filepath.Join(venv, "/bin/pip"), "install", dest)
	log.Printf("running %s ...", c)
	out, err = c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v\n%s", err, out)
	}
	log.Printf("output: %s", out)
}
