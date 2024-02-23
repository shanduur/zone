package exec

import (
	"fmt"
	"testing"

	"github.com/shanduur/zone/zonego/pkg/types/zoneadm"
)

func TestParse(t *testing.T) {
	cmd, err := (&Parser{Commander: New()}).Parse(&zoneadm.Zoneadm{
		MarkIncomplete: &zoneadm.MarkIncomplete{},
	})
	t.Errorf("got: %v", func() []string {
		x := []string{}

		for _, a := range cmd.Args() {
			x = append(x, fmt.Sprintf(`"%v"`, a))
		}

		return x
	}())
	t.Errorf("err: %v", err)
}
