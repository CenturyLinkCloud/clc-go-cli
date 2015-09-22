package main

import (
	"encoding/json"
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

	if parseErr != nil {
		return nil, parseErr
	}
	err = p.postProcess(res)
	return res, err
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
				param.Children, err = p.parseTable(reqSec, resSec, false, strings.Title(param.Name+" Definition"))
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
	res.UrlParameters, err = p.parseTable(reqSec, resSec, true, "URI Parameters", "URI Properties", "URI and Querystring Parameters")
	if err != nil {
		return nil, err
	}
	res.ContentParameters, err = p.parseTable(reqSec, resSec, false, "Content Properties")
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
	res.ResParameters, err = p.parseTable(resSec, nil, false, "Entity Definition")
	if err != nil {
		return nil, err
	}
	err = parseChildren(res.ResParameters)
	if err != nil {
		return nil, err
	}
	res.ResExample, err = p.parseExample(resSec, nil, "Examples")
	return res, err
}

func (p *parser) parseUrl(startNode, endNode *html.Node, headerText string) (string, string, error) {
	p.logger.Log("parseUrl called")
	res, err := p.findNodeByHeader(startNode, endNode, atom.Pre, atom.Code, 1, headerText)
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
	return array[0], strings.TrimSpace(array[1]), nil
}

func (p *parser) parseExample(startNode, endNode *html.Node, headerText string) (interface{}, error) {
	p.logger.Log("parseExample called")
	res, err := p.findNodeByHeader(startNode, endNode, atom.Pre, atom.Code, 2, headerText)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	content := ""
	for n := res.FirstChild; n != nil; n = n.NextSibling {
		if n.Type == html.TextNode {
			content += n.Data
		}
	}
	if content == "" {
		return nil, nil
	}
	data := new(interface{})
	p.logger.Log("Converting object: %s", content)
	err = json.Unmarshal([]byte(content), data)
	if err != nil && (err.Error() == "invalid character '}' looking for beginning of object key string" || err.Error() == "invalid character ']' looking for beginning of value") {
		i := strings.LastIndex(content, ",")
		content = content[:i] + content[i+1:]
		err = json.Unmarshal([]byte(content), data)
	}
	if err != nil && err.Error() == "invalid character ',' after top-level value" {
		content = "[" + content + "]"
		err = json.Unmarshal([]byte(content), data)
	}
	if err != nil && err.Error() == "invalid character '\"' after object key:value pair" {
		content = strings.Replace(content, "/v2/groups/acct/31d13f501459411ba59304f3d47486eb/scheduledActivities\"", "/v2/groups/acct/31d13f501459411ba59304f3d47486eb/scheduledActivities\",", -1)
		err = json.Unmarshal([]byte(content), data)
	}
	if err != nil && err.Error() == "invalid character '(' looking for beginning of value" {
		content = strings.Replace(content, "(/v2/groups/ALIAS/GROUP/statistics?start=2014-04-09T20:00:00&sampleInterval=01:00:00)", "", -1)
		err = json.Unmarshal([]byte(content), data)
	}
	return *data, err
}

func (p *parser) parseTable(startNode, endNode *html.Node, capitalize bool, headerText ...string) ([]*ParameterDef, error) {
	p.logger.Log("parseTable called")
	table, err := p.findNodeByHeader(startNode, endNode, atom.Table, atom.Table, 1, headerText...)
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
		itemName := cell.FirstChild.Data
		if capitalize {
			itemName = strings.Title(itemName)
		}
		item.Name = itemName
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

func (p *parser) findNodeByHeader(startNode, endNode *html.Node, containerType, elemType atom.Atom, containerMaxRemoteness int, headerText ...string) (*html.Node, error) {
	header := p.findNextNode(startNode, endNode, headerText...)
	if header != nil {
		p.logger.LogNode("findNodeByHeader header found:", header)

		//if next node is paragraf we don't need to return error, because this is a valid case
		//just return nil
		_, err := p.findNextNodeByType(header, atom.P, 2)
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
	p.logger.Log("findNodeByHeader header not found %v", headerText)
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

func (p *parser) findNextNode(startNode, endNode *html.Node, text ...string) *html.Node {
	for c := startNode.NextSibling; c != endNode; c = c.NextSibling {
		for _, item := range text {
			if c.FirstChild != nil && strings.HasSuffix(c.FirstChild.Data, item) {
				return c
			}
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

func (p *parser) postProcess(api []*ApiDef) error {
	indexToDelete := make([]bool, len(api))
	for i := 0; i < len(api); i++ {
		for j := i + 1; j < len(api); j++ {
			if api[i].Url == api[j].Url && api[i].Method == api[j].Method {
				if api[i].ContentParameters == nil && api[j].ContentParameters != nil {
					indexToDelete[i] = true
				} else if api[j].ContentParameters == nil && api[i].ContentParameters != nil {
					indexToDelete[j] = true
				} else {
					indexToDelete[i] = true
					//return fmt.Errorf("Same API found, but can't decide what wersion to delete, api1: %#v, api2: %#v", api[i], api[j])
				}
			}
		}
	}

	j := 0
	for i := 0; i < len(api); i++ {
		if indexToDelete[i] {
			j--
		} else if i != j {
			api[j] = api[i]
		}
		j++
	}
	api = api[0:j]

	return nil
}
