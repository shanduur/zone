package exec

import "errors"

const (
	// Key represents the key prefix used to identify special flags.
	Key = ":+"

	// KeyPositional represents the key prefix used to identify positional arguments.
	KeyPositional = Key + "pos"

	// KeyOnce represents the key prefix used to identify flags and subcommands that can only be used once.
	KeyOnce = Key + "once"
)

var (
	// ErrInvalidCommand represents the error returned when a command is invalid.
	ErrInvalidCommand = errors.New("invalid command, too many repetitions")
)
