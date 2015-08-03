package options

import (
	"fmt"
	"reflect"
)

func ExtractFrom(parsedArgs map[string]interface{}) (*Options, error) {
	res := &Options{}
	if val, ok := parsedArgs["User"]; ok {
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("User must be string.")
		}
		delete(parsedArgs, "User")
		res.User = val.(string)
	}
	if val, ok := parsedArgs["Password"]; ok {
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("Password must be string")
		}
		delete(parsedArgs, "Password")
		res.Password = val.(string)
	}
	if val, ok := parsedArgs["Profile"]; ok {
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("Profile must be string.")
		}
		delete(parsedArgs, "Profile")
		res.Profile = val.(string)
	}
	if val, ok := parsedArgs["Trace"]; ok {
		if reflect.ValueOf(val).Kind() != reflect.Invalid {
			return nil, fmt.Errorf("trace option must not have a value")
		}
		delete(parsedArgs, "Trace")
		res.Trace = true
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
