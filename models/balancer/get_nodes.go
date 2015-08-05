package balancer

type GetNodes struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
	PoolId         string `valid:"required" URIParam:"yes"`
}
