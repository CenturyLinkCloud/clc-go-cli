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

	logger := NewLogger()
	if generate {
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
		api, err := LoadApi(apiPath)
		if err != nil {
			showError("Error while loadin API", err)
			return
		}
		logger.Log("Api def loaded, count: %d", len(api))
		runner := NewRunner(api, logger)
		err = runner.RunTests()
		if err != nil {
			showError("", err)
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
