package balancer

type ListPools struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter"`
}
