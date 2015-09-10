package balancer

type UpdatePool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter"`
	PoolId       string `valid:"required" URIParam:"yes"`
	Method       string `oneOf:"roundRobin,leastConnection"`
	Persistence  string `oneOf:"standard,sticky"`
}
