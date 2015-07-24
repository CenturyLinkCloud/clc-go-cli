package server

type RevertToSnapshotReq struct {
	ServerId   string `valid:"required" URIParam:"yes"`
	SnapshotId string `valid:"required" URIParam:"yes"`
}
