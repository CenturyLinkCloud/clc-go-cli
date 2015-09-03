package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var apiPath string
	flag.StringVar(&apiPath, "api-path", "", "Api path")
	var generate bool
	flag.BoolVar(&generate, "generate", false, "Generate API")

	flag.Parse()

	if generate {
		logger := NewLogger()
		parser := NewParser(logger)
		apiDef, err := parser.ParseApi()
		if err != nil {
			showError("Error while parsing API definition", err)
			return
		}
		err = StoreApi(apiDef, apiPath)
		if err != nil {
			showError("Error while storing API definition", err)
		}
	} else {
		_, err := LoadApi(apiPath)
		if err != nil {
			showError("Error while loadin API", err)
			return
		}
	}
}

func showError(prefix string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", prefix, err.Error())
	} else {
		fmt.Print(prefix + "\n")
	}
	os.Exit(1)
}
