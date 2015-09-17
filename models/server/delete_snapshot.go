package server

type DeleteSnapshotReq struct {
	Server     `argument:"composed" URIParam:"ServerId"`
	SnapshotId string `valid:"required" URIParam:"yes"`
}
