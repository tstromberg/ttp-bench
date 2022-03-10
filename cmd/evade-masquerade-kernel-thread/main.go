// Simulates process masquerading as a kernel thread [T1036.004]
package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/erikdubbelboer/gspt"
	"k8s.io/klog/v2"
)

func main() {
	target := "kernel_task"

	switch runtime.GOOS {
	case "linux":
		target = "[kthreadd]"
	case "windows":
		target = "rundll32.exe"
	}

	klog.Infof("%s -> %s", runtime.GOOS, target)

	gspt.SetProcTitle(target)
	log.Printf("pid %d is hiding as %q and sleeping ...", os.Getpid(), target)
	time.Sleep(60 * time.Second)
}
