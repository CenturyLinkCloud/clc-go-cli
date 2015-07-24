package server

type GetImportsReq struct {
	LocationId string `valid:"required" URIParam:"yes"`
}

type GetImportsRes []AvailableImport

type AvailableImport struct {
	Id            string
	Name          string
	StorageSizeGB int64
	CpuCount      int64
	MemorySizeMB  int64
}
