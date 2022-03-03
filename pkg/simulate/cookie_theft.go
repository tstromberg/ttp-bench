package simulate

import (
	"github.com/zellyn/kooky"
	_ "github.com/zellyn/kooky/allbrowsers" // register cookie store finders!
	"k8s.io/klog/v2"
)

func CookieTheft() error {
	cookies := kooky.ReadCookies(kooky.Valid, kooky.DomainHasSuffix(`google.com`), kooky.Name(`NID`))

	klog.Infof("found %d cookies from google.com!", len(cookies))
	return nil
}
