package flagcmd

import "errors"

var (
	ErrDuplicatedCMD = errors.New("duplicated command registration")
	ErrStdOut        = errors.New("unable to use stdout")
	ErrNoArgs        = errors.New("not enough arguments")
	ErrNoCommand     = errors.New("no command is provided")
)
