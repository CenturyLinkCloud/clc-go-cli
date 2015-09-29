package group

import "github.com/centurylinkcloud/clc-go-cli/base"

type GetBillingReq struct {
	Group `argument:"composed" URIParam:"GroupId" json:"-"`
}

type GetBillingRes struct {
	Date   base.Time
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
