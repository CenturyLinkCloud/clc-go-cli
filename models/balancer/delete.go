package balancer

type Delete struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter"`
}
