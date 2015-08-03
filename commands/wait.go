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
	complex := ComplexResponse{}
	err = json.Unmarshal(bytes, &complex)
	if err == nil && len(complex.Links) != 0 {
		links = complex.Links
	} else {
		flat := models.LinkEntity{}
		err = json.Unmarshal(bytes, &flat)
		if err != nil {
			return nil
		}
		links = []models.LinkEntity{flat}
	}
	for _, link := range links {
		if link.Rel == "status" {
			sr := StatusResponse{Status: "notStarted"}
			statusURL := fmt.Sprintf("%s%s", BaseURL, link.Href)
			for sr.Status == "executing" || sr.Status == "resumed" || sr.Status == "notStarted" {
				cn.ExecuteRequest("GET", statusURL, nil, &sr)
				time.Sleep(200)
			}
			w.Output = sr
			return nil
		}
	}
	return nil
}

func (w *Wait) InputModel() interface{} {
	return &inputStub{}
}
