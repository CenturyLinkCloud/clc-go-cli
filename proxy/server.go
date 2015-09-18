package proxy

import (
	"encoding/json"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
	"net/http"
	"sync"
)

var (
	LoginResponse = &authentication.LoginRes{
		AccountAlias: "ALIAS",
		BearerToken:  "token",
	}

	once        sync.Once
	andOnlyOnce sync.Once
)

type Endpoint struct {
	URL      string
	Response interface{}
}

// Server mocks HTTP requests for the registered endpoints. It is a caller's
// responsibility to clean up the resources by calling CloseServer. For the purpose
// of thread-safety only the first invokations of Server and CloseServer work.
func Server(endpoints []Endpoint) {
	once.Do(func() {
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
	})
}

func CloseServer() {
	andOnlyOnce.Do(release)
}
