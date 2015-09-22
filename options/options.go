package options

type Options struct {
	Output       string
	Filter       string
	Query        string
	User         string
	Password     string
	Profile      string
	AccountAlias string
	Trace        bool
	Help         bool
}

func Get() []string {
	return []string{
		"--help",
		"--from-file",
		"--user",
		"--password",
		"--profile",
		"--account-alias",
		"--trace",
		"--query",
		"--filter",
		"--output",
		"--generate-cli-skeleton",
	}
}
