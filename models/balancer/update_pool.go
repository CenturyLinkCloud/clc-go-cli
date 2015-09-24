package balancer

type UpdatePool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
	PoolId       string `json:"-" valid:"required" URIParam:"yes"`
	Method       string `oneOf:"roundRobin,leastConnection"`
	Persistence  string `oneOf:"standard,sticky"`
}
