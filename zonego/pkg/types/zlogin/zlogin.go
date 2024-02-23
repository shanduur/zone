package zlogin

const Docs = "https://illumos.org/man/1/zlogin"

func (z *Zlogin) NAME() string {
	return "zlogin"
}

type Zlogin struct {
	// The name of the zone to be entered.
	//
	// REQUIRED
	Zonename string `json:"zonename:+pos"`

	// The utility to be run in the specified zone.
	//
	// If provided, the following have NO effect:
	// - Console
	// - DisconnectOnHalt
	//
	// If provided, the following DO have effect:
	// - Safe
	// - RedirectStdin
	Utility string `json:"utility:+pos"`

	// Arguments passed to the utility.
	Arguments []string `json:"arguments:+pos"`

	// Specifies quiet mode operation. In quiet mode, extra messages indicating the the
	// function of zlogin will not be displayed, giving the possibility to present the
	// appearance that the command is running locally rather than in another zone.
	//
	// Equivalent of the '-Q' command.
	Quiet bool `json:"-Q"`

	// Safe login mode. zlogin does minimal processing and does not invoke login or su.
	// The option can not be used if a username is specified through the Login option,
	// and cannot be used with console logins. This mode should only be used to recover a
	// damaged zone when other forms of login have become impossible.
	//
	// Equivalent of the '-S' command.
	Safe bool `json:"-S"`

	// Redirect the input of zlogin to /dev/null. This option is useful when the command
	// running in the local zone and the shell which invokes zlogin both read from
	// standard input.
	//
	// Equivalent of the '-n' command.
	RedirectStdin bool `json:"-n"`

	// Disables the ability to access extended functions or to disconnect from the login
	// by using the escape sequence character.
	//
	// Equivalent of the '-E' command.
	NoEscape bool `json:"-E"`

	// Connects to the zone console.
	//
	// Equivalent of the '-C' command.
	Console bool `json:"-C"`

	// Disconnect from the console when the zone halts. This option may only be used if
	// the -C option is specified.
	//
	// Equivalent of the '-d' command.
	DisconnectOnHalt bool `json:"-d"`

	// Specifies a different escape character, c, for the key sequence used to access
	// extended functions and to disconnect from the login. The default escape character
	// is the tilde (~).
	//
	// Equivalent of the '-e [c]' command.
	EscapeCharacter *rune `json:"-e,omitempty"`

	// Specifies a different username for the zone login. If you do not use this option,
	// the zone username used is "root". This option is invalid if the Console option is
	// specified.
	//
	// Equivalent of the '-l [username]' command.
	Login *string `json:"-l,omitempty"`
}
