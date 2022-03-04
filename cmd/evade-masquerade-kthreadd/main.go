package main

import (
	"log"
	"os"
	"time"

	"github.com/erikdubbelboer/gspt"
)

func main() {
	gspt.SetProcTitle("[kthreadd]")
	log.Printf("pid %d is hiding as [kthreadd] and sleeping ...", os.Getpid())
	time.Sleep(60 * time.Second)
}
