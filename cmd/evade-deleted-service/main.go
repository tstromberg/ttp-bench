//go:build !windows

// Simulates a service running by a binary which no longer exists
package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	a := os.Args
	log.Printf("delete yourself (you have no chance to win): %v", a)
	if err := os.Remove(a[0]); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	url := "https://suspicious-ioc.blogspot.com/"
	log.Printf("making connection to %s", url)
	_, err := http.Get(url)
	if err != nil {
		log.Printf("%s returned error: %v (don't care)", url, err)
	}

	timeout := 45 * time.Second
	log.Printf("waiting around for %s ...", timeout)
	time.Sleep(timeout)
}
