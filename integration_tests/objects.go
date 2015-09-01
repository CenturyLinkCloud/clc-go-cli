package main

type ApiDef struct {
	Method            string
	Url               string
	UrlExample        string
	UrlParameters     []*ParameterDef
	ContentParameters []*ParameterDef
	ContentExample    string
	ResExample        string
	ResParameters     []*ParameterDef
}

type ParameterDef struct {
	Name        string
	Type        string
	Description string
	IsRequired  bool
	Children    []*ParameterDef
}
