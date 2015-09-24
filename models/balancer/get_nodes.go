package balancer

type GetNodes struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
	PoolId       string `valid:"required" URIParam:"yes"`
}
