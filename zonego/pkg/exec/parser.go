package exec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Parser struct {
	Commander Commander // The Commander interface used to execute commands.
}

// MustParse parses the given command and returns a Runner, panicking on error.
func (p *Parser) MustParse(command interface{ NAME() string }) Runner {
	r, err := p.Parse(command)
	if err != nil {
		panic(err)
	}

	return r
}

// Parse parses the given command and returns a Runner and an error.
func (p *Parser) Parse(command interface{ NAME() string }) (Runner, error) {
	name, flags, err := parse(command.NAME(), command)
	if err != nil {
		return nil, fmt.Errorf("parsing failed: %w", err)
	}

	return p.Commander.Command(name, flags...), nil
}

// hasKey checks if the given string contains the Key prefix.
func hasKey(s string) bool {
	return strings.Contains(s, Key)
}

// multipart splits the given string on spaces and returns an array of strings.
// If the string is empty, it returns itself as a single-element array.
func multipart(s string) []string {
	x := strings.Split(s, " ")
	if len(x) <= 0 {
		return []string{s}
	}

	return x
}

// stripKeys removes the Key prefix from the given string.
// If the string doesn't contain the Key prefix, it returns the original string.
func stripKeys(s string) string {
	x := strings.Split(s, Key)
	if len(x) <= 0 {
		return s
	}

	return x[0]
}

// parse parses the given command name and value and returns the parsed command name,
// a slice of flags, and an error.
func parse(name string, command any) (string, []string, error) {
	buf := new(bytes.Buffer) // Buffer to store the encoded command
	once := false            // Flag to track if a KeyOnce flag has been used

	// Encode the command to JSON
	err := json.NewEncoder(buf).Encode(command)
	if err != nil {
		return "", nil, fmt.Errorf("unable to encode command %T: %w", command, err)
	}

	// Decode the JSON into a map
	ast := make(map[string]any)
	err = json.NewDecoder(buf).Decode(&ast)
	if err != nil {
		return "", nil, fmt.Errorf("unable to decode ast %T: %w", ast, err)
	}

	flags := []string{}               // Slice to store flags
	posArgs := make(map[int][]string) // Map to store positional arguments

	for k, v := range ast {
		if strings.Contains(k, KeyOnce) {
			if once {
				return "", nil, ErrInvalidCommand
			}

			once = true
		}

		switch v := v.(type) {
		case bool:
			if v {
				flags = append(flags, k)
			}

		case string:
			if strings.Contains(k, ":+pos") {
				posArgs[1] = append(posArgs[1], v)
			} else {
				flags = append(flags, k, v)
			}

		case *string:
			if v != nil {
				if strings.Contains(k, KeyPositional) {
					posArgs[1] = append(posArgs[1], *v)
				} else {
					flags = append(flags, k, *v)
				}
			}

		case []string:
			if strings.Contains(k, KeyPositional) {
				posArgs[1] = append(posArgs[1], v...)
			} else {
				flags = append(flags, k)
				flags = append(flags, v...)
			}

		case map[string]any:
			n, fls, err := parse(k, v)
			if err != nil {
				return name, flags, err
			}

			posArgs[1] = append(posArgs[1], multipart(n)...)
			posArgs[1] = append(posArgs[1], fls...)
		}
	}

	for _, pos := range posArgs {
		flags = append(flags, pos...)
	}

	if hasKey(name) {
		name = stripKeys(name)
	}

	return name, flags, nil
}
