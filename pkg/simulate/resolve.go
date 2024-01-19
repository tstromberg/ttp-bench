package simulate

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"time"
)

func ResolveRandom() error {
	for i := 0; i < 32; i++ {
		bytes := make([]byte, 8)
		if _, err := rand.Read(bytes); err != nil {
			return fmt.Errorf("read: %w", err)
		}

		host := fmt.Sprintf("%s.dns.%d.eu.org", hex.EncodeToString(bytes), i)
		log.Printf("looking up %s ...", host)
		time.Sleep(time.Millisecond * 50)
		_, err := net.LookupHost(host)

		if err != nil {
			if de, ok := err.(*net.DNSError); ok {
				if de.IsNotFound {
					continue
				}
			}

			return fmt.Errorf("lookup %s: %w", host, err)
		}
	}
	return nil
}
