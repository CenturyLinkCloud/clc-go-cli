package server

type RevertToSnapshotReq struct {
	Server     `json:"-" argument:"composed" URIParam:"ServerId"`
	SnapshotId string `json:"-" valid:"required" URIParam:"yes"`
}
