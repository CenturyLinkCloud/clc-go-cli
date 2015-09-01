package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "--generate" {
		logger := NewLogger()
		parser := NewParser(logger)
		apiDef, err := parser.ParseApi()
		if err != nil {
			showError("Error while parsing API definition", nil)
			return
		}
		err = StoreApi(apiDef)
		if err != nil {
			showError("Error while storing API definition", nil)
		}
	}
	if len(args) != 0 {
		showError("Ussage: 'integration_tests' or 'integration_tests --generate'", nil)
		return
	}
	_, err := LoadApi()
	if err != nil {
		showError("Error while loadin API", err)
		return
	}
}

func showError(prefix string, err error) {
	if err != nil {
		fmt.Printf("%s: %s", prefix, err.Error())
	} else {
		fmt.Print(prefix)
	}
	os.Exit(1)
}
