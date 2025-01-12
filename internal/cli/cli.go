package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cl1ckname/cdf/internal/handler"
	"github.com/cl1ckname/cdf/internal/pkg/domain"
)

var ErrInvalidArgs = errors.New("invalid args")

func ParseCall(arguments []string) (*handler.Call, error) {
	args, kwargs, err := ParseFlags(arguments)
	if err != nil {
		return nil, err
	}
	cmd, err := parseCommand(args)
	if err != nil {
		return nil, err
	}

	if len(args) >= 1 {
		args = args[1:]
	}
	call := handler.Call{
		Kwargs: kwargs,
		Args:   args,
		Code:   cmd,
	}

	return &call, nil
}

func parseCommand(args handler.Args) (*domain.Command, error) {
	if len(args) < 1 {
		return nil, nil
	}
	cmd := args[0]
	code, err := domain.ParseCommand(cmd)
	if err != nil {
		return nil, err
	}
	return &code, nil
}

func ParseFlags(flags []string) (handler.Args, handler.Kwargs, error) {
	if len(flags) < 1 {
		return nil, nil, fmt.Errorf("args should contain program name: %w", ErrInvalidArgs)
	}
	flagsTripProgramName := flags[1:]
	kwargs := make(handler.Kwargs)
	args := handler.Args{}
	for _, arg := range flagsTripProgramName {
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
