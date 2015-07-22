package datacenter

import (
	"fmt"
)

type GetReq struct {
	DataCenter string `valid:"required" URIParam:"yes"`
	GroupLinks string `valid:"required" URIParam:"yes"`
}

type GetRes struct {
	Id    string
	Name  string
	Links []map[string]interface{}
}

func (r *GetReq) Validate() error {
	if r.GroupLinks != "false" && r.GroupLinks != "true" {
		return fmt.Errorf("group-links value must be either true or false.")
	}
	return nil
}
