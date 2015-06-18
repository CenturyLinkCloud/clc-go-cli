package base

type Connection interface {
	ExecuteRequest(verb string, url string, reqModel interface{}, resModel interface{}) (err error)
}
