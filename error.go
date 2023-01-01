package flagcmd

import "errors"

var (
	ErrExists = errors.New("duplicated command registration")
	ErrStdOut = errors.New("unable to use stdout")
	ErrNoArgs = errors.New("not enough arguments")
	ErrNilCmd = errors.New("no command is provided")
)
