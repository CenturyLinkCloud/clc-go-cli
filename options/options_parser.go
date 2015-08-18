package options

import (
	"fmt"
	"reflect"
)

func ExtractFrom(parsedArgs map[string]interface{}) (*Options, error) {
	res := &Options{}
	if val, ok := parsedArgs["Help"]; ok {
		if reflect.ValueOf(val).Kind() != reflect.Invalid {
			return nil, fmt.Errorf("help option must not have a value")
		}
		delete(parsedArgs, "Help")
		res.Help = true
	}
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
	if val, ok := parsedArgs["Output"]; ok {
		delete(parsedArgs, "Output")
		res.Output = val.(string)
	}
	if val, ok := parsedArgs["Query"]; ok {
		delete(parsedArgs, "Query")
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("Query must be string.")
		}
		res.Query = val.(string)
	}
	if val, ok := parsedArgs["Filter"]; ok {
		delete(parsedArgs, "Filter")
		if _, ok := val.(string); !ok {
			return nil, fmt.Errorf("Filter must be string.")
		}
		res.Filter = val.(string)
	}
	return res, nil
}

func AreToBeSaved(parsedArgs map[string]interface{}) (bool, error) {
	if val, ok := parsedArgs["GenerateCliSkeleton"]; ok {
		if reflect.ValueOf(val).Kind() != reflect.Invalid {
			return false, fmt.Errorf("generate-cli-skeleton option must not have a value")
		}
		delete(parsedArgs, "GenerateCliSkeleton")
		return true, nil
	}
	return false, nil
}

func AreToBeTakenFromFile(parsedArgs map[string]interface{}) (bool, string, error) {
	if file, ok := parsedArgs["FromFile"]; !ok {
		return false, "", nil
	} else {
		if len(parsedArgs) != 1 {
			return false, "", fmt.Errorf("No other options are allowed to be with the from-file.")
		}
		if filepath, ok := file.(string); !ok {
			return false, "", fmt.Errorf("Invalid file path: %v.", file)
		} else {
			return true, filepath, nil
		}
	}
}
