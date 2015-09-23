package alert

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Policy struct {
	PolicyId   string `json:"-"`
	PolicyName string `json:"-"`
}

var (
	policiesURL = "https://api.ctl.io/v2/alertPolicies/{accountAlias}"
)

func (p *Policy) Validate() error {
	if (p.PolicyId == "") == (p.PolicyName == "") {
		return fmt.Errorf("Exactly one of the policy-id and policy-name parameters must be specified")
	}
	return nil
}

func (p *Policy) InferID(cn base.Connection) error {
	if p.PolicyName == "" {
		return nil
	}

	policies := &ListRes{}
	err := cn.ExecuteRequest("GET", policiesURL, nil, policies)
	if err != nil {
		return err
	}

	matched := []string{}
	for _, pl := range policies.Items {
		if p.PolicyName == pl.Name {
			matched = append(matched, pl.Id)
		}
	}

	switch len(matched) {
	case 0:
		return fmt.Errorf("There are no policies with name '%s'", p.PolicyName)
	case 1:
		p.PolicyId = matched[0]
		return nil
	default:
		return fmt.Errorf("There are more than one policy with name '%s'. Please, specify an ID.", p.PolicyName)
	}
}

func (p *Policy) GetNames(cn base.Connection, property string) ([]string, error) {
	if property != "PolicyName" {
		return nil, nil
	}

	policies := &ListRes{}
	err := cn.ExecuteRequest("GET", policiesURL, nil, policies)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, policy := range policies.Items {
		names = append(names, policy.Name)
	}
	return names, nil
}
