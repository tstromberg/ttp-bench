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
	target := "[kthreadd]"
	if runtime.GOOS == "darwin" {
		target = "kernel_task"
	}

	klog.Infof("%s -> %s", runtime.GOOS, target)

	gspt.SetProcTitle(target)
	log.Printf("pid %d is hiding as %q and sleeping ...", os.Getpid(), target)
	time.Sleep(60 * time.Second)
}
