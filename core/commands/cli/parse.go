package cli

import "errors"

var (
	// ErrIncorrectNumCLIArguments means that the input on the command line
	// had the wrong number of input arguments (either too few or too many)
	ErrIncorrectNumCLIArguments = errors.New("cli: wrong number of CLI arguments")
)

/*
CommandType is a command given to trinity on the command line
*/
type CommandType int

const (
	// Initialize the node
	Initialize CommandType = iota
	// Start the node
	Start
	// Version of trinity
	Version
	// None of the above
	None
)

/*
Parse pulls out all the commands from the CLI input.
This is a temporary implementation just to get basic
functionality from the command line for the time being.
*/
func Parse(input []string) (CommandType, error) {
	if len(input) != 2 {
		return None, ErrIncorrectNumCLIArguments
	}
	for _, token := range input {
		if token == "init" {
			return Initialize, nil
		} else if token == "start" {
			return Start, nil
		} else if token == "version" {
			return Version, nil
		}
	}

	return None, nil
}
