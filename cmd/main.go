package main

import (
	"flag"
	"fmt"
	"os"
)

var cwd = flag.String("cwd-file", "", "")

func main() {
	flag.Parse()
	fmt.Println(*cwd)
	f, err := os.Open(*cwd)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = os.WriteFile(*cwd, []byte("/home"), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
