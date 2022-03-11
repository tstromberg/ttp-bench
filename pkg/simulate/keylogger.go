package simulate

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/MarinX/keylogger"
)

func listenKeyboard(kbd string) error {
	if os.Geteuid() != 0 {
		log.Printf("effective uid is %d, not 0 (sniffing may not work)", os.Geteuid())
	}

	k, err := keylogger.New(kbd)
	if err != nil {
		return fmt.Errorf("keyboard: %w", err)
	}
	defer k.Close()

	events := k.Read()
	for e := range events {
		if e.KeyPress() {
			log.Printf("sniffed key press on %s (hiding the value for privacy)", kbd)
		}
	}

	return nil
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}

func Keylogger() error {
	var wg sync.WaitGroup
	timeout := 10 * time.Second

	for _, dev := range keylogger.FindAllKeyboardDevices() {
		log.Printf("listening for keystrokes on %s (timeout=%s) ...", dev, timeout)
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			listenKeyboard(d)
			log.Printf("%s done", d)
		}(dev)
	}

	st := waitTimeout(&wg, timeout)
	log.Printf("our job here is done, timeout=%v", st)
	return nil
}
