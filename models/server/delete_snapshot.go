package server

type DeleteSnapshotReq struct {
	ServerId   string `valid:"required" URIParam:"yes"`
	SnapshotId string `valid:"required" URIParam:"yes"`
}
