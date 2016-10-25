package webhook

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Webhook struct {
	Name          string
	Configuration Configuration
	Links         []models.LinkEntity
}

type Configuration struct {
	Recursive  bool
	TargetUris []TargetUri
}

type TargetUri struct {
	TargetUri string
}
