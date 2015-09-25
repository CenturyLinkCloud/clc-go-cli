package autoscale

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Policy struct {
	PolicyId   string
	PolicyName string
}

var (
	policiesURL = "https://api.ctl.io/v2/autoscalePolicies/{accountAlias}"
)

func validatePolicy(policyId, policyName string) error {
	if (policyId == "") == (policyName == "") {
		return fmt.Errorf("Exactly one of the policy-id and policy-name parameters must be specified")
	}
	return nil
}

func inferID(policyName string, cn base.Connection) (string, error) {
	if policyName == "" {
		return "", nil
	}

	policies := []Entity{}
	err := cn.ExecuteRequest("GET", policiesURL, nil, &policies)
	if err != nil {
		return "", err
	}

	matched := []string{}
	for _, pl := range policies {
		if policyName == pl.Name {
			matched = append(matched, pl.Id)
		}
	}

	switch len(matched) {
	case 0:
		return "", fmt.Errorf("There are no policies with name '%s'", policyName)
	case 1:
		return matched[0], nil
	default:
		return "", fmt.Errorf("There are more than one policy with name '%s'. Please, specify an ID.", policyName)
	}
}

func getNames(cn base.Connection, property string) ([]string, error) {
	if property != "PolicyName" {
		return nil, nil
	}

	policies := []Entity{}
	err := cn.ExecuteRequest("GET", policiesURL, nil, &policies)
	if err != nil {
		return nil, err
	}

	names := []string{}
	for _, policy := range policies {
		names = append(names, policy.Name)
	}
	return names, nil
}
func (p *Policy) Validate() error {
	return validatePolicy(p.PolicyId, p.PolicyName)
}

func (p *Policy) InferID(cn base.Connection) error {
	id, err := inferID(p.PolicyName, cn)
	if err == nil && id != "" {
		p.PolicyId = id
	}
	return err
}

func (p *Policy) GetNames(cn base.Connection, property string) ([]string, error) {
	return getNames(cn, property)
}
