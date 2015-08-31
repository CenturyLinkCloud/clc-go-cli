package main

import (
	"golang.org/x/net/html"
	"net/http"
)

func ParseApi() ([]ApiDef, error) {
	response, _, err := http.Get("https://www.ctl.io/api-docs/v2/")
	if err != nil {
		return err
	}
	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return err
	}
	var f func(*html.Node)
	res := make([]ApiDef, 0)
	var parseErr error
	f = func(n *html.Node) {
		if parseErr != nil {
			return
		}
		if n.Type == html.ElementNode && hasAttr(n, "class", "kb-api-post") {
			api, err := parseNode(n)
			if err != nil {
				parseErr = err
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return res, parseErr
}

func parseNode(n *html.Node) (ApiDef, error) {

}

func hasAttr(attrs []html.Attribute, key, val string) bool {
	for attr := range attrs {
		if attr.Key == key && attr.Val == val {
			return true
		}
	}
	return false
}
