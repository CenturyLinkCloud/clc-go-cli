package group

import (
	"fmt"
	"regexp"
	"time"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/errors"
)

var (
	WeekDays = []string{
		"sun",
		"mon",
		"tue",
		"wed",
		"thu",
		"fri",
		"sat",
	}
)

type SetScheduledActivities struct {
	Group               `URIParam:"GroupId" argument:"composed" json:"-"`
	Status              string    `oneOf:"on,off" valid:"required"`
	Type                string    `oneOf:"archive,createsnapshot,delete,deletesnapshot,pause,poweron,reboot,shutdown" valid:"required"`
	BeginDateUtc        time.Time `json:"beginDateUTC"`
	Repeat              string    `oneOf:"never,daily,weekly,monthly,customWeekly" valid:"required"`
	CustomWeeklyDays    []string
	Expire              string    `oneOf:"never,afterDate,afterCount" valid:"required"`
	ExpireCount         int64     `json:",omitempty"`
	ExpireDateUtc       time.Time `json:"-"`
	ExpireDateUTCString string    `json:"expireDateUTC,omitempty" argument:"ignore"`
	TimeZoneOffset      string    `valid:"required"`
}

func (s *SetScheduledActivities) Validate() error {
	zeroTime := time.Time{}

	if err := s.Group.Validate(); err != nil {
		return err
	}

	if s.BeginDateUtc == zeroTime {
		return errors.EmptyField("begin-date-utc")
	}

	if s.Repeat == "customWeekly" {
		if len(s.CustomWeeklyDays) == 0 {
			return fmt.Errorf("When repeat is customWeekly the custom-weekly-days field must be set and non-empty")
		}
		for _, day := range s.CustomWeeklyDays {
			if !IsWeekDay(day) {
				return fmt.Errorf("Every custom week day must be one of sun, mon, tue, wed, thu, fri, or sat")
			}
		}
	}

	if s.Expire == "afterDate" && s.ExpireDateUtc == zeroTime {
		return fmt.Errorf("When expire is afterDate the expire-date-utc field must be set and non-empty")
	} else if s.Expire == "afterCount" && s.ExpireCount == 0 {
		return fmt.Errorf("When expire is afterCount the expire-count field must be set and non-empty")
	}

	matched, err := regexp.MatchString(`^(-)?[[:digit:]]{2}:[[:digit:]]{2}$`, s.TimeZoneOffset)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("The time-zone-offset value must be in the format (-)hh:mm")
	}

	return nil
}

func (s *SetScheduledActivities) ApplyDefaultBehaviour() error {
	zeroTime := time.Time{}
	if s.ExpireDateUtc != zeroTime {
		s.ExpireDateUTCString = s.ExpireDateUtc.Format(base.SERVER_TIME_FORMAT)
	}
	return nil
}

func IsWeekDay(day string) bool {
	for _, option := range WeekDays {
		if option == day {
			return true
		}
	}
	return false
}
