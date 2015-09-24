package balancer

type Get struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter" json:"-"`
}
