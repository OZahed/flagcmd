package flagcmd

import (
	"flag"
	"fmt"
)

type SubCommandHandler func(*Command) error
type SubCommandErrorHandler func(SubCommandError)

type SubCommandError struct {
	Where *Command
	Why   error
}

func (s SubCommandError) Error() string {
	return fmt.Sprintf("subcommand %s got error with %s,\nUsage: %s",
		s.Where.Name,
		s.Why.Error(),
		s.Where.Usage,
	)
}

type Command struct {
	Name         string                 `validate:"required,min=1"`
	Desc         string                 `validate:"required"`
	Usage        string                 `validate:"required"`
	FlagSet      *flag.FlagSet          `validate:"required"`
	Handler      SubCommandHandler      `validate:"required"`
	ErrorHandler SubCommandErrorHandler `validate:"required"`
	parsed       bool                   `validate:"-"`
}

func (s *Command) Help() string {
	return s.Usage
}

func (s *Command) Parsed() bool {
	return s.parsed
}

func (s *Command) SetPared(p bool) {
	s.parsed = p
}

func (s *Command) ParseFlags(args []string) {
	if err := s.FlagSet.Parse(args); err != nil {
		LogErrorf("Could not parse the values %w", err)
	}
}
