package main

import (
	"fmt"
	"os"

	"github.com/cl1ckname/cdf/internal/app"
)

var version string

func main() {
	if err := app.Run(version, os.Args...); err != nil {
		fmt.Println(err)
	}
}
