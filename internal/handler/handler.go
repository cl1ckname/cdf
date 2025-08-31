package handler

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/pkg/commands"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

type cmdmap = map[domain.Command]Handler

const (
	ShortHelp = "h"
	LongHelp  = "help"

	ListFormat = "format"
)

type Marks struct {
	help   commands.Help
	add    commands.Add
	list   commands.List
	remove commands.Remove
	move   commands.Move
	shell  commands.Shell
}

func NewMarks(
	help commands.Help,
	add commands.Add,
	list commands.List,
	remove commands.Remove,
	move commands.Move,
	shell commands.Shell,
) Marks {
	return Marks{
		help:   help,
		add:    add,
		list:   list,
		remove: remove,
		move:   move,
		shell:  shell,
	}
}

func (h Marks) Permorm(call Call) error {
	if call.Code == nil {
		return h.performFlag(call.Args)
	}
	_, shortHelp := call.Kwargs[ShortHelp]
	_, longHelp := call.Kwargs[LongHelp]
	if shortHelp || longHelp {
		return h.help.Execute(call.Code)
	}

	return h.performCommand(*call.Code, call.Args, call.Kwargs)
}

func (h Marks) performCommand(code domain.Command, args Args, kwargs Kwargs) error {
	commands := cmdmap{
		domain.CommandHelp:   h.Help,
		domain.CommandAdd:    h.Add,
		domain.CommandList:   h.List,
		domain.CommandRemove: h.Remove,
		domain.CommandMove:   h.Move,
		domain.CommandShell:  h.Shell,
	}

	cmd, ok := commands[code]
	if !ok {
		return domain.ErrUnknownCommand
	}

	return cmd(args, kwargs)
}

func (h Marks) performFlag(args Args) error {
	return h.Help(args, nil)
}

func (h Marks) Help(args Args, _ Kwargs) error {
	if len(args) == 0 {
		return h.help.Execute(nil)
	}
	arg := args[0]
	cmd, err := domain.ParseCommand(arg)
	if err != nil {
		return err
	}
	return h.help.Execute(&cmd)
}

func (h Marks) Add(args Args, _ Kwargs) error {
	if len(args) < 2 {
		return fmt.Errorf("alias mark path required")
	}
	alias := args[0]
	path := args[1]

	return h.add.Execute(alias, path)
}

func (h Marks) List(_ Args, kw Kwargs) error {
	f := kw[ListFormat]
	format, ok := domain.ParseFormat(&f)
	if !ok {
		return fmt.Errorf("invalid format: %s", f)
	}
	return h.list.Execute(format, kw)
}

func (h Marks) Remove(args Args, _ Kwargs) error {
	if len(args) < 1 {
		return fmt.Errorf("one alias required: %w", domain.ErrInvalidAlias)
	}
	alias := args[0]
	return h.remove.Execute(alias)
}

func (h Marks) Move(args Args, kw Kwargs) error {
	if ac := len(args); ac != 1 {
		return fmt.Errorf("required 1 arg (alias), got: %d", ac)
	}
	alias := args[0]

	if err := h.move.Execute(alias); err != nil {
		return err
	}
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
