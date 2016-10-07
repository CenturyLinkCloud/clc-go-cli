// +build integration

package integration_tests

import (
	"flag"
	"os"
	"testing"
)

var (
	apiPath  = flag.String("api-path", "", "The path to the API file")
	generate = flag.Bool("generate", false, "Flag to indicate we should generate the API file")
	clcTrace = flag.Bool("clc-trace", false, "Output trace statements for calls to CLC")
)

func TestMain(m *testing.M) {
	flag.Parse()

	if *generate {
		//Generate the API file
		logger := NewLogger()
		parser := NewParser(logger)
		apiDef, err := parser.ParseApi()
		if err != nil {
			logger.Logf("Error while parsing API definition: %v", err)
			os.Exit(-1)
		}
		err = StoreApi(apiDef, *apiPath)
		if err != nil {
			logger.Logf("Error while storing API definition: %v", err)
			os.Exit(-2)
		}
		os.Exit(0)
	}

	result := m.Run()

	os.Exit(result)

}

func TestCommands(t *testing.T) {
	api, err := LoadApi(*apiPath)
	if err != nil {
		t.Errorf("Error while loading API: %v", err)
		t.Fail()
		return
	}
	t.Logf("Api def loaded, count: %d", len(api))

	runner := NewRunner(api, *clcTrace)
	err = runner.RunTests(t)

	if err != nil {
		t.Errorf("%s\n", err.Error())
		t.Fail()
	}
}
