package balancer

type CreatePool struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
	Port           int64  `valid:"required"`
	Method         string
	Persistence    string
}
