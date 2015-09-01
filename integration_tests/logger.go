package main

import (
	"fmt"
	"golang.org/x/net/html"
)

type Logger interface {
	Log(format string, a ...interface{})
	LogNode(message string, n *html.Node)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) Log(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

func (l *logger) LogNode(message string, n *html.Node) {
	if n == nil {
		fmt.Printf("%s: <nil> \n", message)
	} else {
		fmt.Printf("%s: Data = %s, Attrs = %+v\n", message, n.Data, n.Attr)
	}
}
