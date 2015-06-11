package base

type Options struct {
	Output   int
	Filter   []FilterOption
	Query    []string
	User     string
	Password string
	Profile  string
	Key      string
}

type FilterOption struct {
	PropertyName string
	Operation    int
	Value        string
}
