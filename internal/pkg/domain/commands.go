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
	stub := struct{}{}
	commandMap := map[string]struct{}{
		"help":   stub,
		"add":    stub,
		"move":   stub,
		"list":   stub,
		"remove": stub,
		"shell":  stub,
	}
	_, ok := commandMap[s]
	if !ok {
		return "", fmt.Errorf("unknown command %s: %w", s, ErrUnknownCommand)
	}
	return Command(s), nil
}
