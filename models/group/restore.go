package group

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type RestoreReq struct {
	Group           `argument:"composed" URIParam:"GroupId"`
	TargetGroupId   string
	TargetGroupName string `json:"-"`
}

type RestoreRes struct {
	IsQueued     bool
	Links        []models.LinkEntity `json:",omitempty"`
	ErrorMessage string              `json:",omitempty"`
}

func (r *RestoreReq) Validate() error {
	if (r.TargetGroupId == "") == (r.TargetGroupName == "") {
		return fmt.Errorf("Exactly one of the target-group-id and taret-group-name parameters must be specified")
	}
	return nil
}

func (r *RestoreReq) InferID(cn base.Connection) error {
	err := r.Group.InferID(cn)
	if err != nil {
		return err
	}

	if r.TargetGroupName == "" {
		return nil
	}

	id, err := IDByName(cn, "all", r.TargetGroupName)
	if err != nil {
		return err
	}
	r.TargetGroupId = id
	return nil
}

func (r *RestoreReq) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "TargetGroupName" && property != "GroupName" {
		return nil, nil
	}

	// Group and Target Group need the same list of names.
	return GetNames(cn, "all")
}
