// Simulates process masquerading as a kernel thread [T1036.004]
package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/erikdubbelboer/gspt"
)

func main() {
	target := "cloudphotod"

	switch runtime.GOOS {
	case "linux":
		target = "(sd-pam)"
	case "windows":
		target = "rundll32.exe"
	}

	log.Printf("%s -> %s", runtime.GOOS, target)

	gspt.SetProcTitle(target)
	log.Printf("pid %d is hiding as %q and sleeping ...", os.Getpid(), target)
	time.Sleep(60 * time.Second)
}
