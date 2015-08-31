package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "--generate" {
		apiDef, err := ParseApi()
		if err != nil {
			showError("Error while parsing API definition")
			return
		}
		err = StoreApi(apiDef)
		if err != nil {
			showError("Error while storing API definition")
		}
	}
	if len(args) != 0 {
		showError("Ussage: 'integration_tests' or 'integration_tests --generate'")
		return
	}
	apiDef, err := LoadApi()
	if err != nil {
		showError("Error while loadin API", err)
		return
	}
}

func showError(prefix string, err error) {
	if err != nil {
		fmt.Printf("%s: %s", prefix, err.String())
	} else {
		fmt.Print(prefix)
	}
	os.Exit(1)
}
