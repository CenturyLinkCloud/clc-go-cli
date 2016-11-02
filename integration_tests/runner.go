package integration_tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	cli "github.com/centurylinkcloud/clc-go-cli"
	"github.com/centurylinkcloud/clc-go-cli/base"
	clc "github.com/centurylinkcloud/clc-go-cli/cmd_runner"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
	arg_parser "github.com/centurylinkcloud/clc-go-cli/parser"
)

// Runner represents a integration test suite
type Runner interface {
	RunTests(t *testing.T) error
}

type runner struct {
	api      []*ApiDef
	serveMux *http.ServeMux
	server   *httptest.Server
	trace    bool
}

// NewRunner will create a new instance of a Runner
// for a supplied api definition.
func NewRunner(api []*ApiDef, trace bool) Runner {
	r := &runner{}
	r.api = api
	r.addServeMux()
	r.trace = trace
	return r
}

func (r *runner) RunTests(t *testing.T) error {
	if r.trace {
		os.Setenv("CLC_TRACE", "true")
	}

	for _, command := range cli.AllCommands {
		if cmdBase, ok := command.(*commands.CommandBase); ok {
			var subTestName = fmt.Sprintf("%s %s", cmdBase.ExcInfo.Resource, cmdBase.ExcInfo.Command)
			t.Run(subTestName, func(t *testing.T) {
				err := r.TestCommand(cmdBase, t)
				if err != nil {
					t.Errorf("Error executing test: %v", err)
					return
				}

				t.Logf("Test completed without error")

			})
		}
	}

	return nil
}

func (r *runner) TestCommand(cmd *commands.CommandBase, t *testing.T) (err error) {
	if cmd.ExcInfo.Verb == "PATCH" {
		t.Skipf("HTTP PATCH isn't supported by the integration tests")
		return nil
	}
	if strings.Contains(cmd.ExcInfo.Url, "?") {
		t.Skipf("URLs with a query string aren't supported by the integration tests")
		return nil
	}

	apiDef := r.findApiDef(cmd.ExcInfo.Url, cmd.ExcInfo.Verb)
	if apiDef == nil {
		t.Skipf("Api definition for url %s and method %s not found", cmd.ExcInfo.Url, cmd.ExcInfo.Verb)
		return nil
	}

	//TODO: Handle commands that are wrappers around execute package. Ignore the commands for now
	if r.isExecutePackageWrapper(cmd) {
		t.Skipf("Commands that are a wrapper around execute package aren't supported by the integration tests.")
		return nil
	}

	args := []string{cmd.ExcInfo.Resource, cmd.ExcInfo.Command}
	defaultID := "some-id"
	r.initialModifyContent(apiDef)
	t.Logf("API url: %s\n", apiDef.Url)
	url := apiDef.Url
	url = strings.Replace(url, "https://api.ctl.io", "", -1)
	url = strings.Replace(url, "{accountAlias}", "ALIAS", -1)
	var contentExampleString []byte
	if apiDef.ContentExample != nil {
		contentExampleString, err = json.Marshal(apiDef.ContentExample)
		if err != nil {
			return err
		}
	}
	var resExampleString []byte
	if apiDef.ResExample != nil {
		r.modifyResExample(apiDef)
		resExampleString, err = json.Marshal(apiDef.ResExample)
		if err != nil {
			return err
		}
	}
	modifiedContentExampleString, err := r.postModifyContent(apiDef)
	if err != nil {
		return err
	}
	if apiDef.ContentExample != nil {
		args = append(args, string(modifiedContentExampleString))
	}
	url = strings.ToLower(url)
	for _, param := range apiDef.UrlParameters {
		paramName := strings.Replace(param.Name, "IP", "Ip", -1)
		paramName = strings.Replace(paramName, "ID", "Id", -1)
		if paramName != "AccountAlias" && paramName != "LocationId" && !strings.EqualFold(paramName, "sourceAccountAlias") {
			args = append(args, arg_parser.DenormalizePropertyName(paramName), defaultID)
			url = strings.Replace(url, "{"+strings.ToLower(paramName)+"}", defaultID, -1)
		} else if paramName == "LocationId" {
			args = append(args, "--data-center", defaultID)
			url = strings.Replace(url, "{locationid}", defaultID, -1)
		}
	}
	args = append(args, "--user", "user", "--password", "password")
	err = r.addHandler(url, string(resExampleString), func(req string) error {
		t.Logf("Json1: %s", string(contentExampleString))
		t.Logf("Json2: %s", req)
		return r.compareJson(string(contentExampleString), req)
	})
	if err != nil {
		return err
	}

	t.Logf("Args: %v", args)
	res := clc.Run(args)
	t.Logf("Result received: %s", res)
	if res == "" {
		return nil
	}
	obj := new(interface{})
	err = json.Unmarshal([]byte(res), obj)
	if err != nil {
		//if we can't unmarshal result - this is most likely a error message
		return fmt.Errorf(res)
	}
	return r.deepCompareObjects("", r.postModifyResExample(apiDef.ResExample), *obj)
}

