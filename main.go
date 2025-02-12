package main

import (
	"os"

	"github.com/cl1ckname/cdf/internal/app"
)

var version string

func main() {
	sys := app.System{
		Stdout:  os.Stdout,
		Stderr:  os.Stderr,
		Args:    os.Args,
		Version: version,
	}
	if err := app.Run(sys); err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
