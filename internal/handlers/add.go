package handlers

import (
	"fmt"

	"github.com/cl1ckname/cdf/internal/commands"
	"github.com/cl1ckname/cdf/internal/state"
)

func Add(s state.State, args commands.Args, _ commands.Kwargs) error {
	if len(args) < 1 {
		return fmt.Errorf("mark path required")
	}
	path := args[0]
	return s.Append(path)
}
