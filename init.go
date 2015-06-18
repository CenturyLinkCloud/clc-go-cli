package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

var AllCommands []Command = make([]Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
	})
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info commands.CommandExcInfo) {
	cmd := &commands.CommandBase{
		Input:   inputModel,
		Output:  outputModel,
		ExcInfo: info,
	}
	AllCommands = append(AllCommands, cmd)
}
