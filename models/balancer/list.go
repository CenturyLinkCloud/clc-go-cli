package balancer

type List struct {
	DataCenter string `valid:"required" URIParam:"yes"`
}
