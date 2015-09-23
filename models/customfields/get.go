package customfields

type GetRes struct {
	Id         string
	Name       string
	IsRequired bool
	Type       string `oneOf:"text,checkbox,option"`
	Options    []NameValue
}

type NameValue struct {
	Name  string
	Value string
}
