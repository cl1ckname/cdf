package embed

import (
	_ "embed"
)

//go:embed shell.fish
var FishShell []byte
