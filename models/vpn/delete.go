package vpn

type DeleteReq struct {
	VpnId string `valid:"required" URIParam:"yes"`
}
