package cli

type Options struct {
	Output   int
	Filter   []FilterOption
	Query    []string
	User     string
	Password string
	Profile  string
}

type FilterOption struct {
	PropertyName string
	Operation    int
	Value        string
}
