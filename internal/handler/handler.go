package handler

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
)

type cmdmap = map[Code]Handler

type Marks struct {
	add  commands.Add
	list commands.List
	move commands.Move
}

func NewMarks(
	add commands.Add,
	list commands.List,
	move commands.Move,
) Marks {
	return Marks{
		add:  add,
		list: list,
		move: move,
	}
}

func (m Marks) Permorm(call Call) error {
	commands := cmdmap{
		CodeAdd:  m.Add,
		CodeList: m.List,
		CodeMove: m.Move,
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

const movePathWriePathkey = "cwd-file"

func (h Marks) Move(args Args, kw Kwargs) error {
	cwd, ok := kw[movePathWriePathkey]
	if !ok {
		return fmt.Errorf("--cwd-file required")
	}
	if ac := len(args); ac != 1 {
		return fmt.Errorf("required 1 arg (alias), got: %d", ac)
	}
	alias := args[0]

	path, err := h.move.Execute(alias, cwd)
	if err != nil {
		return err
	}
	fmt.Printf("you're in %s now\n", path)
	return nil
}
