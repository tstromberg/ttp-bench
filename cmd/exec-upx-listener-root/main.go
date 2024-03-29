// New unsigned obfuscated binary listening from a hidden directory as root
package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"time"

	cp "github.com/otiai10/copy"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

var listenPort = ":19"

func main() {
	log.Printf("args: %s", os.Args)
	if len(os.Args) > 1 {

		// make outgoing connection
		endpoint := "http://dprkportal.kp/239582395810"
		log.Printf("fetching %s", endpoint)
		http.Get(endpoint)

		log.Printf("listening from %s at %s", os.Args[0], listenPort)
		l, err := net.Listen("tcp", listenPort)
		if err != nil {
			log.Panicf("listen failed: %v", err)
		}
		defer l.Close()
		l.Accept()
		os.Exit(0)
	}

	upx, err := exec.LookPath("upx")
	if err != nil {
		log.Fatalf("upx not found: %v", err)
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

	if err := os.Chmod(dest, 0o777); err != nil {
		log.Fatalf("chmod failed: %v", err)
	}

	tf.Close()

	c := exec.Command(upx, "-f", dest)
	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("run failed: %v\n%s", err, bs)
	}

	// macOS will kill this process: https://github.com/upx/upx/issues/424
	//
	// ASP: Security policy would not allow process: 55547, /private/var/tmp/.XXXX3857572708
	//
	// We've kept this check on macOS because it's still a test to see if
	// a SIEM or EDR is relaying these security policy interceptions.
	err = iexec.InteractiveTimeout(62*time.Second, dest, "--omg", "--wtf", "--bbq")
	if err != nil {
		log.Fatalf("err: %v (possibly killed by SIP?)", err)
	}
}
