// Simulates theft of web session cookies [T1539]
package main

import (
	"log"

	"github.com/browserutils/kooky"
	_ "github.com/browserutils/kooky/browser/all"
)

func main() {
	for _, st := range kooky.FindAllCookieStores() {
		log.Printf("found cookie store: %s -> %s", st.Browser(), st.FilePath())
	}

	for _, c := range kooky.ReadCookies(kooky.DomainHasSuffix(`google.com`), kooky.Name(`NID`)) {
		log.Printf("found google.com NID cookie expiring at %s", c.Expires)
	}
}
