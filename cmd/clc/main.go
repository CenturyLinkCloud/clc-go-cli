package main

import (
	"fmt"
	"os"

	"github.com/centurylinkcloud/clc-go-cli/autocomplete"
	runner "github.com/centurylinkcloud/clc-go-cli/cmd_runner"
	"github.com/centurylinkcloud/clc-go-cli/state"
)

func main() {
	args := os.Args[1:]

	var output string
	if len(args) >= 1 && args[len(args)-1] == "--generate-bash-completion" {
		output = autocomplete.Run(args[:len(args)-1])
		// A shell autocomplete handler is expected to get the autocomplete
		// options from the "completion" file from the configuration folder.
		state.WriteToFile([]byte(output), "completion", 0666)
		return
	} else {
		output = runner.Run(args)
	}
	if output != "" {
		fmt.Printf("%s\n", output)
	}
}
