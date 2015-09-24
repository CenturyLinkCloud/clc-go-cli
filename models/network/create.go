package network

type CreateReq struct {
	DataCenter string `json:"-" valid:"required" URIParam:"yes"`
}
