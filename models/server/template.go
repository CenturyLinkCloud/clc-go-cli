package server

import (
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"time"
)

var (
	TemplateTimeout = 200
)

func LoadTemplates(cn base.Connection) ([]string, error) {
	datacenters := []datacenter.ListRes{}
	URL := fmt.Sprintf("%s/v2/datacenters/{accountAlias}", base.URL)
	err := cn.ExecuteRequest("GET", URL, nil, &datacenters)
	if err != nil {
		return nil, err
	}

	done := make(chan error)
	capabilities := make([]datacenter.GetDCRes, len(datacenters))
	for i, ref := range datacenters {
		go loadCapabilities(ref, capabilities, i, cn, done)
	}

	count := 0
ConsumingCapabilities:
	for {
		select {
		case err := <-done:
			if err != nil {
				return nil, err
			}
			count += 1
			if count == len(datacenters) {
				break ConsumingCapabilities
			}
		case <-time.After(time.Second * time.Duration(TemplateTimeout)):
			return nil, fmt.Errorf("Request timeout error.")
		}
	}

	templates := map[string]bool{}
	for _, c := range capabilities {
		for _, t := range c.Templates {
			templates[t.Name] = true
		}
	}
	keys := []string{}
	for k := range templates {
		keys = append(keys, k)
	}
	return keys, nil
}

func loadCapabilities(ref datacenter.ListRes, capabilities []datacenter.GetDCRes, dcnum int, cn base.Connection, done chan<- error) {
	link, err := models.GetLink(ref.Links, "deploymentCapabilities")
	if err != nil {
		done <- err
	}

	d := datacenter.GetDCRes{}
	URL := fmt.Sprintf("%s%s?groupLinks=false", base.URL, link)
	err = cn.ExecuteRequest("GET", URL, nil, &d)
	if err != nil {
		done <- err
	}

	capabilities[dcnum] = d
	done <- nil
}
