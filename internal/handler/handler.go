package handler

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
)

type cmdmap = map[Code]Handler

type Marks struct {
	add commands.Add
}

func NewMarks(add commands.Add) Marks {
	return Marks{
		add: add,
	}
}

func (m Marks) Permorm(call Call) error {
	commands := cmdmap{
		CodeAdd: m.Add,
	}

	cmd, ok := commands[call.Code]
	if !ok {
		return ErrUnknownCommand
	}
	return cmd(call.Args, call.Kwargs)
}

func (h Marks) Add(args Args, _ Kwargs) error {
	if len(args) < 2 {
		return fmt.Errorf("alias mark path required")
	}
	alias := args[0]
	path := args[1]

	return h.add.Execute(alias, path)
}
