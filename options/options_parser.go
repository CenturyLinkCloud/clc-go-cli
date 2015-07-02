package options

func LoadOptions(parsedArgs map[string]interface{}) (*Options, error) {
	res := &Options{}
	if val, ok := parsedArgs["User"]; ok {
		delete(parsedArgs, "User")
		res.User = val.(string)
	}
	if val, ok := parsedArgs["Password"]; ok {
		delete(parsedArgs, "Password")
		res.User = val.(string)
	}
	return res, nil
}