func (r *runner) addServeMux() {
	baseMux := http.NewServeMux()
	r.serveMux = http.NewServeMux()
	r.server = httptest.NewServer(baseMux)
	connection.BaseUrl = r.server.URL + "/"

	baseMux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		req.URL.Path = strings.ToLower(req.URL.Path)
		r.serveMux.ServeHTTP(w, req)
	}))
}

func (r *runner) addLoginHandler() error {
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
	r.addHandlerBase("/v2/authentication/login", string(response), checker)
	return nil
}

func (r *runner) addHandler(url string, response string, checker func(string) error) error {
	r.addServeMux()
	err := r.addLoginHandler()
	if err != nil {
		return err
	}
	r.addHandlerBase(url, response, checker)
	return nil
}

func (r *runner) addHandlerBase(url string, response string, checker func(string) error) {
	url = strings.ToLower(url)
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

func (r *runner) findApiDef(url, method string) *ApiDef {
	for _, apiDef := range r.api {
		apiDef.Url = strings.Replace(apiDef.Url, "{sourceAccountAlias}", "{accountAlias}", -1)
		apiURL := strings.Replace(apiDef.Url, "locationId", "DataCenter", -1)
		apiURL = strings.Replace(apiURL, "Network", "NetworkId", -1)
		if strings.EqualFold(apiURL, url) && apiDef.Method == method {
			return apiDef
		}
	}
	return nil
}

func (r *runner) postModifyResExample(obj interface{}) interface{} {
	switch obj.(type) {
	case map[string]interface{}:
		m := obj.(map[string]interface{})
		for key, value := range m {
			m[key] = r.postModifyResExample(value)
		}
	case []interface{}:
		array := obj.([]interface{})
		for i, child := range array {
			array[i] = r.postModifyResExample(child)
		}
	case string:
		str := obj.(string)
		t, err := time.Parse(base.SERVER_TIME_FORMAT, str)
		if err == nil {
			return t.Format(base.TIME_FORMAT)
		}
	}
	return obj
}

func (r *runner) modifyResExample(apiDef *ApiDef) {
	additionalProperties := []additionalProperty{
		{"POST", "https://api.ctl.io/v2/groups/{accountAlias}", "serversCount", 1},
		{"GET", "https://api.ctl.io/v2/servers/{accountAlias}/{serverId}", "os", "some-os"},
	}
	for _, prop := range additionalProperties {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			if obj, ok := apiDef.ResExample.(map[string]interface{}); ok {
				obj[prop.Name] = prop.Value
			}
			if array, ok := apiDef.ResExample.([]interface{}); ok {
				for _, obj := range array {
					obj.(map[string]interface{})[prop.Name] = prop.Value
				}
			}
		}
	}
}

func (r *runner) initialModifyContent(apiDef *ApiDef) {
	additionalProperties := []additionalProperty{
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}", "isManagedBackup", true},
	}
	for _, prop := range additionalProperties {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			apiDef.ContentExample.(map[string]interface{})[prop.Name] = prop.Value
		}
	}

	missedExamples := []missedExample{
		{"POST", "https://api.ctl.io/v2/operations/{accountAlias}/servers/startMaintenance", []interface{}{"WA1ALIASWB01", "WA1ALIASWB02"}},
		{"POST", "https://api.ctl.io/v2/groups/{accountAlias}/{groupId}/restore", map[string]interface{}{"targetGroupId": "WA1ALIASWB02"}},
	}
	for _, prop := range missedExamples {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			apiDef.ContentExample = prop.Example
		}
	}
	urlProperties := []convertProperty{
		{"PUT", "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{dataCenter}/{Network}", "Network", "NetworkId"},
		{"POST", "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{dataCenter}/{Network}/release", "Network", "NetworkId"},
		{"PUT", "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{dataCenter}/{firewallPolicy}", "DestinationAccountAlias", "FirewallPolicy"},
	}
	for _, prop := range urlProperties {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			for _, param := range apiDef.UrlParameters {
				if param.Name == prop.OldName {
					param.Name = prop.NewName
				}
			}
			apiDef.Url = strings.Replace(apiDef.Url, prop.OldName, prop.NewName, -1)
		}
	}
}

func (r *runner) isExecutePackageWrapper(cmd *commands.CommandBase) bool {
	wrappers := []ExecutePackageWrapper{
		{"os-patch", "apply"},
	}

	for _, wrapper := range wrappers {
		if cmd.ExcInfo.Resource == wrapper.Resource && cmd.ExcInfo.Command == wrapper.Command {
			return true
		}
	}
	return false
}

