package affinity

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
)

type Policy struct {
	PolicyId   string `URIParam:"yes"`
	PolicyName string
}

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

	url := "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}"
	policies := &ListRes{}
	err := cn.ExecuteRequest("GET", url, nil, policies)
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
