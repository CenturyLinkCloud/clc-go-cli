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

type PatchDetails struct {
	Execution_id  string
	Status        string
	Start_time    base.Time
	End_time      base.Time
	Duration      string
	Begin_message string
	End_message   string
	Patches       []struct {
		Start_time          base.Time
		End_time            base.Time
		Patch_begin_message string
		Patch_end_message   string
		Status              string
	}
}
