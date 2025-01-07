package handler

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
)

type cmdmap = map[Code]Handler

type Marks struct {
	add  commands.Add
	list commands.List
}

func NewMarks(
	add commands.Add,
	list commands.List,
) Marks {
	return Marks{
		add:  add,
		list: list,
	}
}

func (m Marks) Permorm(call Call) error {
	commands := cmdmap{
		CodeAdd:  m.Add,
		CodeList: m.List,
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

func (h Marks) List(_ Args, _ Kwargs) error {
	return h.list.Execute()
}
