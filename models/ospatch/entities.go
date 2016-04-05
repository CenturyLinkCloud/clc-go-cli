package ospatch

import "github.com/centurylinkcloud/clc-go-cli/base"

type PatchInfo struct {
	Execution_id  string
	Status        string
	Start_time    base.Time
	End_time      base.Time
	Init_messages []struct {
		Start_time         base.Time
		End_time           base.Time
		Init_begin_message string
		Init_end_message   string
	}
}
