//go:build !illumos
// +build !illumos

package exec

import "fmt"

type cmd struct {
	args []string
}

func (c *cmd) Args() []string {
	return c.args
}

func (c *cmd) Run() error {
	fmt.Printf("running: %v\n", c.args)
	return nil
}

func (e *Executor) Command(n string, args ...string) Runner {
	return &cmd{
		args: append([]string{n}, args...),
	}
}
