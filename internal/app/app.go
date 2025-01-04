package app

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/commands"
)

func Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("command should be provided")
	}
	cs := args[1]
	cmd, err := commands.Parse(cs)
	if err != nil {
		return err
	}
	if cmd == commands.Add {

	}
	return fmt.Errorf("unknown: %w", err)
}
