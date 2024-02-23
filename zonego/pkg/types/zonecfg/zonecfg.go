package zonecfg

const Docs = "https://illumos.org/man/8/zonecfg"

func (z *Zonecfg) NAME() string {
	return "zonecfg"
}

type Zonecfg struct{}
