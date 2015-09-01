package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"strings"
)

type CheckRes int

const (
	CheckResReturn   CheckRes = iota
	CheckResContinue CheckRes = iota
	CheckResApply    CheckRes = iota
)

type Parser interface {
	ParseApi() ([]*ApiDef, error)
}

type parser struct {
	logger Logger
}

func NewParser(logger Logger) Parser {
	return &parser{logger: logger}
}

func (p *parser) ParseApi() ([]*ApiDef, error) {
	response, err := http.Get("https://www.ctl.io/api-docs/v2/")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	res := make([]*ApiDef, 0)
	var parseErr error
	firstApiSectionId := "alert-policies-create-alert-policy"
	firstSectionReached := false
	p.findNode(doc,
		func(n *html.Node) CheckRes {
			if parseErr != nil {
				return CheckResReturn
			}
			if p.hasAttr(n.Attr, "id", firstApiSectionId) {
				firstSectionReached = true
			}
			if p.hasAttr(n.Attr, "class", "kb-api-post") && firstSectionReached {
				return CheckResApply
			}
			return CheckResContinue
		},
		func(n *html.Node) {
			api, err := p.parseNode(n)
			if err != nil {
				parseErr = err
			} else {
				res = append(res, api)
			}
		})

	return res, parseErr
}

func (p *parser) parseNode(n *html.Node) (*ApiDef, error) {
	p.logger.LogNode("Parsing node", n)
	var apiSection *html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if p.hasAttr(c.Attr, "data-api-section", "") {
			apiSection = c
			break
		}
	}
	if apiSection == nil {
		return nil, fmt.Errorf("'data-api-section' attribute missed.")
	}

	parseChildren := func(parameters []*ParameterDef) {
		for _, param := range parameters {
			if param.Type == "complex" {
				_, param.Children = p.parseTable(apiSection, strings.Title(param.Name+" Definition"))
			}
		}
	}

	res := &ApiDef{}
	apiSection, res.Method, res.Url = p.parseUrl(apiSection, "Structure")
	apiSection, _, res.UrlExample = p.parseUrl(apiSection, "Example")
	apiSection, res.UrlParameters = p.parseTable(apiSection, "URI Parameters")
	apiSection, res.ContentParameters = p.parseTable(apiSection, "Content Properties")
	parseChildren(res.ContentParameters)
	apiSection, res.ContentExample = p.parseExample(apiSection, "Examples", "Response")
	apiSection, res.ResParameters = p.parseTable(apiSection, "Entity Definition")
	parseChildren(res.ResParameters)
	apiSection, res.ResExample = p.parseExample(apiSection, "Examples", "")
	return res, nil
}

func (p *parser) parseUrl(n *html.Node, header string) (*html.Node, string, string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		p.logger.LogNode("ParseUrl data", c.FirstChild)
		if c.FirstChild != nil && c.FirstChild.Data == header {
			c = c.NextSibling
			p.logger.LogNode("Parse Url success:", c.FirstChild)
			res := ""
			p.findNode(doc,
				func(n *html.Node) CheckRes {
					if n.DataAtom == atom.Code {
						return CheckResApply
					}
					return CheckResContinue
				},
				func(n *html.Node) {
					res = n.FirstChild.Data
				})
			return c.NextSibling, res[0], res[1]
		}
	}
	return n, "", ""
}

func (p *parser) parseExample(n *html.Node, header, stopAt string) (*html.Node, string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if stopAt != "" && c.LastChild.Data == stopAt {
			return n, ""
		}
		if c.FirstChild != nil && c.FirstChild.Data == header {
			c = c.NextSibling
			p.logger.LogNode("Parse Example success:", c)
			res := ""
			p.findNode(doc,
				func(n *html.Node) CheckRes {
					if n.DataAtom == atom.Code {
						return CheckResApply
					}
					return CheckResContinue
				},
				func(n *html.Node) {
					res = n.FirstChild.Data
				})
			return c.NextSibling, res
		}
	}
	return n, ""
}

func (p *parser) parseTable(n *html.Node, header string) (*html.Node, []*ParameterDef) {
	res := []*ParameterDef{}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.FirstChild != nil && c.FirstChild.Data == header && c.NextSibling.DataAtom == atom.Table {
			c = c.NextSibling

			p.logger.Log("Parse Table success: %+v", c)
			for row := c.LastChild.FirstChild; row != nil; row = row.NextSibling {
				item := &ParameterDef{}
				cell := row.FirstChild
				item.Name = cell.FirstChild.Data
				cell = cell.NextSibling

				item.Type = cell.FirstChild.Data
				cell = cell.NextSibling

				item.Description = cell.FirstChild.Data
				cell = cell.NextSibling

				if cell != nil {
					item.IsRequired = cell.FirstChild.Data == "Yes"
				}
				res = append(res, item)
			}
			return c.NextSibling, res
		}
	}
	return n, nil
}

func (p *parser) hasAttr(attrs []html.Attribute, key, val string) bool {
	for _, attr := range attrs {
		attrKey := attr.Key
		attrVal := attr.Val
		if key == "" {
			attrKey = ""
		}
		if val == "" {
			attrVal = ""
		}
		if attrKey == key && attrVal == val {
			return true
		}
	}
	return false
}

func (p *parser) findNode(n *html.Node, checker func(*html.Node) CheckRes, action func(*html.Node)) {
	checkRes := checker(n)
	if checkRes == CheckResReturn {
		return
	}
	if checkRes == CheckResApply {
		action(n)
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			p.findNode(c, checker, action)
		}
	}
}
