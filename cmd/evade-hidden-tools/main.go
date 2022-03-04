package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	dir := "/var/tmp/.hidden"
	log.Printf("creating %s ...", dir)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		log.Fatalf("mkdir failed: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		log.Fatalf("chdir failed: %v", err)
	}

	url := "https://github.com/kubernetes/minikube/releases/download/v1.25.2/minikube-" + runtime.GOOS + "-" + runtime.GOARCH

	log.Printf("downloading %s to %s", url, dir)
	c := exec.Command("curl", "-L", "-o", "xxx", url)
	if err := c.Run(); err != nil {
		log.Fatalf("run failed: %v", err)
	}

	if err := os.Chmod("./xxx", 0o700); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	log.Printf("running %s/xxx ...", dir)
	c = exec.Command("./xxx", "version")
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	log.Printf("output: %s", out)
}
