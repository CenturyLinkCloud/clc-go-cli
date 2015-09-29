package base

import (
	"encoding/json"
	"time"
)

type NilField struct {
	Set bool
}

type Time string

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	if s == "" {
		*t = Time(s)
		return nil
	}

	var parsed time.Time
	parsed, err = time.Parse(SERVER_TIME_FORMAT, s)
	if err != nil {
		return err
	}
	*t = Time(parsed.Format(TIME_FORMAT))
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*t))
}
