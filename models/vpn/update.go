package vpn

type UpdateReq struct {
	VpnId string `json:"-" valid:"required" URIParam:"yes"`

	Local  LocalPropertiesUpdateReq
	Remote RemotePropertiesUpdateReq
	Ipsec  IpSecUpdateReq
	Ike    IkeUpdateReq
}

type LocalPropertiesUpdateReq struct {
	Subnets []string
}

type RemotePropertiesUpdateReq struct {
	SiteName   string
	DeviceType string
	Address    string
	Subnets    []string
}

type IpSecUpdateReq struct {
	Encryption string `oneOf:"aes128,aes192,aes256,tripleDES,optional"`
	Hashing    string `oneOf:"sha1_96,sha1_256,md5,optional"`
	Protocol   string `oneOf:"esp,ah,optional"`
	Pfs        string `oneOf:"disabled,group1,group2,group5,optional"`
	Lifetime   string `oneOf:"3600,28800,86400,optional"`
}

type IkeUpdateReq struct {
	Encryption         string `oneOf:"aes128,aes192,aes256,tripleDES"`
	Hashing            string `oneOf:"sha1_96,sha1_256,md5,optional"`
	DiffieHellmanGroup string `oneOf:"group1,group2,group5"`
	PreSharedKey       string
	Lifetime           string `oneOf:"3600,28800,86400"`
	Mode               string `oneOf:"main,aggresive"`
	DeadPeerDetection  string `oneOf:"true,false,optional"`
	NatTraversal       string `oneOf:"true,false,optional"`
	RemoteIdentity     string
}
