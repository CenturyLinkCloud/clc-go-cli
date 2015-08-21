package balancer

type Create struct {
	DataCenter  string `valid:"required" URIParam:"yes"`
	Name        string `valid:"required"`
	Description string `valid:"required"`
	Status      string `oneOf:"enabled,disabled"`
}
