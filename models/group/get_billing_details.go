package group

import (
	"time"
)

type GetBillingReq struct {
	GroupId string `valid:"required" URIParam:"yes"`
}

type GetBillingRes struct {
	Date   time.Time
	Groups map[string]GroupBilling
}

type GroupBilling struct {
	Name    string
	Servers map[string]ServerBilling
}

type ServerBilling struct {
	TemplateCost    float64
	ArchiveCost     float64
	MonthlyEstimate float64
	MonthToDate     float64
	CurrentHour     float64
}
