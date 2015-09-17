package models

import (
	"time"
)

type ChangeInfo struct {
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedData time.Time
	ModifiedBy   string
}
