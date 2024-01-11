// New unsigned binary listening from a hidden directory
package main

import (
	"log"
	"net"
	"os"
	"os/exec"

	cp "github.com/otiai10/copy"
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

	tf, err := os.CreateTemp("/var/tmp", ".XXXX")
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

	c := exec.Command(dest, "ttp-bench")
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v\n%s", err, bs)
	}
	log.Printf("output: %s", bs)

}
