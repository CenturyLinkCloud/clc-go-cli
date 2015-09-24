package balancer

type Update struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
	Name         string `valid:"required"`
	Description  string `valid:"required"`
	Status       string `json:",omitempty" oneOf:"enabled,disabled"`
}
