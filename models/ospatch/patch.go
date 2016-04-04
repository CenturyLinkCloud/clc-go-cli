package ospatch

import (
	"encoding/json"

	"github.com/centurylinkcloud/clc-go-cli/errors"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
)

var (
	packages = map[string]server.PackageDef{
		"Windows2012": server.PackageDef{
			PackageId: "b229535c-a313-4a31-baf8-6aa71ff4b9ed",
		},
		"RedHat": server.PackageDef{
			PackageId: "c3c6642e-24e1-4c37-b56a-1cf1476ee360",
			Parameters: map[string]string{
				"patch.debug.mode": "false",
			},
		},
	}
)

type Patch struct {
	ServerIds []string `json:"servers"`
	OsType    string   `valid:"required" oneOf:"Windows2012,RedHat"`
}

func (p *Patch) Validate() error {
	if len(p.ServerIds) == 0 {
		return errors.EmptyField("server-ids")
	}
	return nil
}

func (p *Patch) MarshalJSON() ([]byte, error) {
	return json.Marshal(server.ExecutePackage{
		ServerIds: p.ServerIds,
		Package:   packages[p.OsType],
	})
}

type List struct {
	server.Server `argument:"composed" URIParam:"ServerId"`
}
