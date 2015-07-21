package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

var AllCommands []base.Command = make([]base.Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
	})
	registerCommandBase(&server.DeleteReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "delete",
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
