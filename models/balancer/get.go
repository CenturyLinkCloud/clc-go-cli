package balancer

type Get struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
}
