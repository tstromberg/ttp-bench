package simulate

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
)

func DNSOverHTTPS() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	c := doh.Use(doh.CloudflareProvider, doh.GoogleProvider)

	for i := 0; i < 16; i++ {
		bytes := make([]byte, 8)
		if _, err := rand.Read(bytes); err != nil {
			return fmt.Errorf("read: %w", err)
		}

		host := fmt.Sprintf("%s.blogspot.com", hex.EncodeToString(bytes))
		log.Printf("looking up TXT record for %s ...", host)

		_, err := c.Query(ctx, dns.Domain(host), dns.TypeTXT)
		if err != nil {
			return fmt.Errorf("query: %w", err)
		}
		time.Sleep(500 * time.Millisecond)
	}

	c.Close()
	return nil
}
