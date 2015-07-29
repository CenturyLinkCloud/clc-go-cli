package customfields

type GetRes struct {
	Id         string
	Name       string
	IsRequired bool
	Type       string
	Options    []NameValue
}

type NameValue struct {
	Name   string
	String string
}
