package group

import "fmt"

type SetHAPolicy struct {
	Group        `URIParam:"GroupId" argument:"composed" json:"-"`
	PolicyId     string         `valid:"required"`
	LoadBalancer LBForHARequest `json:"LoadBalancerPool"`
}

type LBForHARequest struct {
	Id          string
	PublicPort  int64
	PrivatePort int64
}

func (s *SetHAPolicy) Validate() error {
	err := s.Group.Validate()
	if err != nil {
		return err
	}

	lb := s.LoadBalancer
	if lb.Id == "" || lb.PublicPort == 0 || lb.PrivatePort == 0 {
		return fmt.Errorf("LoadBalancer: Id, PublicPort, and PrivatePort are required")
	}
	return nil
}
