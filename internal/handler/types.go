package handler

import (
	"errors"
	"fmt"
)

type Code int

const (
	CodeAdd Code = iota
	CodeList
	CodeMove
	CodeRemove
	CodeShell
)

const RecordSeparator = "="

var ErrUnknownCommand = errors.New("unknown command")

func Parse(s string) (Code, error) {
	commandMap := map[string]Code{
		"add":    CodeAdd,
		"move":   CodeMove,
		"list":   CodeList,
		"remove": CodeRemove,
		"shell":  CodeShell,
	}
	cmd, ok := commandMap[s]
	if !ok {
		return 0, fmt.Errorf("unknown command %s: %w", s, ErrUnknownCommand)
	}
	return cmd, nil
}

type Kwargs = map[string]string
type Args = []string
type Handler = func(args Args, kwargs Kwargs) error

type Command struct {
	Handler     Handler
	Description string
}

type Call struct {
	Kwargs Kwargs
	Args   Args
	Code   *Code
}
