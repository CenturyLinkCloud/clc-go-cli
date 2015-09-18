package proxy

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"net/http"
	"net/http/httptest"
)

var (
	server            = &httptest.Server{}
	realConnectionURL = connection.BaseUrl
	realBaseURL       = base.URL
)

func acquire() *http.ServeMux {
	mux := http.NewServeMux()
	server = httptest.NewServer(mux)
	connection.BaseUrl = server.URL + "/"
	base.URL = server.URL + "/"
	return mux
}

func release() {
	server.Close()
	connection.BaseUrl = realConnectionURL
	base.URL = realBaseURL
}
