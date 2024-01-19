package simulate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

func Traitor(args ...string) error {
	bin := fmt.Sprintf("traitor-%s", runtime.GOARCH)
	url := fmt.Sprintf("https://github.com/liamg/traitor/releases/download/v0.0.14/%s", bin)
	td := os.TempDir()
	os.Chdir(td)

	log.Printf("Downloading %s to %s ...", url, td)
	if _, err := grab.Get(".", url); err != nil {
		return fmt.Errorf("grab: %w", err)
	}

	if err := os.Chmod(bin, 0o777); err != nil {
		return fmt.Errorf("chmod failed: %v", err)
	}

	c := exec.Command("strip", bin)
	if err := c.Run(); err != nil {
		log.Printf("strip failed: %v", err)
	}

	defer func() {
		log.Printf("removing %s", bin)
		os.Remove(bin)
	}()

	return iexec.InteractiveTimeout(75*time.Second, "./"+bin, args...)
}
