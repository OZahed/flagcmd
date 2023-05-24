package flagcmd

import (
	"fmt"
	"os"
)

var (
	appName string
	cmdMap  = make(map[string]*Command, 0)
)

func init() {
	appName = os.Args[0]
}

func SetAppName(name string) {
	appName = name
}

func DefaultErrHandler(e SubCommandError) {
	println(fmt.Sprintf("executing command %s ejected with error %s", e.Where.Name, e.Why))
}

func DefaultHandler(c *Command) error {
	if _, err := fmt.Printf("sub command %s called\n", c.Name); err != nil {
		return ErrStdOut
	}
	c.Help()
	return nil
}

func RegisterCommand(cmd *Command) error {
	if cmd == nil {
		return ErrNilCmd
	}
	if _, ok := cmdMap[cmd.Name]; ok {
		return ErrExists
	}

	cmdMap[cmd.Name] = cmd
	return nil
}
