package formatter_provider_test

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/formatter_provider"
	"github.com/centurylinkcloud/clc-go-cli/formatters"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"reflect"
	"testing"
)

type formatterProviderTestCase struct {
	opts options.Options
	conf config.Config
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
	{
		opts: options.Options{
			Output: "csv",
		},
		res: &formatters.CsvFormatter{},
	},
	// Does not accept other values.
	{
		opts: options.Options{
			Output: "xml",
		},
		err: "Unknown output 'xml'. Must be one of the following: json, table, text, csv.",
	},
	// Reads values from config.
	{
		conf: config.Config{
			DefaultFormat: "json",
		},
		res: &formatters.JsonFormatter{},
	},
	{
		conf: config.Config{
			DefaultFormat: "table",
		},
		res: &formatters.TableFormatter{},
	},
	{
		conf: config.Config{
			DefaultFormat: "text",
		},
		res: &formatters.TextFormatter{},
	},
	// Gives options a bigger priority.
	{
		conf: config.Config{
			DefaultFormat: "table",
		},
		opts: options.Options{
			Output: "json",
		},
		res: &formatters.JsonFormatter{},
	},
	// Complains about invalid data read from the config.
	{
		conf: config.Config{
			DefaultFormat: "xml",
		},
		err: "Invalid config value for DefaultFormat: 'xml'. Must be one of the following: json, table, text.",
	},
}

func TestFormatterProvider(t *testing.T) {
	for i, testCase := range testCases {
		if testCase.skip {
			t.Logf("Skipping %d test case.", i+1)
			continue
		}
		t.Logf("Executing %d test case.", i+1)
		res, err := formatter_provider.GetOutputFormatter(&testCase.opts, &testCase.conf)
		if (err != nil || testCase.err != "") && err.Error() != testCase.err {
			t.Errorf("Invalid error.\n Expected: %s,\n obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && !reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result.\n expected %#v,\n obtained %#v", testCase.res, res)
		}
	}
}
