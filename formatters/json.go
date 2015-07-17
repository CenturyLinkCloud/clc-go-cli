package formatters

import (
	"encoding/json"
	"fmt"
)

type JsonFormatter struct{}

func (f *JsonFormatter) FormatOutput(model interface{}) (res string, err error) {
	byteRes, err := json.MarshalIndent(model, "", "    ")
	return fmt.Sprintf("%s\n", string(byteRes)), err
}
