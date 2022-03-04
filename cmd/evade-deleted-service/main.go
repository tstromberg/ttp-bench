package main

import (
	"log"
	"os"
	"time"
)

func main() {
	a := os.Args
	log.Printf("delete yourself (you have no chance to win): %v", a)
	if err := os.Remove(a[0]); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	log.Printf("waiting around for a minute ...")
	time.Sleep(1 * time.Minute)
}
