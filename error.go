package flagcmd

import "errors"

var (
	ErrExists = errors.New("duplicated command registration")
	ErrStdOut = errors.New("unable to use stdout")
	ErrNoArgs = errors.New("not enough arguments")
	ErrNilCmd = errors.New("no command is provided")

	// Tags
	ErrNilPointer = errors.New("can not parse a nil value")
	ErrNoPtr      = errors.New("can not parse non pointer")
)
