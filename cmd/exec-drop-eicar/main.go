// Simulates droppping a known virus signature (EICAR) onto filesystem
package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	// https://en.wikipedia.org/wiki/EICAR_test_file
	e1 := `X5O!P%@AP[4\PZX54(P^)7CC)7}$EI`
	e2 := `CAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*`

	tf, err := ioutil.TempFile(os.TempDir(), "eicar.*.exe")
	if err != nil {
		log.Fatal(err)
	}

	tf.WriteString(e1 + e2)
	defer func() {
		log.Printf("removing %s (did anyone notice?) ...", tf.Name())
		os.Remove(tf.Name())
	}()

	log.Printf("Dropped %s with known anti-virus signature (EICAR)", tf.Name())
	time.Sleep(30 * time.Second)
}
