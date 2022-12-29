package flagcmd

import (
	"flag"
	"fmt"
)

type SubCommandHandler func(*SubCommand) error
type SubCommandErrorHandler func(SubCommandError)

type SubCommandError struct {
	Where *SubCommand
	Why   error
}

func (s SubCommandError) Error() string {
	return fmt.Sprintf("subcommand %s got error with %s,\nUsage: %s",
		s.Where.SubCommandName,
		s.Why.Error(),
		s.Where.Usage,
	)
}

type SubCommand struct {
	SubCommandName string
	SortDesc       string
	LongDesc       string
	Usage          string
	FlagSet        *flag.FlagSet
	Handler        SubCommandHandler
	ErrorHandler   SubCommandErrorHandler
	parsed         bool
}

func (s *SubCommand) Help() string {
	return s.Usage
}

func (s *SubCommand) Parsed() bool {
	return s.parsed
}

func (s *SubCommand) SetPared(p bool) {
	s.parsed = p
}

func (s *SubCommand) ParseFlags(args []string) {
	if err := s.FlagSet.Parse(args); err != nil {
		LogErrorf("Could not parse the values %w", err)
	}
}
