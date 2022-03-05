package main

import (
	"log"
	"os/exec"
)

func main() {
	c := exec.Command("nc", "-l", "12345")
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	log.Printf("output: %s", bs)
}
