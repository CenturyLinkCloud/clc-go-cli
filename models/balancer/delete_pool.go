package balancer

type DeletePool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
	PoolId       string `valid:"required" URIParam:"yes"`
}
