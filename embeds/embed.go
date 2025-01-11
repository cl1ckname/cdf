package embeds

import (
	_ "embed"
)

var (
	//go:embed shell.fish
	FishShell []byte

	//go:embed shell.bash
	BashShell []byte
)
