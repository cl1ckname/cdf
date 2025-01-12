package handler

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type cmdmap = map[Code]Handler

type Marks struct {
	add   commands.Add
	list  commands.List
	move  commands.Move
	shell commands.Shell
}

func NewMarks(
	add commands.Add,
	list commands.List,
	move commands.Move,
	shell commands.Shell,
) Marks {

	return Marks{
		add:   add,
		list:  list,
		move:  move,
		shell: shell,
	}
}

func (h Marks) Permorm(call Call) error {
	if call.Code != nil {
		return h.performCommand(*call.Code, call.Args, call.Kwargs)
	}
	return h.performFlag(call.Args)
}

func (h Marks) performCommand(code Code, args Args, kwargs Kwargs) error {
	commands := cmdmap{
		CodeAdd:   h.Add,
		CodeList:  h.List,
		CodeMove:  h.Move,
		CodeShell: h.Shell,
	}

	cmd, ok := commands[code]
	if !ok {
		return ErrUnknownCommand
	}
	return cmd(args, kwargs)
}

func (h Marks) performFlag(args Args) error {
	return nil
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

func (h Marks) Shell(args Args, kw Kwargs) error {
	if len(args) != 1 {
		return fmt.Errorf("one shell name required, got %d", len(args))
	}
	arg := args[0]
	shell, err := domain.ParseShell(arg)
	if err != nil {
		return err
	}
	return h.shell.Execute(shell)
}
