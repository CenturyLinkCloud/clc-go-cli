package main

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/autocomplete"
	"os"
)

func main() {
	args := os.Args[1:]

	var output string
	if len(args) >= 1 && args[len(args)-1] == "--generate-bash-completion" {
		output = autocomplete.Run(args[:len(args)-1])
	} else {
		output = Run(args)
	}
	fmt.Printf("%s\n", output)
}
