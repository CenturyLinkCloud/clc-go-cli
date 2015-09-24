package models

import (
	"time"
)

type ChangeInfo struct {
	CreatedDate  time.Time
	CreatedBy    string
	ModifiedDate time.Time
	ModifiedBy   string
}
