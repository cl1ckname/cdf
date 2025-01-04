package cmd

import (
	"fmt"
	"strings"

	"github.com/cl1ckname/cdf/internal/commands"
	"github.com/cl1ckname/cdf/internal/handlers"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
)

var (
	addCmd = commands.Command{
		Handler:     handlers.Add,
		Description: "Add new mark",
	}
)

var cmdmap = map[commands.Code]commands.Command{
	commands.Add: addCmd,
}

func Run(arguments ...string) error {
	if len(arguments) < 1 {
		return fmt.Errorf("invalid os.Args")
	}
	args, kwargs, err := parseFlags(arguments[1:])
	if len(args) < 1 {
		return fmt.Errorf("no command: %w", err)
	}
	cmd := args[0]
	argsTrimCmd := args[1:]
	code, err := commands.Parse(cmd)
	if err != nil {
		return err
	}
	handler, ok := cmdmap[code]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd)
	}

	cdfCatalog := catalog.New("/home/clickname/.config/cdf", filesystem.FS)
	storage := store.New(cdfCatalog)
	if err := storage.Init(); err != nil {
		return err
	}

	if err := handler.Handler(storage, argsTrimCmd, kwargs); err != nil {
		return err
	}

	return nil
}

func parseFlags(flags []string) (commands.Args, commands.Kwargs, error) {
	kwargs := make(commands.Kwargs)
	args := commands.Args{}
	for _, arg := range flags {
		kwarg, err := isKwarg(arg)
		if err != nil {
			return nil, nil, err
		}
		if !kwarg {
			args = append(args, arg)
			continue
		}
		key, val, err := parseKwarg(arg)
		if err != nil {
			return nil, nil, err
		}
		kwargs[key] = val
	}
	return args, kwargs, nil
}

func isKwarg(f string) (bool, error) {
	if len(f) == 0 {
		return false, fmt.Errorf("empty flag")
	}
	if f[0] != '-' {
		return false, nil
	}
	f = strings.TrimPrefix(f, "-")
	f = strings.TrimPrefix(f, "-")
	if len(f) == 0 {
		return false, fmt.Errorf("empty flag")
	}
	return true, nil
}

func parseKwarg(f string) (string, string, error) {
	isNotFlag := f[0] != '-'
	if isNotFlag {
		return "", "", fmt.Errorf("invalid flag: %s", f)
	}
	f = strings.TrimPrefix(f, "-")
	f = strings.TrimPrefix(f, "-")
	keyValue := strings.Split(f, "=")
	if len(keyValue) > 2 {
		return "", "", fmt.Errorf("invalid flag: %s", f)
	}
	if len(keyValue) == 1 {
		return keyValue[0], "true", nil
	}
	return keyValue[0], keyValue[1], nil
}
