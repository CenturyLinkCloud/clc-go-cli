package server

type RevertToSnapshotReq struct {
	Server     `argument:"composed" URIParam:"ServerId"`
	SnapshotId string `valid:"required" URIParam:"yes"`
}
