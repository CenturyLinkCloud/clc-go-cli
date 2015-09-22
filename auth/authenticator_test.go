package auth_test

import (
	"github.com/centurylinkcloud/clc-go-cli/auth"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/config"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/options"
	"log"
	"os"
	"reflect"
	"testing"
)

var testCases = []testParam{
	{options: &options.Options{User: "user", Password: "password"}, res: &expectedResult{"user", "password"}},
	{options: &options.Options{User: "user"}, err: "Both --user and --password options should be specified."},
	{env: envParams{user: "user", password: "password"}, res: &expectedResult{"user", "password"}},
	{env: envParams{user: "user"}, err: "Both CLC_USER and CLC_PASSWORD environment variables should be specified."},
	{
		options: &options.Options{Profile: "p1"},
		config:  &config.Config{Profiles: map[string]config.Profile{"p1": {User: "user", Password: "password"}}},
		res:     &expectedResult{"user", "password"},
	},
	{
		config: &config.Config{Profiles: map[string]config.Profile{"p1": {User: "user", Password: "password"}}},
		env:    envParams{profile: "p1"},
		res:    &expectedResult{"user", "password"},
	},
	{
		config: &config.Config{Profiles: map[string]config.Profile{"p1": {User: "user"}}},
		env:    envParams{profile: "p1"},
		err:    "Incorrect profile 'p1'. Both User and Password should be specified.",
	},
	{
		config: &config.Config{Profiles: map[string]config.Profile{"p1": {User: "user", Password: "password"}}},
		env:    envParams{profile: "p2"},
		err:    "Profile 'p2' doesn't exist.",
	},
	{config: &config.Config{User: "user", Password: "password"}, res: &expectedResult{"user", "password"}},
	{config: &config.Config{User: "user"}, err: "Incorrect config. Both User and Password should be specified."},
	{err: "No credentials provided. Use 'clc login --help' to view list of all authentication options."},
}

type testParam struct {
	options *options.Options
	config  *config.Config
	env     envParams
	err     string
	res     *expectedResult
}

type envParams struct {
	user, password, profile string
}

type expectedResult struct {
	user, password string
}

func TestAuthenticator(t *testing.T) {
	for i, testCase := range testCases {
		t.Logf("Executing %d test case.", i+1)
		os.Setenv("CLC_USER", testCase.env.user)
		os.Setenv("CLC_PASSWORD", testCase.env.password)
		os.Setenv("CLC_PROFILE", testCase.env.profile)
		var res expectedResult
		connection.NewConnection = func(username, password, accountAlias string, logger *log.Logger) (cn base.Connection, err error) {
			res = expectedResult{username, password}
			return nil, nil
		}
		if testCase.options == nil {
			testCase.options = &options.Options{}
		}
		if testCase.config == nil {
			testCase.config = &config.Config{}
		}
		_, err := auth.AuthenticateCommand(testCase.options, testCase.config)
		if testCase.err != "" && err.Error() != testCase.err {
			t.Errorf("Invalid error. Expected: %s, obtained %s", testCase.err, err.Error())
		}
		if testCase.res != nil && reflect.DeepEqual(testCase.res, res) {
			t.Errorf("Invalid result. expected %#v, obtained %#v", testCase.res, res)
		}
	}
}
