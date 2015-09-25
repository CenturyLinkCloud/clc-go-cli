package autoscale

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id                     string
	Name                   string
	ResourceType           string
	ThresholdPeriodMinutes int64
	ScaleUpIncrement       int64
	Range                  Range
	ScaleUpThreshold       int64
	ScaleDownThreshold     int64
	ScaleDownWindow        ScaleDownWindow
	Links                  []models.LinkEntity
}

type Range struct {
	Min int
	Max int
}

type ScaleDownWindow struct {
	Start string
	End   string
}
