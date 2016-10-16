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
	"runtime"
	"strings"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/errors"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
)

const OriginalBaseUrl = "https://api.ctl.io/"

var (
	userAgent = fmt.Sprintf("clc-go-cli-%s-%s", base.BuildVersion, runtime.GOOS)
	//this made a variable instead of a constant for testing purpoises
	BaseUrl = OriginalBaseUrl
)

type connection struct {
	bearerToken  string
	accountAlias string
	logger       *log.Logger
}

var NewConnection = func(username, password, accountAlias string, logger *log.Logger) (base.Connection, error) {
	cn := &connection{
		logger: logger,
	}
	cn.logger.Printf("Creating new connection. Username: %s", username)
	loginReq := &authentication.LoginReq{Username: username, Password: password}
	loginRes := &authentication.LoginRes{}
	err := cn.ExecuteRequest("POST", BaseUrl+"v2/authentication/login", loginReq, loginRes)
	if err != nil {
		return nil, err
	}
	cn.bearerToken = loginRes.BearerToken
	if accountAlias == "" {
		accountAlias = loginRes.AccountAlias
	}
	cn.accountAlias = accountAlias
	cn.logger.Printf("Updating connection. Bearer: %s, Alias: %s", cn.bearerToken, accountAlias)
	return cn, nil
}

func (cn *connection) GetAccountAlias() string {
	return cn.accountAlias
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
	value := reflect.ValueOf(model)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		panic("ExtractURIParams was called with the model not being a struct.")
	}
	meta := value.Type()

	var newURI = uri
	for i := 0; i < meta.NumField(); i++ {
		fieldMeta := meta.Field(i)
		uriTag := fieldMeta.Tag.Get("URIParam")
		if uriTag == "" {
			continue
		}

		field := value.FieldByIndex([]int{i})
		if uriTag == "yes" {
			if field.Kind() != reflect.String {
				panic("Fields marked by URIParam tag with value 'yes' must be strings.")
			}
			stub := fmt.Sprintf("{%s}", fieldMeta.Name)
			if strings.Contains(uri, stub) {
				newURI = strings.Replace(newURI, stub, field.String(), 1)
			}
		} else {
			if field.Kind() != reflect.Struct {
				panic("Fields marked by URIParam tag with a field name must be structs.")
			}
			for _, tag := range strings.Split(uriTag, ",") {
				subField := field.FieldByName(tag)
				if subField.Kind() != reflect.String {
					panic("Fields pointed to by a URIParam tag must be strings.")
				}
				stub := fmt.Sprintf("{%s}", tag)
				newURI = strings.Replace(newURI, stub, subField.String(), 1)
			}
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
	if BaseUrl != OriginalBaseUrl {
		url = strings.Replace(url, OriginalBaseUrl, BaseUrl, -1)
	}
	var inputData io.Reader
	if reqModel != nil {
		if verb == "POST" || verb == "PUT" || verb == "PATCH" || verb == "DELETE" {
			b, err := json.Marshal(reqModel)
			if err != nil {
				return nil, err
			}
			if string(b) != "{}" {
				inputData = bytes.NewReader(b)
			}
		}
		url = ExtractURIParams(url, reqModel)
	}
	url = strings.Replace(url, "{accountAlias}", cn.accountAlias, 1)
	url = FilterQuery(url)
	req, err = http.NewRequest(verb, url, inputData)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("CLC-ALIAS", cn.accountAlias)
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
			var payloadArray []interface{}
			if err := json.Unmarshal(resBody, &payload); err == nil {
				if errors, ok := payload["modelState"]; ok {
					bytes, err := json.Marshal(errors)
					if err == nil {
						reason = string(bytes)
					}
				} else if errors, ok := payload["message"]; ok {
					if errMsg, ok := errors.(string); ok {
						reason = errMsg
					}
				} else if errors, ok := payload["error"]; ok {
					if errMsg, ok := errors.(string); ok {
						reason = errMsg
					}
				}
			} else if err := json.Unmarshal(resBody, &payloadArray); err == nil {
				for _, p := range payloadArray {
					if pMap, ok := p.(map[string]interface{}); ok {
						if errors, ok := pMap["message"]; ok {
							if errMsg, ok := errors.(string); ok {
								reason += "\n  " + errMsg
							}
						}
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
	if stringPtr, ok := resModel.(*string); ok {
		*stringPtr = ""
		return
	}
	if binaryModel, ok := resModel.(*base.BinaryResponse); ok {
		var byts []byte
		byts, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		*binaryModel = base.BinaryResponse(byts)
		return
	}
	err = cn.decodeResponse(res, resModel)
	return
}

func (cn *connection) decodeResponse(res *http.Response, resModel interface{}) (err error) {
	if resModel == nil {
		return
	}
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(resModel)
	return
}
