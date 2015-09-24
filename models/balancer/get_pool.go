package balancer

type GetPool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
	PoolId       string `valid:"required" URIParam:"yes"`
}
