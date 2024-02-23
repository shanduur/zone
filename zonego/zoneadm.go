package zonego

import "github.com/shanduur/zone/zonego/pkg/types/zoneadm"

type Zoneadmer interface{}

func Zoneadm(zonename string) Zoneadmer {
	return zadm{
		Zoneadm: &zoneadm.Zoneadm{
			ZoneName: &zonename,
		},
	}
}

type zadm struct {
	*zoneadm.Zoneadm
}

func (z *zadm) Attach() error {
	z.Zoneadm.Attach = &zoneadm.Attach{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Boot() error {
	z.Zoneadm.Boot = &zoneadm.Boot{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Clone() error {
	z.Zoneadm.Clone = &zoneadm.Clone{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Detach() error {
	z.Zoneadm.Detach = &zoneadm.Detach{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Halt() error {
	z.Zoneadm.Halt = &zoneadm.Halt{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Help() error {
	z.Zoneadm.Help = &zoneadm.Help{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Install() error {
	z.Zoneadm.Install = &zoneadm.Install{}

	return parser.MustParse(z).Run()
}

func (z *zadm) List() error {
	z.Zoneadm.List = &zoneadm.List{}

	return parser.MustParse(z).Run()
}

func (z *zadm) MarkIncomplete() error {
	z.Zoneadm.MarkIncomplete = &zoneadm.MarkIncomplete{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Move() error {
	z.Zoneadm.Move = &zoneadm.Move{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Ready() error {
	z.Zoneadm.Ready = &zoneadm.Ready{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Reboot() error {
	z.Zoneadm.Reboot = &zoneadm.Reboot{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Shutdown() error {
	z.Zoneadm.Shutdown = &zoneadm.Shutdown{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Uninstall() error {
	z.Zoneadm.Uninstall = &zoneadm.Uninstall{}

	return parser.MustParse(z).Run()
}

func (z *zadm) Verify() error {
	z.Zoneadm.Verify = &zoneadm.Verify{}

	return parser.MustParse(z).Run()
}
