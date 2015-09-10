package balancer

type DeletePool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter"`
	PoolId       string `valid:"required" URIParam:"yes"`
}
