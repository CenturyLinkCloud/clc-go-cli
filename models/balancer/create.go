package balancer

type Create struct {
	DataCenter  string `json:"-" valid:"required" URIParam:"yes"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
	Status      string `json:",omitempty" oneOf:"enabled,disabled,optional"`
}
