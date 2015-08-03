package connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"reflect"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/errors"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
)

//this made a variable instead of a constant for testing purpoises
var BaseUrl = "https://api.tier3.com/v2/"

type connection struct {
	bearerToken  string
	accountAlias string
	logger       *log.Logger
}

var NewConnection = func(username string, password string, logger *log.Logger) (base.Connection, error) {
	cn := &connection{
		logger: logger,
	}
	cn.logger.Printf("Creating new connection. Username: %s", username)
	loginReq := &authentication.LoginReq{username, password}
	loginRes := &authentication.LoginRes{}
	err := cn.ExecuteRequest("POST", BaseUrl+"authentication/login", loginReq, loginRes)
	if err != nil {
		return nil, err
	}
	cn.bearerToken = loginRes.BearerToken
	cn.accountAlias = loginRes.AccountAlias
	cn.logger.Printf("Updating connection. Bearer: %s, Alias: %s", cn.bearerToken, cn.accountAlias)
	return cn, nil
}

func (cn *connection) ExecuteRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error) {
	req, err := cn.prepareRequest(verb, url, reqModel)
	if err != nil {
		return
	}
	reqDump, _ := httputil.DumpRequest(req, true)
	cn.logger.Printf("Sending request: %s", reqDump)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	resDump, _ := httputil.DumpResponse(res, true)
	cn.logger.Printf("Response received: %s", resDump)
	err = cn.processResponse(res, resModel)
	return
}

func ExtractURIParams(uri string, model interface{}) string {
	meta := reflect.TypeOf(model)
	var value reflect.Value
	if meta.Kind() == reflect.Ptr {
		meta = meta.Elem()
		value = reflect.ValueOf(model).Elem()
	}
	if meta.Kind() != reflect.Struct {
		panic("ExtractURIParams was called with the model not being a struct.")
	}

	var newURI = uri
	for i := 0; i < meta.NumField(); i++ {
		fieldMeta := meta.Field(i)
		stub := fmt.Sprintf("{%s}", fieldMeta.Name)
		if fieldMeta.Tag.Get("URIParam") != "" && strings.Contains(uri, stub) {
			field := value.FieldByName(fieldMeta.Name)
			if field.Kind() != reflect.String {
				panic("Fields marked by URIParam tag must be strings.")
			}
			newURI = strings.Replace(newURI, stub, field.String(), 1)
		}
	}
	return newURI
}

func FilterQuery(raw string) string {
	uri, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	query, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		return raw
	}
	for k, v := range query {
		if len(v) == 1 && v[0] == "" {
			query.Del(k)
		}
	}
	uri.RawQuery = query.Encode()
	return uri.String()
}

func (cn *connection) prepareRequest(verb string, url string, reqModel interface{}) (req *http.Request, err error) {
	var inputData io.Reader
	if reqModel != nil {
		if verb == "POST" || verb == "PUT" || verb == "PATCH" {
			b, err := json.Marshal(reqModel)
			if err != nil {
				return nil, err
			}
			inputData = bytes.NewReader(b)
		}
		url = ExtractURIParams(url, reqModel)
	}
	url = strings.Replace(url, "{accountAlias}", cn.accountAlias, 1)
	url = FilterQuery(url)
	req, err = http.NewRequest(verb, url, inputData)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	if cn.bearerToken != "" {
		req.Header.Add("Authorization", "Bearer "+cn.bearerToken)
	}
	return req, err
}

func (cn *connection) processResponse(res *http.Response, resModel interface{}) (err error) {
	switch res.StatusCode {
	case 200, 201, 202, 204:
	default:
		reason := ""
		if resBody, err := ioutil.ReadAll(res.Body); err == nil {
			var payload map[string]interface{}
			if err := json.Unmarshal(resBody, &payload); err == nil {
				if errors, ok := payload["modelState"]; ok {
					bytes, err := json.Marshal(errors)
					if err == nil {
						reason = string(bytes)
					}
				}
			}
		}
		return &errors.ApiError{
			StatusCode:  res.StatusCode,
			ApiResponse: resModel,
			Reason:      reason,
		}
	}
	if resModel == nil {
		return err
	}
	err = cn.decodeResponse(res, resModel)
	if err != nil {
		return err
	}
	return err
}

func (cn *connection) decodeResponse(res *http.Response, resModel interface{}) (err error) {
	if resModel == nil {
		return
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(resModel)
	return
}
