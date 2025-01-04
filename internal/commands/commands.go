package commands

import (
	"errors"
	"fmt"

	"github.com/cl1ckname/cdf/internal/state"
)

type Code int

const (
	Add Code = iota
	Move
	Remove
)

var ErrUnknownCommand = errors.New("unknown command")

func Parse(s string) (Code, error) {
	commandMap := map[string]Code{
		"add":    Add,
		"move":   Move,
		"remove": Remove,
	}
	cmd, ok := commandMap[s]
	if !ok {
		return 0, fmt.Errorf("unknown command %s: %w", s, ErrUnknownCommand)
	}
	return cmd, nil
}

type Kwargs = map[string]string
type Args = []string
type Handler = func(state state.State, args Args, kwargs Kwargs) error

type Command struct {
	Handler     Handler
	Description string
}
