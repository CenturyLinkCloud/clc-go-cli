package server

type DeleteSnapshotReq struct {
	Server     `argument:"composed" URIParam:"ServerId" json:"-"`
	SnapshotId string `valid:"required" URIParam:"yes" json:"-"`
}
