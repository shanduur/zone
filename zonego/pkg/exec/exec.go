package exec

import (
	"io"
	"os"
)

type Commander interface {
	Command(string, ...string) Runner
}

type Runner interface {
	Run() error
	Args() []string
}

func New(opts ...Option) *Executor {
	ex := &Executor{
		stdout: os.Stdout,
		stderr: os.Stderr,
		stdin:  os.Stdin,
	}

	for _, opt := range opts {
		opt.apply(ex)
	}

	return ex
}

type Executor struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

type Option interface {
	apply(*Executor)
}

type option struct {
	applyFunc func(*Executor)
}

func (o *option) apply(e *Executor) {
	o.applyFunc(e)
}

func WithStdin(r io.Reader) Option {
	return &option{applyFunc: func(e *Executor) {
		e.stdin = r
	}}
}

func WithStdout(w io.Writer) Option {
	return &option{applyFunc: func(e *Executor) {
		e.stdout = w
	}}
}

func WithStderr(w io.Writer) Option {
	return &option{applyFunc: func(e *Executor) {
		e.stderr = w
	}}
}
