package domain

import (
	"errors"
	"fmt"
)

type Command string

const (
	CommandHelp   Command = "help"
	CommandAdd    Command = "add"
	CommandList   Command = "list"
	CommandMove   Command = "move"
	CommandRemove Command = "remove"
	CommandShell  Command = "shell"
)

var ErrUnknownCommand = errors.New("unknown command")

func ParseCommand(s string) (Command, error) {
	switch i := Command(s); i {
	case CommandHelp, CommandAdd, CommandList, CommandMove, CommandRemove, CommandShell:
		return i, nil
	default:
		return "", fmt.Errorf("unknown command %s: %w", s, ErrUnknownCommand)
	}
}
