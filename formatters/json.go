package formatters

import (
	"fmt"
	"encoding/json"
)

type JsonFormatter struct{}

func (f *JsonFormatter) FormatOutput(model interface{}) (res string, err error) {
	byteRes, err := json.MarshalIndent(model, "", "    ")
	return fmt.Sprintf("%s\n", string(byteRes)), err
}
