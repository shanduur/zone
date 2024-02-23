package zonego

import (
	"github.com/shanduur/zone/zonego/pkg/types/zlogin"
)

func Zlogin(zone string, utility string, args ...string) error {
	z := &zlogin.Zlogin{
		Zonename:      zone,
		Utility:       utility,
		Arguments:     args,
		RedirectStdin: true,
	}

	return parser.MustParse(z).Run()
}
