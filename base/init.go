package base

import (
	"github.com/altoros/century-link-cli/commands"
	"github.com/altoros/century-link-cli/models/server"
)

var AllCommands []Command = make([]Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Method:   "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
	})
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info CommandExcInfo) {
	cmd := commandBase{
		inputModel:  inputModel,
		outputModel: outputModel,
		excInfo:     info,
	}
	AllCommands := append(AllCommands, cmd)
}
