package zoneadm

const Docs = "https://illumos.org/man/8/zoneadm"

func (z Zoneadm) NAME() string {
	return "zoneadm"
}

type Zoneadm struct {
	// Specify an alternate root (boot environment). This option can only be used in
	// conjunction with the "list" and "mark" subcommands.
	//
	// Equivalent of the '-R' flag.
	Root *string `json:"-R,omitempty"`

	// Unique identifier for a zone, as assigned by libuuid(3LIB). If this option is
	// present and the argument is a non-empty string, then the zone matching the UUID is
	// selected instead of the one named by the -z option, if such a zone is present.
	//
	// Equivalent of the '-R' flag.
	UUIDMatch *string `json:"-u,omitempty"`

	// String identifier for a zone.
	//
	// Equivalent of the '-z' flag.
	ZoneName *string `json:"-z,omitempty"`

	Attach         *Attach         `json:"attach:+once,omitempty"`
	Boot           *Boot           `json:"boot:+once,omitempty"`
	Clone          *Clone          `json:"clone:+once,omitempty"`
	Detach         *Detach         `json:"detach:+once,omitempty"`
	Halt           *Halt           `json:"halt:+once,omitempty"`
	Help           *Help           `json:"help:+once,omitempty"`
	Install        *Install        `json:"install:+once,omitempty"`
	List           *List           `json:"list:+once,omitempty"`
	MarkIncomplete *MarkIncomplete `json:"mark incomplete:+once,omitempty"`
	Move           *Move           `json:"move:+once,omitempty"`
	Ready          *Ready          `json:"ready:+once,omitempty"`
	Reboot         *Reboot         `json:"reboot:+once,omitempty"`
	Shutdown       *Shutdown       `json:"shutdown:+once,omitempty"`
	Uninstall      *Uninstall      `json:"uninstall:+once,omitempty"`
	Verify         *Verify         `json:"verify:+once,omitempty"`
}
