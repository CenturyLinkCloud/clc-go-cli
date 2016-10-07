package integration_tests

type ApiDef struct {
	Method            string
	Url               string
	UrlExample        string
	UrlParameters     []*ParameterDef
	ContentParameters []*ParameterDef
	ContentExample    interface{}
	ResExample        interface{}
	ResParameters     []*ParameterDef
}

type ParameterDef struct {
	Name        string
	Type        string
	Description string
	IsRequired  bool
	Children    []*ParameterDef
}
