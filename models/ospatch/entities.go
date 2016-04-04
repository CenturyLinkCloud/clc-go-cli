package ospatch

import "time"

type PatchInfo struct {
	Execution_id  string
	Status        string
	Start_time    time.Time
	End_time      time.Time
	Init_messages []map[string]interface{}
}
