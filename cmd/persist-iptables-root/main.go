package main

import (
	"log"
	"os/exec"
	"time"
)

func main() {
	c := exec.Command("iptables", "-I", "INPUT", "-p", "tcp", "--dport", "12345", "-j", "ACCEPT")
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	log.Printf("output: %s", bs)

	time.Sleep(15 * time.Second)

	c = exec.Command("iptables", "-D", "INPUT", "-p", "tcp", "--dport", "12345", "-j", "ACCEPT")
	log.Printf("running %s ...", c)
	bs, err = c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v", err)
	}
	log.Printf("output: %s", bs)

}
