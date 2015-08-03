package options

type Options struct {
	Output   string
	Filter   []FilterOption
	Query    string
	User     string
	Password string
	Profile  string
	Trace    bool
	Help     bool
}

type FilterOption struct {
	PropertyName string
	Operation    string
	Value        string
}
