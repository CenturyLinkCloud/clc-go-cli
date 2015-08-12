package help

import (
	"bytes"
	"text/template"
)

type Command struct {
	Brief     string
	Arguments []Argument
}

type Argument struct {
	Name        string
	Description []string
}

var commandHelpTemplate = `{{.Brief}}

Parameters:
{{range .Arguments}}
	{{.Name}}
{{range .Description}} {{ printf "\t\t" }} {{ . }} {{ printf "\n" }} {{end}}
{{end}}`

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
