package models

type PortRestriction struct {
	Protocol string
	Port     int64
	PortTo   int64 `json:",omitempty"`
}

type SourceRestriction struct {
	CIDR string
}
