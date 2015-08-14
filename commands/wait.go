package commands

import (
	"encoding/json"
	"fmt"
	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/state"
	"time"
)

type Wait struct {
	CommandBase
}

type ComplexResponse struct {
	Links []models.LinkEntity
}

type StatusResponse struct {
	Status string
}

func NewWait(info CommandExcInfo) *Wait {
	w := Wait{}
	w.ExcInfo = info
	return &w
}

func (w *Wait) Execute(cn base.Connection) error {
	w.Output = "Nothing to wait for."
	bytes, err := state.LoadLastResult()
	if err != nil {
		return nil
	}

	var links []models.LinkEntity
	c := ComplexResponse{}
	l := models.LinkEntity{}
	status := models.Status{}

	if err = json.Unmarshal(bytes, &l); err == nil && l.Href != "" {
		links = []models.LinkEntity{l}
	} else if err = json.Unmarshal(bytes, &c); err == nil && len(c.Links) > 0 {
		links = c.Links
	} else {
		json.Unmarshal(bytes, &status)
	}

	if len(links) > 0 {
		for _, link := range links {
			if link.Rel == "status" {
				w.Output = ping(cn, fmt.Sprintf("%s%s", BaseURL, link.Href))
				return nil
			}
		}
	} else if status.URI != "" {
		w.Output = ping(cn, fmt.Sprintf("%s%s", BaseURL, status.URI))
	}
	return nil
}

func (w *Wait) InputModel() interface{} {
	return &inputStub{}
}

func ping(cn base.Connection, URL string) (status StatusResponse) {
	status = StatusResponse{Status: "notStarted"}
	for status.Status == "executing" || status.Status == "resumed" || status.Status == "notStarted" {
		cn.ExecuteRequest("GET", URL, nil, &status)
		time.Sleep(200)
	}
	return
}
