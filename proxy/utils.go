package proxy

import (
	"github.com/centurylinkcloud/clc-go-cli/connection"
	"net/http"
	"net/http/httptest"
)

var (
	server            = &httptest.Server{}
	realConnectionURL = connection.BaseUrl
)

func acquire() *http.ServeMux {
	mux := http.NewServeMux()
	server = httptest.NewServer(mux)
	connection.BaseUrl = server.URL + "/"
	return mux
}

func release() {
	server.Close()
	connection.BaseUrl = realConnectionURL
}
