// Simulates an overflow where Google Chrome spawns a shell [T1189]
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/tstromberg/ttp-bench/pkg/simulate"
)

func main() {
	globs := []string{
		"/opt/google/chrome*/chrome",
		"/Applications/Google Chrome*.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/*/Helpers/Google Chrome Helper (Renderer).app/Contents/MacOS/Google Chrome Helper (Renderer)",
	}

	args := "--type=renderer --ioc --display-capture-permissions-policy-allowed --change-stack-guard-on-fork=enable --lang=en-US --num-raster-threads=4 --enable-main-frame-before-activation --renderer-client-id=7 --launch-time-ticks=103508166127 --shared-files=v8_context_snapshot_data:100"

	dest := ""
	for _, g := range globs {
		paths, err := filepath.Glob(g)
		if err != nil {
			log.Printf("glob error: %v", err)
			continue
		}

		for _, p := range paths {
			log.Printf("found %s", p)
			dest = p
			break
		}
	}

	if dest == "" {
		log.Fatalf("unable to find a chrome browser to emulate")
	}

	// I am chrome!
	if filepath.Base(os.Args[0]) == filepath.Base(dest) {
		if err := simulate.SpawnShellID(); err != nil {
			log.Fatalf("spawn: %v", err)
		}
		os.Exit(0)
	}
	if err := simulate.ReplaceAndLaunch(os.Args[0], dest, args); err != nil {
		log.Fatalf("replace and launch: %v", err)
	}
}
