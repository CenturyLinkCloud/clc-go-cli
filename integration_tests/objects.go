package main

type ApiDef struct {
	Method            string
	Url               string
	UrlExample        string
	UrlParameters     []ParameterDef
	ContentExample    string
	ContentParameters []ParameterDef
	ResExample        string
	ResProperties     []ResPropertyDef
}

type ParameterDef struct {
	Name        string
	Type        string
	Description string
	IsRequired  bool
	Children    map[string]ParameterDef
}

type ResPropertyDef struct {
	Name        string
	Type        string
	Description string
	Children    map[string]ResPropertyDef
}
