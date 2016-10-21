package main

import (
	"flag"
	"os"

	integration "github.com/centurylinkcloud/clc-go-cli/integration_tests"
)

func main() {
	logger := integration.NewLogger()

	var apiPath = flag.String("api-path", "", "The path to the API file")
	flag.Parse()

	if *apiPath == "" {
		logger.Logf("ERROR: The api-path command line argument must be specified.")
		os.Exit(-1)
	}

	parser := integration.NewParser(logger)

	apiDef, err := parser.ParseApi()
	if err != nil {
		logger.Logf("Error while parsing API definition: %v", err)
		os.Exit(-2)
	}

	err = integration.StoreApi(apiDef, *apiPath)
	if err != nil {
		logger.Logf("Error while storing API definition: %v", err)
		os.Exit(-3)
	}

	os.Exit(0)
}
