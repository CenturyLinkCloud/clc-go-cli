package balancer

type Update struct {
	DataCenter     string `valid:"required" URIParam:"yes"`
	LoadBalancerId string `valid:"required" URIParam:"yes"`
	Name           string `valid:"required"`
	Description    string `valid:"required"`
	Status         string
}
