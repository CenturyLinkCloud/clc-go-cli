package connection_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
)

type requestWithURIParams struct {
	param1 string `URIParam:"yes"`
	param2 string `URIParam:"yes"`
	param3 string
}

type requestWithURIParamsComposed struct {
	requestWithURIParams `URIParam:"param3"`
	param4               string `URIParam:"yes"`
}

var serveMux *http.ServeMux
var server *httptest.Server

func initTest() {
	serveMux = http.NewServeMux()
	server = httptest.NewServer(serveMux)
	connection.BaseUrl = server.URL + "/"
}

func finishTest() {
	server.Close()
}

func addHandler(t *testing.T, url string, reqModel, resModel interface{}) {
	serveMux.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		if reqModel != nil {
			reqContent, err := ioutil.ReadAll(req.Body)
			if err != nil {
				t.Error(err)
			}
			reqModel1 := reflect.New(reflect.ValueOf(reqModel).Elem().Type()).Interface()
			err = json.Unmarshal(reqContent, reqModel1)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(reqModel, reqModel1) {
				t.Errorf("Expected: %#v, obtained: %#v", reqModel, reqModel1)
			}
		}
		js, err := json.Marshal(resModel)
		if err != nil {
			t.Error(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

func newConnection(t *testing.T, registerHandler bool) (base.Connection, error) {
	if registerHandler {
		resModel := &authentication.LoginRes{AccountAlias: "ALIAS", BearerToken: "token"}
		reqModel := &authentication.LoginReq{Username: "user", Password: "password"}
		addHandler(t, "/authentication/login", reqModel, resModel)
	}
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return connection.NewConnection("user", "password", logger)
}

type testReqModel struct {
	P1, P2 string
}

type testResModel struct {
	P1, P2 string
}

func TestNewConnection(t *testing.T) {
	initTest()
	defer finishTest()
	cn, err := newConnection(t, true)
	if err != nil {
		t.Error(err)
	}
	//test that bearer token and account alias are attached to subsequent requests
	serveMux.HandleFunc("/some-url/ALIAS", func(w http.ResponseWriter, req *http.Request) {
		if h, ok := req.Header["Authorization"]; !ok || len(h) == 0 || h[0] != "Bearer token" {
			t.Errorf("Incorrect request: bearer token not found, headers: %#v", req.Header)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(""))
	})
	err = cn.ExecuteRequest("GET", connection.BaseUrl+"some-url/{accountAlias}", nil, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestNewConnectionError(t *testing.T) {
	initTest()
	defer finishTest()
	_, err := newConnection(t, false)
	if err == nil || err.Error() != "Error occured while sending request to API. Status code: 404." {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestExtractURIParams(t *testing.T) {
	r := requestWithURIParams{
		param1: "1",
		param2: "2",
		param3: "3",
	}
	uri := "http://some-url/{param1}?param2=5&param3={param3}"
	got := connection.ExtractURIParams(uri, r)
	expected := "http://some-url/1?param2=5&param3={param3}"
	if got != expected {
		t.Errorf("\nInvalid result.\nExpected: %s\nGot:%s", expected, got)
	}

	uri = "http://some-url/{param3}?x=y&param1={param1}&param2={param2}"
	got = connection.ExtractURIParams(uri, r)
	expected = "http://some-url/{param3}?x=y&param1=1&param2=2"

	if got != expected {
		t.Errorf("\nInvalid result.\nExpected: %s\nGot:%s", expected, got)
	}

	composed := requestWithURIParamsComposed{
		param4: "4",
	}
	composed.param1 = "1"
	composed.param3 = "3"
	uri = "http://some-url?x=y&param1={param1}&param2={param2}&p={param4}&p3={param3}"
	got = connection.ExtractURIParams(uri, composed)
	expected = "http://some-url?x=y&param1={param1}&param2={param2}&p=4&p3=3"

	if got != expected {
		t.Errorf("\nInvalid result.\nExpected: %s\nGot:%s", expected, got)
	}
}

func TestFilterQuery(t *testing.T) {
	cases := [][]string{
		{
			"http://some-url/?",
			"http://some-url/",
		},
		{
			"http://some-url?param1=",
			"http://some-url",
		},
		{
			"http://some-url?param1=1&",
			"http://some-url?param1=1",
		},
		{
			"http://some-url?param1=&param2=&param3=5",
			"http://some-url?param3=5",
		},
	}

	for _, testCase := range cases {
		raw, expected := testCase[0], testCase[1]
		got := connection.FilterQuery(raw)
		if got != expected {
			t.Errorf("\nInvalid result.\nExpected: %s\nGot:%s", expected, got)
		}
	}
}
