package balancer

type UpdatePool struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
	PoolId         string `valid:"required" URIParam:"yes"`
	Method         string
	Persistence    string
}
