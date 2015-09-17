package group

import (
	"fmt"
	"regexp"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

type GetStatsReq struct {
	Group          `argument:"composed" URIParam:"GroupId"`
	Start          string `URIParam:"yes"`
	End            string `URIParam:"yes"`
	SampleInterval string `URIParam:"yes"`
	Type           string `URIParam:"yes" oneOf:"latest,hourly,realtime"`
}

type GetStatsRes struct {
	Name  string
	Stats []Stats
}

type Stats struct {
	Timestamp                time.Time
	Cpu                      float64
	CpuPercent               float64
	MemoryMB                 float64
	MemoryPercent            float64
	NetworkReceivedKbs       float64
	NetworkTransmittedKbs    float64
	DiskUsageTotalCapacityMB float64
	DiskUsage                []DiskUsage
	GuestDiskUsage           []GuestDiskUsage
}

type DiskUsage struct {
	Id         string
	CapacityMB int64
}

type GuestDiskUsage struct {
	Path       string
	CapacityMB int64
	ConsumedMB int64
}

func (g *GetStatsReq) Validate() error {
	if g.Type == "latest" {
		return nil
	}
	if g.Start == "" || g.SampleInterval == "" {
		return fmt.Errorf("For the types hourly and realtime both start and sample-interval must be set.")
	}
	_, err := time.Parse(timeFormat, g.Start)
	if err != nil {
		return fmt.Errorf("start must be in YYYY-MM-DD hh:mm:ss format.")
	}
	if g.End != "" {
		_, err := time.Parse(timeFormat, g.End)
		if err != nil {
			return fmt.Errorf("end must be in YYYY-MM-DD hh:mm:ss format.")
		}
	}
	match, err := regexp.Match("^[[:digit:]]{2}:[[:digit:]]{2}:[[:digit:]]{2}$", []byte(g.SampleInterval))
	if err != nil || !match {
		return fmt.Errorf("sample-interval must be in 02:00:00 format.")
	}
	return nil
}
