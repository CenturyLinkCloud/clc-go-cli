package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
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
	registerCommandBase(&server.UpdateReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "update",
	})
	registerCommandBase(&server.GetReq{}, &server.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "get",
	})
	registerCommandBase(&server.GetCredentialsReq{}, &server.GetCredentialsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/credentials",
		Resource: "server",
		Command:  "get-credentials",
	})
	registerCommandBase(&server.GetImportsReq{}, &server.GetImportsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}/{LocationId}/available",
		Resource: "server",
		Command:  "get-imports",
	})
	registerCommandBase(&server.GetIPAddressReq{}, &server.GetIPAddressRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "get-public-ip-address",
	})
	registerCommandBase(&server.AddIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses",
		Resource: "server",
		Command:  "add-public-ip-address",
	})
	registerCommandBase(&server.RemoveIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "remove-public-ip-address",
	})
	registerCommandBase(&server.UpdateIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "update-public-ip-address",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOn",
		Resource: "server",
		Command:  "power-on",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOff",
		Resource: "server",
		Command:  "power-off",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/pause",
		Resource: "server",
		Command:  "pause",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/reset",
		Resource: "server",
		Command:  "reset",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/shutDown",
		Resource: "server",
		Command:  "shut-down",
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/archive",
		Resource: "server",
		Command:  "archive",
	})
	registerCommandBase(&server.RestoreReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/restore",
		Resource: "server",
		Command:  "restore",
	})
	registerCommandBase(&server.CreateSnapshotReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/createSnapshot",
		Resource: "server",
		Command:  "create-snapshot",
	})
	registerCommandBase(&server.RevertToSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}/restore",
		Resource: "server",
		Command:  "revert-to-snapshot",
	})
	registerCommandBase(&server.DeleteSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}",
		Resource: "server",
		Command:  "delete-snapshot",
	})

	registerCommandBase(&group.GetReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "get",
	})
	registerCommandBase(&group.CreateReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}",
		Resource: "group",
		Command:  "create",
	})
	registerCommandBase(&group.DeleteReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "delete",
	})

	registerCommandBase(&datacenter.ListReq{}, &[]datacenter.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}",
		Resource: "data-center",
		Command:  "list",
	})
	registerCommandBase(&datacenter.GetReq{}, &datacenter.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}/{DataCenter}?groupLinks={GroupLinks}",
		Resource: "data-center",
		Command:  "get",
	})
	registerCommandBase(&datacenter.GetDCReq{}, &datacenter.GetDCRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}/{DataCenter}/deploymentCapabilities",
		Resource: "data-center",
		Command:  "get-deployment-capabilities",
	})

	registerCustomCommand(commands.NewGroupList(commands.CommandExcInfo{
		Resource: "group",
		Command:  "list",
	}))
	registerCustomCommand(commands.NewServerList(commands.CommandExcInfo{
		Resource: "server",
		Command:  "list",
	}))
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info commands.CommandExcInfo) {
	cmd := &commands.CommandBase{
		Input:   inputModel,
		Output:  outputModel,
		ExcInfo: info,
	}
	AllCommands = append(AllCommands, cmd)
}

func registerCustomCommand(cmd base.Command) {
	AllCommands = append(AllCommands, cmd)
}
