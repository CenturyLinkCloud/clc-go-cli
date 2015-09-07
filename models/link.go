package models

import (
	"fmt"
)

type LinkEntity struct {
	Rel   string
	Href  string
	Id    string
	Name  string
	Verbs []string
}

func GetLink(links []LinkEntity, resource string) (string, error) {
	for _, link := range links {
		if link.Rel == resource {
			return link.Href, nil
		}
	}
	return "", fmt.Errorf("No %s link found", resource)
}
