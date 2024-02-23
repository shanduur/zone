package zonego

import "github.com/shanduur/zone/zonego/pkg/exec"

var parser = &exec.Parser{Commander: exec.New()}

func SetParser(p *exec.Parser) {
	if p != nil {
		parser = p
	}
}