func (r *runner) postModifyContent(apiDef *ApiDef) (string, error) {
	contentExample := apiDef.ContentExample
	if array, ok := contentExample.([]interface{}); ok {
		if _, ok := array[0].(string); ok {
			contentExample = map[string]interface{}{"serverIds": array}
		} else {
			contentExample = map[string]interface{}{"nodes": array}
		}
	}

	for _, param := range apiDef.ContentParameters {
		if param.Type == "dateTime" {
			contentExample.(map[string]interface{})[param.Name] = strings.Replace(contentExample.(map[string]interface{})[param.Name].(string), "T", " ", -1)
			contentExample.(map[string]interface{})[param.Name] = strings.Replace(contentExample.(map[string]interface{})[param.Name].(string), "Z", "", -1)
		}
		if param.Name == "source" {
			v := contentExample.(map[string]interface{})["source"]
			contentExample.(map[string]interface{})["sources"] = v
			delete(contentExample.(map[string]interface{}), "source")
		}
		if param.Name == "destination" {
			v := contentExample.(map[string]interface{})["destination"]
			contentExample.(map[string]interface{})["destinations"] = v
			delete(contentExample.(map[string]interface{}), "destination")
		}
	}
	exampleProperties := []convertProperty{
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}", "password", "rootPassword"},
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}", "memoryGB", "memoryGb"},
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}", "isManagedOS", "isManagedOs"},
		{"POST", "https://api.ctl.io/v2/servers/{accountAlias}/{serverId}/publicIPAddresses", "cidr", "CIDR"},
		{"PUT", "https://api.ctl.io/v2/servers/{accountAlias}/{serverId}/publicIPAddresses/{publicIP}", "cidr", "CIDR"},
		{"POST", "https://api.ctl.io/v2/vmImport/{accountAlias}", "password", "rootPassword"},
		{"POST", "https://api.ctl.io/v2/vmImport/{accountAlias}", "memoryGB", "memoryGb"},
		{"POST", "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}", "location", "dataCenter"},
		{"PUT", "https://api.ctl.io/v2/servers/{accountAlias}/{serverId}/cpuAutoscalePolicy", "id", "policyId"},
		{"PUT", "https://api.ctl.io/v2/groups/{accountAlias}/{groupId}/horizontalAutoscalePolicy/", "loadBalancerPool", "loadBalancer"},
		{"POST", "https://api.ctl.io/v2/groups/{accountAlias}/{groupId}/defaults", "memoryGB", "memoryGb"},
		{"POST", "https://api.ctl.io/v2/operations/{accountAlias}/servers/executePackage", "servers", "serverIds"},
	}
	data, err := json.Marshal(contentExample)
	if err != nil {
		return "", err
	}
	strData := string(data)
	for _, prop := range exampleProperties {
		if apiDef.Method == prop.Method && apiDef.Url == prop.Url {
			strData = strings.Replace(strData, prop.OldName, prop.NewName, -1)
		}
	}
	return strData, err
}

func (r *runner) compareJson(json1, json2 string) error {
	if strings.TrimSpace(json1) == "" && strings.TrimSpace(json2) == "" {
		return nil
	}
	obj1 := new(interface{})
	obj2 := new(interface{})
	err := json.Unmarshal([]byte(json1), obj1)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(json2), obj2)
	if err != nil {
		return err
	}
	return r.deepCompareObjects("", *obj1, *obj2)
}

func (r *runner) deepCompareObjects(prefix string, obj1 interface{}, obj2 interface{}) error {
	if obj1 == nil && obj2 == nil {
		return nil
	}
	if obj1 == nil || obj2 == nil {
		if array, ok := obj1.([]interface{}); ok && len(array) == 0 {
			return nil
		}
		if array, ok := obj2.([]interface{}); ok && len(array) == 0 {
			return nil
		}
		return fmt.Errorf("Mismatch in property %s. Values: \n%v \n%v", prefix, obj1, obj2)
	}
	switch obj1.(type) {
	case string, float64, bool:
		if fmt.Sprintf("%v", obj1) != fmt.Sprintf("%v", obj2) {
			return fmt.Errorf("Mismatch in property %s. Values: %v %v", prefix, obj1, obj2)
		}
		return nil
	case []interface{}:
		array1 := obj1.([]interface{})
		array2 := obj2.([]interface{})
		if len(array1) == 0 && len(array2) == 0 {
			return nil
		}
		if len(array1) != len(array2) {
			return fmt.Errorf("Different array length for property %s - %d %d. Values %v %v", prefix, len(array1), len(array2), obj1, obj2)
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
		//defferent map length should be considered a valid case
		/*if len(map1) != len(map2) {
			var keys1, keys2 []string
			for key1, _ := range map1 {
				keys1 = append(keys1, key1)
			}
			for key2, _ := range map2 {
				keys2 = append(keys2, key2)
			}
			return fmt.Errorf("Different map length for property %s - %d %d. Keys:\n%v \n%v", prefix, len(map1), len(map2), keys1, keys2)
		}*/
		for key, value := range map1 {
			if key == "links" {
				continue
			}
			var correspondingValue interface{}
			for key2, val2 := range map2 {
				if strings.ToLower(key) == strings.ToLower(key2) {
					correspondingValue = val2
					break
				}
			}
			res := r.deepCompareObjects(prefix+"."+key, value, correspondingValue)
			if res != nil {
				return res
			}
		}
	}
	return nil
}

type convertProperty struct {
	Method, Url, OldName, NewName string
}

type additionalProperty struct {
	Method, Url, Name string
	Value             interface{}
}

type missedExample struct {
	Method, Url string
	Example     interface{}
}

type ExecutePackageWrapper struct {
	Resource, Command string
}
