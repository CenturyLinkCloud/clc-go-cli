package proxy

import (
	"encoding/json"
	"github.com/centurylinkcloud/clc-go-cli/models/authentication"
	"net/http"
)

// Login mocks the authentication facility in that every authentication request
// completes successfully. It is a caller's responsibility to call CloseLogin
// to return to the original state. Useful for testing.
func Login() {
	mux := acquire()

	mux.HandleFunc("/authentication/login", func(w http.ResponseWriter, r *http.Request) {
		res := &authentication.LoginRes{AccountAlias: "ALIAS", BearerToken: "token"}
		js, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	})
}

// CloseLogin releases the control over the server connections acquired by Login.
func CloseLogin() {
	release()
}
