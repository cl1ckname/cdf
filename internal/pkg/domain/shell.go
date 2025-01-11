package domain

import "errors"

type Shell string

const (
	FishShell Shell = "fish"
	BashShell Shell = "bash"
	ZshShell  Shell = "zsh"
)

var ErrUnsupportedShell = errors.New("unsupported shell")

func ParseShell(s string) (Shell, error) {
	shellMap := map[string]Shell{
		string(FishShell): FishShell,
		string(BashShell): BashShell,
		string(ZshShell):  ZshShell,
	}
	shell, ok := shellMap[s]
	if !ok {
		return shell, ErrUnsupportedShell
	}
	return shell, nil
}
