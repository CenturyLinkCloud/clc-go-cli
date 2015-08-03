package options

import (
	"fmt"
)

func ExtractFrom(parsedArgs map[string]interface{}) (*Options, error) {
	res := &Options{}
	if val, ok := parsedArgs["User"]; ok {
		delete(parsedArgs, "User")
		res.User = val.(string)
	}
	if val, ok := parsedArgs["Password"]; ok {
		delete(parsedArgs, "Password")
		res.Password = val.(string)
	}
	if val, ok := parsedArgs["Format"]; ok {
		delete(parsedArgs, "Format")
		res.Output = val.(string)
	}
	if val, ok := parsedArgs["Query"]; ok {
		delete(parsedArgs, "Query")
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("Query must be string.")
		}
		res.Query = val.(string)
	}
	return res, nil
}
