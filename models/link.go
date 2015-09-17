package models

type LinkEntity struct {
	Rel   string
	Href  string
	Id    string   `json:",omitempty"`
	Name  string   `json:",omitempty"`
	Verbs []string `json:",omitempty"`
}
