package simulate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	cp "github.com/otiai10/copy"
)

// SpawnShell will spawn at least two suspicious children
func SpawnShellID() error {
	c := exec.Command("bash", "-c", "id")
	log.Printf("running %s from %s", c, os.Args[0])
	bs, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run failed: %v\n%s", err, bs)
	}
	fmt.Printf("%s\n", bs)
	time.Sleep(5 * time.Second)
	return nil
}

func ReplaceAndLaunch(src string, dest string, args string) error {
	s, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}

	cs, err := os.Stat(dest)
	csize := int64(0)
	if err == nil {
		csize = cs.Size()
	}

	if s.Size() == csize {
		log.Printf("%s already appears to be fake", dest)
	} else {
		log.Printf("backing up %s ...", dest)
		if err := os.Rename(dest, dest+".iocbak"); err != nil {
			return fmt.Errorf("rename: %w", err)
		}
	}

	defer func() {
		log.Printf("restoring %s ...", dest)
		os.Rename(dest+".bak", dest)
	}()

	log.Printf("populating %s ...", dest)
	if err := cp.Copy(src, dest); err != nil {
		return fmt.Errorf("copy: %v", err)
	}

	if err := os.Chmod(dest, 0o755); err != nil {
		return fmt.Errorf("chmod failed: %v", err)
	}

	c := exec.Command("sh", "-c", dest, args)

	// If we are root, swap to the user who ran ioc-bench
	if syscall.Geteuid() == 0 {
		user := os.Getenv("DOAS_USER")
		if user == "" {
			user = os.Getenv("SUDO_USER")
		}
		if user == "" {
			user = "nobody"
		}
		c = exec.Command("/usr/bin/su", user, "-c", fmt.Sprintf(`"%s" %s`, dest, args))
	}

	log.Printf("running %s ...", c)
	bs, err := c.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run failed: %v\n%s", err, bs)
	}
	log.Printf("output: %s", bs)
	return nil
}
