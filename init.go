package cli

import (
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/help"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/affinity"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"github.com/centurylinkcloud/clc-go-cli/models/balancer"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/firewall"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/models/network"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

var AllCommands []base.Command = make([]base.Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
		Help: help.Command{
			Brief: `Creates a new server. Use this API operation when you want to create a new server from a standard or custom template, or clone an existing server.`,
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server"},
				},
				{
					"--group-id",
					[]string{"ID of the parent group."},
				},
				{
					"--source-server-id",
					[]string{"ID of the server to use a source."},
				},
				{
					"--is-managed-os",
					[]string{
						"Whether to create the server as managed or not. Default is false.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--is-managed-backup",
					[]string{
						"Whether to add managed backup to the server. Must be a managed OS server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--ip-address",
					[]string{
						"IP address to assign to the server. If not provided, one will be assigned automatically.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--root-password",
					[]string{"Password of administrator or root user on server. If not provided, one will be generated automatically."},
				},
				{
					"--source-server-password",
					[]string{
						"Password of the source server, used only when creating a clone from an existing server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--cpu",
					[]string{"Number of processors to configure the server with (1-16). Ignored for bare metal servers."},
				},
				{
					"--cpu-autoscale-policy-id",
					[]string{
						"ID of the vertical CPU Autoscale policy to associate the server with.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--memory-gb",
					[]string{
						"Number of GB of memory to configure the server with (1-128).",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--type",
					[]string{"Whether to create a standard, hyperscale, or bareMetal server."},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage.",
						"If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--anti-affinity-policy-id",
					[]string{
						"For standard servers, whether to use standard or premium storage. If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale. Ignored for bare metal servers.",
					},
				},
				{
					"--custom-fields",
					[]string{"Collection of custom field ID-value pairs to set for the server."},
				},
				{
					"--additional-disks",
					[]string{"Collection of disk parameters. Ignored for bare metal servers."},
				},
				{
					"--ttl",
					[]string{"Date/time that the server should be deleted. Ignored for bare metal servers."},
				},
				{
					"--packages",
					[]string{"Collection of packages to run on the server after it has been built. Ignored for bare metal servers."},
				},
				{
					"--configuration-id",
					[]string{
						"Specifies the identifier for the specific configuration type of bare metal server to deploy.",
						"Ignored for standard and hyperscale servers.",
					},
				},
				{
					"--os-type",
					[]string{
						"Specifies the OS to provision with the bare metal server. Currently, the only supported OS types",
						"are redHat6_64Bit, centOS6_64Bit, windows2012R2Standard_64Bit.",
						"Ignored for standard and hyperscale servers.",
					},
				},
			},
		},
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
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/startMaintenance",
		Resource: "server",
		Command:  "start-maintenance-mode",
		Help: help.Command{
			Brief: `Sends a start maintenance mode operation to a list of servers and adds operation to queue.`,
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to start maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/stopMaintenance",
		Resource: "server",
		Command:  "stop-maintenance-mode",
		Help: help.Command{
			Brief: `Sends a stop maintenance mode operation to a list of servers and adds operation to queue.`,
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"List of server IDs to stop maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.Import{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}",
		Resource: "server",
		Command:  "import",
		Help: help.Command{
			Brief: `Imports a new server from an uploaded OVF.`,
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server."},
				},
				{
					"--group-id",
					[]string{"ID of the parent group."},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--root-password",
					[]string{
						"Password of administrator or root user on server. This password must match",
						"the one set on the server being imported or the import will fail.",
					},
				},
				{
					"--cpu",
					[]string{
						"Number of processors to configure the server with (1-16). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--memoryGB",
					[]string{
						"Number of GB of memory to configure the server with (1-128). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--type",
					[]string{"Whether to create standard or hyperscale server"},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage. If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
					},
				},
				{
					"--custom-fields",
					[]string{"Collection of custom field ID-value pairs to set for the server."},
				},
				{
					"--ovf-id",
					[]string{"The identifier of the OVF that defines the server to import."},
				},
				{
					"--ovf-os-type",
					[]string{
						"The OS type of the server being imported. Currently, the only supported OS types",
						"are redHat6_64Bit, windows2008R2DataCenter_64bit, and windows2012R2DataCenter_64Bit.",
					},
				},
			},
		},
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
	registerCommandBase(&group.GetBillingReq{}, &group.GetBillingRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/billing",
		Resource: "group",
		Command:  "get-billing-details",
	})
	registerCommandBase(&group.GetStatsReq{}, &[]group.GetStatsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/statistics?start={Start}&end={End}&sampleInterval={SampleInterval}&type={Type}",
		Resource: "group",
		Command:  "get-monitoring-statistics",
	})
	registerCommandBase(&group.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "update",
	})
	registerCommandBase(&group.GetReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/archive",
		Resource: "group",
		Command:  "archive",
	})
	registerCommandBase(&group.RestoreReq{}, &group.RestoreRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/restore",
		Resource: "group",
		Command:  "restore",
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

	registerCommandBase(&network.ListReq{}, &[]network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}",
		Resource: "network",
		Command:  "list",
	})
	registerCommandBase(&network.GetReq{}, &network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}?ipAddresses={IpAddresses}",
		Resource: "network",
		Command:  "get",
	})
	registerCommandBase(&network.ListIpAddresses{}, &[]network.IpAddress{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}/ipAddresses?type={Type}",
		Resource: "network",
		Command:  "list-ip-addresses",
	})
	registerCommandBase(&network.CreateReq{}, &network.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/claim",
		Resource: "network",
		Command:  "create",
	})
	registerCommandBase(&network.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}",
		Resource: "network",
		Command:  "update",
	})
	registerCommandBase(&network.ReleaseReq{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{Network}/release",
		Resource: "network",
		Command:  "release",
	})

	registerCommandBase(&alert.CreateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "create",
	})
	registerCommandBase(nil, &alert.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "list",
	})
	registerCommandBase(&alert.GetReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "get",
	})
	registerCommandBase(&alert.UpdateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "update",
	})
	registerCommandBase(&alert.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "delete",
	})

	registerCommandBase(&affinity.CreateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "create",
	})
	registerCommandBase(nil, &affinity.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "list",
	})
	registerCommandBase(&affinity.GetReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "get",
	})
	registerCommandBase(&affinity.UpdateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "update",
	})
	registerCommandBase(&affinity.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "delete",
	})

	registerCommandBase(&firewall.CreateReq{}, &firewall.CreateRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}",
		Resource: "firewall-policy",
		Command:  "create",
	})
	registerCommandBase(&firewall.ListReq{}, &[]firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}?destinationAccount={DestinationAccountAlias}",
		Resource: "firewall-policy",
		Command:  "list",
	})
	registerCommandBase(&firewall.GetReq{}, &firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "get",
	})
	registerCommandBase(&firewall.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "update",
	})
	registerCommandBase(&firewall.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{SourceAccountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "delete",
	})

	registerCommandBase(&balancer.CreatePool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "create",
	})
	registerCommandBase(&balancer.Create{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "create",
	})
	registerCommandBase(&balancer.ListPools{}, &[]balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "list",
	})
	registerCommandBase(&balancer.List{}, &[]balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "list",
	})
	registerCommandBase(&balancer.GetPool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "get",
	})
	registerCommandBase(&balancer.Get{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "get",
	})
	registerCommandBase(&balancer.UpdatePool{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "update",
	})
	registerCommandBase(&balancer.Update{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "update",
	})
	registerCommandBase(&balancer.GetNodes{}, &[]balancer.Node{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "get-nodes",
	})
	registerCommandBase(&balancer.UpdateNodes{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "update-nodes",
	})
	registerCommandBase(&balancer.DeletePool{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "delete",
	})
	registerCommandBase(&balancer.Delete{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "delete",
	})

	registerCommandBase(nil, &[]customfields.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/accounts/{accountAlias}/customFields",
		Resource: "custom-fields",
		Command:  "get",
	})

	registerCustomCommand(commands.NewGroupList(commands.CommandExcInfo{
		Resource: "group",
		Command:  "list",
	}))
	registerCustomCommand(commands.NewServerList(commands.CommandExcInfo{
		Resource: "server",
		Command:  "list",
	}))
	registerCustomCommand(commands.NewWait(commands.CommandExcInfo{
		Resource: "wait",
	}))
	registerCustomCommand(commands.NewLogin(commands.CommandExcInfo{
		Resource: "login",
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
