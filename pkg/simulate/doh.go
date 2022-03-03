package simulate

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/likexian/doh-go"
	"github.com/likexian/doh-go/dns"
	"k8s.io/klog/v2"
)

func DNSOverHTTPS() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	c := doh.Use(doh.CloudflareProvider, doh.GoogleProvider)

	for i := 0; i < 8; i++ {
		bytes := make([]byte, 8)
		if _, err := rand.Read(bytes); err != nil {
			return fmt.Errorf("read: %w", err)
		}

		host := fmt.Sprintf("%s.blogspot.com", hex.EncodeToString(bytes))
		klog.Infof("looking up TXT record for %s ...", host)

		_, err := c.Query(ctx, dns.Domain(host), dns.TypeTXT)
		if err != nil {
			return fmt.Errorf("query: %w", err)
		}
	}

	c.Close()
	return nil
}
