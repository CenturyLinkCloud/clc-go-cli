package formatter_provider_test

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/formatter_provider"
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"reflect"
	"testing"
)

type formatterProviderTestCase struct {
	opts options.Options
	res  base.Formatter
	err  string
	skip bool
}

var testCases = []formatterProviderTestCase{
	// Defaults to json.
	{
		opts: options.Options{},
		res:  &formatters.JsonFormatter{},
	},
	// Accepts text, json and table options.
	{
		opts: options.Options{
			Output: "json",
		},
		res: &formatters.JsonFormatter{},
	},
	{
		opts: options.Options{
			Output: "text",
		},
		res: &formatters.TextFormatter{},
	},
	{
		opts: options.Options{
			Output: "table",
		},
		res: &formatters.TableFormatter{},
	},
	// Does not accept other values.
	{
		opts: options.Options{
			Output: "xml",
		},
		err: "Unknown output 'xml'. Must be one of the following: json, table, text.",
	},
}

func FormatterProviderTest(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := formatter_provider.GetOutputFormatter(&testCase.opts)
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
