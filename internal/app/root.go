package app

import (
	"fmt"
	"strings"

	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/store"
	"github.com/cl1ckname/cdf/internal/store/catalog"
	"github.com/cl1ckname/cdf/internal/store/filesystem"
)

type Cmdmap = map[handler.Code]handler.Command

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
	code, err := handler.Parse(cmd)
	if err != nil {
		return err
	}

	cdfCatalog := catalog.New("/home/clickname/.config/cdf", filesystem.FS)
	storage := store.New(cdfCatalog)
	if err := storage.Init(); err != nil {
		return err
	}

	marksHandler := handler.NewHandler(storage)

	commands := cmdmap(marksHandler)

	command, ok := commands[code]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd)
	}
	if err := command.Handler(argsTrimCmd, kwargs); err != nil {
		return err
	}

	return nil
}

func cmdmap(h handler.Marks) Cmdmap {
	var (
		addCmd = handler.Command{
			Handler:     h.Add,
			Description: "Add new mark",
		}
	)

	var cmdmap = Cmdmap{
		handler.CodeAdd: addCmd,
	}

	return cmdmap
}

func parseFlags(flags []string) (handler.Args, handler.Kwargs, error) {
	kwargs := make(handler.Kwargs)
	args := handler.Args{}
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
