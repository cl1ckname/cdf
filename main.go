package main

import (
	"fmt"
	"os"

	"github.com/cl1ckname/cdf/cmd"
)

func main() {
	if err := cmd.Run(os.Args...); err != nil {
		fmt.Println(err)
	}
}
