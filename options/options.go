package options

type Options struct {
	Output   string
	Filter   string
	Query    string
	User     string
	Password string
	Profile  string
	Trace    bool
	Help     bool
}

func Get() []string {
	return []string{
		"--help",
		"--from-file",
		"--user",
		"--password",
		"--profile",
		"--trace",
		"--query",
		"--filter",
		"--output",
		"--generate-cli-skeleton",
	}
}
