package network

import (
	"fmt"
)

type ListIpAddresses struct {
	DataCenter string `valid:"required" URIParam:"yes"`
	Network    string `valid:"required" URIParam:"yes"`
	Type       string `URIParam:"yes"`
}

func (l *ListIpAddresses) Validate() error {
	if l.Type == "" {
		return nil
	}
	t := l.Type
	if t != "claimed" && t != "free" && t != "all" {
		return fmt.Errorf("type value must be one of the claimed, free, all.")
	}
	return nil
}
