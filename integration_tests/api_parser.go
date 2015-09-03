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
	p.findNode(doc,
		func(n *html.Node) CheckRes {
			if parseErr != nil {
				return CheckResReturn
			}
			if p.hasAttr(n.Attr, "class", "kb-post-content kb-post-content--api") {
				return CheckResApply
			}
			return CheckResContinue
		},
		func(n *html.Node) {
			api, err := p.parseApiNode(n)
			if err != nil {
				parseErr = err
			} else if api != nil {
				res = append(res, api)
			}
		})

	return res, parseErr
}

func (p *parser) parseApiNode(n *html.Node) (*ApiDef, error) {
	p.logger.LogNode("-----------Parsing node", n)
	reqSec := p.findNextNode(n.FirstChild, nil, "URL")
	if reqSec == nil {
		p.logger.LogNode("Not an API node", n)
		return nil, nil
	}
	resSec := p.findNextNode(reqSec, nil, "Response")
	if resSec == nil {
		return nil, fmt.Errorf("Response section not found")
	}

	parseChildren := func(parameters []*ParameterDef) error {
		for _, param := range parameters {
			if param.Type == "complex" {
				var err error
				param.Children, err = p.parseTable(reqSec, resSec, strings.Title(param.Name+" Definition"))
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	res := &ApiDef{}
	var err error
	res.Method, res.Url, err = p.parseUrl(reqSec, resSec, "Structure")
	if err != nil {
		return nil, err
	}
	_, res.UrlExample, err = p.parseUrl(reqSec, resSec, "Example")
	if err != nil {
		return nil, err
	}
	res.UrlParameters, err = p.parseTable(reqSec, resSec, "URI Parameters")
	if err != nil {
		return nil, err
	}
	res.ContentParameters, err = p.parseTable(reqSec, resSec, "Content Properties")
	if err != nil {
		return nil, err
	}
	err = parseChildren(res.ContentParameters)
	if err != nil {
		return nil, err
	}
	res.ContentExample, err = p.parseExample(reqSec, resSec, "Examples")
	if err != nil {
		return nil, err
	}
	res.ResParameters, err = p.parseTable(resSec, nil, "Entity Definition")
	if err != nil {
		return nil, err
	}
	err = parseChildren(res.ResParameters)
	if err != nil {
		return nil, err
	}
	res.ResExample, err = p.parseExample(resSec, nil, "Examples")
	return res, nil
}

func (p *parser) parseUrl(startNode, endNode *html.Node, headerText string) (string, string, error) {
	p.logger.Log("parseUrl called")
	res, err := p.findNodeByHeader(startNode, endNode, headerText, atom.Pre, atom.Code, 1)
	if err != nil {
		return "", "", err
	}
	if res == nil {
		return "", "", nil
	}
	array := strings.Split(res.FirstChild.Data, " ")
	if len(array) != 2 {
		return "", "", fmt.Errorf("Incorrect format of url section")
	}
	return array[0], array[1], nil
}

func (p *parser) parseExample(startNode, endNode *html.Node, headerText string) (string, error) {
	p.logger.Log("parseExample called")
	res, err := p.findNodeByHeader(startNode, endNode, headerText, atom.Pre, atom.Code, 2)
	if err != nil {
		return "", err
	}
	if res == nil {
		return "", nil
	}
	return res.FirstChild.Data, nil
}

func (p *parser) parseTable(startNode, endNode *html.Node, headerText string) ([]*ParameterDef, error) {
	p.logger.Log("parseTable called")
	table, err := p.findNodeByHeader(startNode, endNode, headerText, atom.Table, atom.Table, 1)
	if err != nil {
		return nil, err
	}
	if table == nil {
		return nil, nil
	}

	res := []*ParameterDef{}

	tbody, err := p.findNextNodeByType(table.FirstChild, atom.Tbody, 2)
	if err != nil {
		return nil, err
	}
	p.logger.LogNode("parseTable tbody found", tbody)
	next := func(n *html.Node) *html.Node {
		for n = n.NextSibling; n != nil && n.DataAtom != atom.Td; n = n.NextSibling {
		}
		p.logger.LogNode("parseTable cell found", n)
		return n
	}
	for row := tbody.FirstChild; row != nil; row = row.NextSibling {
		if row.DataAtom != atom.Tr {
			continue
		}

		p.logger.LogNode("parseTable row found", row)
		item := &ParameterDef{}
		cell := next(row.FirstChild)
		item.Name = cell.FirstChild.Data
		cell = next(cell)

		item.Type = cell.FirstChild.Data
		cell = next(cell)

		item.Description = cell.FirstChild.Data
		cell = next(cell)

		if cell != nil {
			item.IsRequired = cell.FirstChild.Data == "Yes"
		}
		res = append(res, item)
	}
	if len(res) == 0 {
		return nil, fmt.Errorf("Empty table parsed")
	}
	return res, nil
}

func (p *parser) findNodeByHeader(startNode, endNode *html.Node, headerText string, containerType, elemType atom.Atom, containerMaxRemoteness int) (*html.Node, error) {
	header := p.findNextNode(startNode, endNode, headerText)
	if header != nil {
		p.logger.LogNode("findNodeByHeader header found:", header)

		//if next node is paragraf we don't need to return error, because this is a valid case
		//just return nil
		_, err := p.findNextNodeByType(header, atom.P, 1)
		if err == nil {
			return nil, nil
		}
		container, err := p.findNextNodeByType(header, containerType, containerMaxRemoteness)
		if err != nil {
			return nil, err
		}
		p.logger.LogNode("findNodeByHeader container:", container)
		var res *html.Node
		p.findNode(container,
			func(n *html.Node) CheckRes {
				if n.DataAtom == elemType {
					return CheckResApply
				}
				return CheckResContinue
			},
			func(n *html.Node) {
				p.logger.LogNode("findNodeByHeader target found:", n.FirstChild)
				res = n
			})
		if res == nil {
			return nil, fmt.Errorf("Node %v not found", elemType)
		}
		return res, nil
	}
	p.logger.Log("findNodeByHeader header not found %s", headerText)
	return nil, nil
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

func (p *parser) findNextNode(startNode, endNode *html.Node, text string) *html.Node {
	for c := startNode.NextSibling; c != endNode; c = c.NextSibling {
		if c.FirstChild != nil && c.FirstChild.Data == text {
			return c
		}
	}
	return nil
}

func (p *parser) findNextNodeByType(n *html.Node, nodeType atom.Atom, maxNodeRemotenes int) (*html.Node, error) {
	var c *html.Node
	i := 0
	for c = n.NextSibling; c != nil && c.DataAtom != nodeType && i < maxNodeRemotenes; c = c.NextSibling {
		if c.Type == html.ElementNode {
			i++
		}
	}
	if i == maxNodeRemotenes {
		return nil, fmt.Errorf("Next node with type %v and maxRemoteness %d not found", nodeType, maxNodeRemotenes)
	}
	if c == nil {
		return nil, fmt.Errorf("Next node with type %v not found", nodeType)
	}
	return c, nil
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
