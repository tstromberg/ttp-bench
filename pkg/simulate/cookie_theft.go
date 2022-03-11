package simulate

import (
	"log"

	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/allbrowsers" // register cookie store finders!
)

func CookieTheft() error {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`google.com`), kooky.Name(`NID`))
	log.Printf("found %d valid NID cookies from google.com!", len(cookies))
	return nil
}
