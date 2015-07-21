package datacenter

type ListReq struct{}

type ListRes struct {
	Id    string
	Name  string
	Links []map[string]interface{}
}
