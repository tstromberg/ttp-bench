// New unsigned binary listening from a hidden directory
package main

import (
	"log"
	"net"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

var listenPort = ":39999"

func main() {
	log.Printf("args: %s", os.Args)
	if len(os.Args) > 1 {
		log.Printf("listening from %s at %s", os.Args[0], listenPort)
		l, err := net.Listen("tcp", listenPort)
		if err != nil {
			log.Panicf("listen failed: %v", err)
		}
		defer l.Close()
		l.Accept()
		os.Exit(0)
	}

	cfg, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("user config dir: %v", err)
	}

	root := filepath.Join(cfg, "ttp-bench")
	if err := os.MkdirAll(root, 0o777); err != nil {
		log.Fatalf("mkdir: %v", err)
	}

	tf, err := os.CreateTemp(root, ".XXXX")
	if err != nil {
		log.Fatalf("create temp: %v", err)
	}

	defer os.Remove(tf.Name())
	src := os.Args[0]
	dest := tf.Name()

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(src, dest); err != nil {
		log.Fatalf("copy: %v", err)
	}

	if err := os.Chmod(dest, 0o755); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	tf.Close()

	iexec.InteractiveTimeout(70*time.Second, dest, "--not-a-hacker-i-promise")
}
