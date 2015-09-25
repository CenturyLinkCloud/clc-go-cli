package autoscale

import (
	"github.com/centurylinkcloud/clc-go-cli/models"
)

type Entity struct {
	Id                     string
	Name                   string
	ResourceType           string
	ThresholdPeriodMinutes int
	ScaleUpIncrement       int
	Range                  Range
	ScaleUpThreshold       int
	scaleDownThreshold     int
	scaleDownWindow        ScaleDownWindow
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
