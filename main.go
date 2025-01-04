package main

import (
	"fmt"
	"os"

	"github.com/cl1ckname/cdf/internal/app"
)

func main() {
	if err := app.Run(os.Args...); err != nil {
		fmt.Println(err)
	}
}
