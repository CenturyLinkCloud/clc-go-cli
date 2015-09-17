package server

type RevertToSnapshotReq struct {
	ServerId   string `json:"-" valid:"required" URIParam:"yes"`
	SnapshotId string `json:"-" valid:"required" URIParam:"yes"`
}
