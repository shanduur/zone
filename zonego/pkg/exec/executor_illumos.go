//go:build illumos
// +build illumos

package exec

import "os/exec"

type cmd struct {
	*exec.Cmd
}

func (c *cmd) Args() []string {
	return c.Cmd.Args
}

func (e *Executor) Command(n string, args ...string) Runner {
	cmd := exec.Command(n, args...)

	return &cmd{Cmd: cmd}
}
