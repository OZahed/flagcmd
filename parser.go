package flagcmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Variables
var (
	appName    string
	comamnds   map[string]*SubCommand
	baseSubCMD SubCommand
)

// errors

func init() {
	comamnds = make(map[string]*SubCommand, 0)
	baseSubCMD = SubCommand{
		SubCommandName: appName,
		SortDesc:       "",
		LongDesc:       "",
		Handler:        DefaultHandler,
		ErrorHandler:   DefaultErrHandler,
		Usage:          fmt.Sprintf(`%s --help for more help`, appName),
		FlagSet:        nil,
	}
}

func SetAppName(name string) {
	appName = name
}

func DefaultErrHandler(e SubCommandError) {
	println(fmt.Sprintf("executing command %s ejected with error %s", e.Where.SubCommandName, e.Why))
}

func DefaultHandler(s *SubCommand) error {
	if _, err := fmt.Printf("sub command %s called\n", s.SubCommandName); err != nil {
		return ErrStdOut
	}
	s.Help()
	return nil
}

func RegisterSubCommand(sc *SubCommand) error {
	sc.SubCommandName = strings.TrimSpace(sc.SubCommandName)

	if sc.Handler == nil {
		sc.Handler = DefaultHandler
	}
	if sc.ErrorHandler == nil {
		sc.ErrorHandler = DefaultErrHandler
	}
	if sc.FlagSet == nil {
		sc.FlagSet = flag.NewFlagSet(sc.SubCommandName, flag.ExitOnError)
	}
	if _, ok := comamnds[sc.SubCommandName]; ok {
		return ErrDuplicatedCMD
	}

	comamnds[sc.SubCommandName] = sc
	return nil
}

func Parse() error {
	if len(os.Args) < 2 {
		cmd, ok := comamnds[""]
		if !ok {
			allUsage()
			return ErrNoArgs
		}
		execute(cmd)
	}
	for name, cmd := range comamnds {
		if os.Args[1] == name {
			cmd.ParseFlags(os.Args)
			cmd.SetPared(true)
			execute(cmd)
			return nil
		}
	}
	return nil
}

func allUsage() {
	usg := []string{appName, " -- help \n[sub command] [OPT]\nsubcommands:\n"}
	for n, c := range comamnds {
		usg = append(usg, n, "\t", c.SortDesc, "\n")
	}
	strings.Join(usg, "  ")
	fmt.Print(usg)
}

func execute(cmd *SubCommand) {
	if err := cmd.Handler(cmd); err != nil {
		cmd.ErrorHandler(SubCommandError{
			Where: cmd,
			Why:   err,
		})
	}
}
