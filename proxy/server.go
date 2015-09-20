package proxy

import (
	"encoding/json"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
	"net/http"
)

var (
	LoginResponse = &authentication.LoginRes{
		AccountAlias: "ALIAS",
		BearerToken:  "token",
	}
)

type Endpoint struct {
	URL      string
	Response interface{}
}

// Server mocks HTTP requests for the registered endpoints. It is a caller's
// responsibility to clean up the resources by calling CloseServer.
func Server(endpoints []Endpoint) {
	mux := acquire()
	for _, e := range endpoints {
		mux.HandleFunc(e.URL, func(w http.ResponseWriter, r *http.Request) {
			js, err := json.Marshal(e.Response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		})
	}
}

func CloseServer() {
	release()
}
