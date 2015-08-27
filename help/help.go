package help

import (
	"bytes"
	"text/template"
)

type Resource struct {
	Name     string
	Commands []Argument
}

type Command struct {
	Brief     []string
	Arguments []Argument
}

type Argument struct {
	Name        string
	Description []string
}

var resourceHelpTemplate = `Available {{.Name}} commands:
{{range .Commands}}
	{{.Name}}
{{range .Description}}{{ printf "\t\t" }}{{ . }}{{ printf "\n" }}{{end}}
{{end}}`

var commandHelpTemplate = `{{range .Brief}}{{ . }}{{ printf " " }}{{ end }}

PARAMETERS:
{{range .Arguments}}
	{{.Name}}
{{range .Description}} {{ printf "\t\t" }} {{ . }} {{ printf "\n" }} {{end}}
{{end}}
OPTIONS:

	--help
			Shows general help or help for the given resource or command.
	--user
			Specifies a user name for the account.
	--password
			Specifies the password for the given user.
	--profile
			Specifies a profile to use (one from the config file).
	--output
			Specifies the output format - either 'json', 'text' or 'table'.
			Defaults to 'json'.
	--generate-cli-skeleton
			If specified, the command is not actually executed. Instead, JSON with all of
			its arguments is returned. All the arguments and options specified along with
			this option are included in the output.
	--from-file
			Specifies a JSON file to load command arguments and options from. No other
			arguments or options can be specified.
	--query
			Restricts the fields in the result to only those specified with this option.

			Multiple fields can be separated by comma. Nested fields can be queried using dot.
			Multiple nested fields are also separated by comma and must be enclosed in curly
			braces on the deepest level of nesting. Note, that a shell may treat curly braces
			in a special way so put the whole query in quotes to avoid errors.
			Aliases may be set for the nested fields using semicolons inside the braces.

			An example:
				clc server list --query "details.IP-addresses.{I:internal,P:public}"
	--filter
			Filters out the returned entities that do not match the given conditions. Multiple
			conditions are separated via comma. Each condition consists of a field, an
			operation and a value. Supported operations are:

			=	equals, applicable to strings, numbers and booleans
			^= (starts with), $= (ends with), ~= (contains)	these three are only for strings
			<,<=,>,>= 	comparison operators, can be used with numbers and strings
	--trace
			If specified, prints out all the HTTP request/response data.

ENVIRONMENT VARIABLES:

	CLC_USER	Specifies a user name for the account.
	CLC_PASSWORD	Specifies the password for the given user.
	CLC_PROFILE	Specifies a profile to use (one from the config file).
	CLC_TRACE	If specified (any non-empty value fits), prints out all the HTTP request/response data.
`

func ForCommand(cmd Command) string {
	tmpl, err := template.New("command help").Parse(commandHelpTemplate)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, cmd)
	if err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}

func ForResource(r Resource) string {
	tmpl, err := template.New("resource help").Parse(resourceHelpTemplate)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, r)
	if err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}
