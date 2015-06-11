package base

import ()

func ParseArgument(args []string) (res map[string]interface{}, err error) {
	if len(args) == 0 {
		return res, nil
	}

	err := json.Unmarshal(args[0], res)
	if ok, _ := err.(json.UnmarshalTypeError); ok {
		return nil, err
	}
	if err == nil {
		return res, nil
	}

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--") {
			return nil, fmt.Errorf("Invalid option format, option %s should start with --", args[i])
		}
		if i+1 == len(args) {
			return nil, fmt.Errorf("No value for option %s", args[i])
		}
		//single value case
		if i+2 == len(args) || strings.HasPrefix(args[i+2], "--") {
			obj, err := parseObject(args[i+1])
			if err != nil {
				return nil, err
			}
			res[strings.TrimLeft(args[i], "--")] = obj
			i++
		} else {
			array := make([]interface{}, 0)
			for j := i + 1; j < len(args) && !strings.HasPrefix(args[j], "--"); j++ {
				obj, err := parseObject(args[i+1])
				if err != nil {
					return nil, err
				}
				append(array, obj)
			}
			res[strings.TrimLeft(args[i], "--")] = array
			i = j + 1
		}
	}

}
