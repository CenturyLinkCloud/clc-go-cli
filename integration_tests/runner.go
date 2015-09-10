package main

import (
	"encoding/json"
	"fmt"
	cli "github.com/centurylinkcloud/clc-go-cli"
	clc "github.com/centurylinkcloud/clc-go-cli/cmd/clc"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
	arg_parser "github.com/centurylinkcloud/clc-go-cli/parser"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Runner interface {
	RunTests() error
}

func NewRunner(api []*ApiDef, logger Logger) Runner {
	r := &runner{}
	r.logger = logger
	r.api = api
	r.serveMux = http.NewServeMux()
	r.server = httptest.NewServer(r.serveMux)
	connection.BaseUrl = r.server.URL + "/"
	return r
}

type runner struct {
	api      []*ApiDef
	logger   Logger
	serveMux *http.ServeMux
	server   *httptest.Server
}

func (r *runner) RunTests() error {
	os.Setenv("CLC_TRACE", "true")
	err := r.addLoginHandler()
	if err != nil {
		return err
	}
	for _, command := range cli.AllCommands {
		if cmdBase, ok := command.(*commands.CommandBase); ok {
			err := r.TestCommand(cmdBase)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *runner) addLoginHandler() error {
	strErr := clc.Run([]string{"login", "--user", "user", "--password", "password"})
	if strErr != "" {
		return fmt.Errorf(strErr)
	}
	resModel := &authentication.LoginRes{AccountAlias: "ALIAS", BearerToken: "token"}
	response, err := json.Marshal(resModel)
	if err != nil {
		return err
	}
	checker := func(req string) error {
		reqModel := &authentication.LoginReq{}
		err := json.Unmarshal([]byte(req), &reqModel)
		if err != nil {
			return err
		}
		if reqModel.Username != "user" || reqModel.Password != "password" {
			return fmt.Errorf("Incorrect request model: %#v", reqModel)
		}
		return nil
	}
	r.addHandler("/authentication/login", string(response), checker)
	return nil
}

func (r *runner) addHandler(url string, response string, checker func(string) error) {
	r.serveMux.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		reqContent, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		if checker != nil {
			err := checker(string(reqContent))
			if err != nil {
				panic(err)
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})
}

func (r *runner) findApiDef(url, method string) (*ApiDef, error) {
	for _, apiDef := range r.api {
		if apiDef.Url == url && apiDef.Method == method {
			return apiDef, nil
		}
	}
	return nil, fmt.Errorf("Api definition for url %s not found", url)
}

func (r *runner) TestCommand(cmd *commands.CommandBase) (err error) {
	r.logger.Log("------- Testing command %s %s", cmd.ExcInfo.Resource, cmd.ExcInfo.Command)
	defer func() {
		if panicErr := recover(); panicErr != nil {
			err = panicErr.(error)
		}
	}()
	apiDef, err := r.findApiDef(cmd.ExcInfo.Url, cmd.ExcInfo.Verb)
	if err != nil {
		return err
	}
	args := []string{cmd.ExcInfo.Resource, cmd.ExcInfo.Command}
	defaultId := "some-id"
	url := apiDef.Url
	url = strings.Replace(url, "https://api.ctl.io/v2", "", -1)
	url = strings.Replace(url, "{accountAlias}", "ALIAS", -1)
	for _, param := range apiDef.UrlParameters {
		if param.Name != "AccountAlias" {
			args = append(args, arg_parser.DenormalizePropertyName(param.Name), defaultId)
			url = strings.Replace(url, "{"+param.Name+"}", defaultId, -1)
		}
	}
	contentExampleString, err := json.Marshal(apiDef.ContentExample)
	if err != nil {
		return err
	}
	resExampleString, err := json.Marshal(apiDef.ResExample)
	if err != nil {
		return err
	}
	modifiedContentExampleString, err := r.modifyContentParams(apiDef)
	if err != nil {
		return err
	}
	if apiDef.ContentExample != nil {
		args = append(args, string(modifiedContentExampleString))
	}
	r.addHandler(url, string(resExampleString), func(req string) error {
		return r.compareJson(req, string(contentExampleString))
	})
	strErr := clc.Run(args)
	if strErr != "" {
		return fmt.Errorf(strErr)
	}
	return nil
}

func (r *runner) modifyContentParams(apiDef *ApiDef) (string, error) {
	type convertProperty struct {
		Method, Url, OldName, NewName string
	}
	properties := []convertProperty{
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}", "password", "rootPassword"},
	}
	for _, prop := range properties {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			for _, param := range apiDef.ContentParameters {
				if param.Name == prop.OldName {
					mapExample := apiDef.ContentExample.(map[string]interface{})
					mapExample[prop.NewName] = mapExample[prop.OldName]
					delete(mapExample, prop.OldName)
				}
			}
		}
	}
	data, err := json.Marshal(apiDef.ContentExample)
	return string(data), err
}

func (r *runner) compareJson(json1, json2 string) error {
	r.logger.Log("Request: %s", json1)
	r.logger.Log("Content Example: %s", json1)
	obj1 := map[string]interface{}{}
	obj2 := map[string]interface{}{}
	err := json.Unmarshal([]byte(json1), &obj1)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(json2), &obj2)
	if err != nil {
		return err
	}
	return r.deepCompareObjects("", obj1, obj2)
}

func (r *runner) deepCompareObjects(prefix string, obj1 interface{}, obj2 interface{}) error {
	if obj1 == nil && obj2 == nil {
		return nil
	}
	if obj1 == nil || obj2 == nil {
		return fmt.Errorf("Mistmatch in property %s. Values: %v %v", prefix, obj1, obj2)
	}
	switch obj1.(type) {
	case string, float64, bool:
		if obj1 != obj2 {
			return fmt.Errorf("Mistmatch in property %s. Values: %v %v", prefix, obj1, obj2)
		}
		return nil
	case []interface{}:
		array1 := obj1.([]interface{})
		array2 := obj2.([]interface{})
		if len(array1) == 0 && len(array2) == 0 {
			return nil
		}
		if len(array1) != len(array2) {
			return fmt.Errorf("Different array length for property %s - %b %b. Values %v %v", prefix, len(array1), len(array2), obj1, obj2)
		}
		for i := 0; i < len(array1); i++ {
			res := r.deepCompareObjects(prefix+"["+strconv.Itoa(i)+"]", array1[i], array2[i])
			if res != nil {
				return res
			}
		}
	case map[string]interface{}:
		map1 := obj1.(map[string]interface{})
		map2 := obj2.(map[string]interface{})
		for key, value := range map1 {
			//all property names in modesl starts with uppercase, but in returned json they can be lowercase
			//so we make a conversion here
			upperKey := []rune(key)
			upperKey[0] = unicode.ToUpper(upperKey[0])
			key2 := string(upperKey)
			val2 := map2[key2]
			//this is for case when response object by itself contains a map
			//then, keys in this map are not uppercased
			if val2 == nil {
				val2 = map2[key]
			}
			res := r.deepCompareObjects(prefix+"."+key, value, val2)
			if res != nil {
				return res
			}
		}
	}
	return nil
}
