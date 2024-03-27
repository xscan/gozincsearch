package endpoints

import (
	"fmt"
	"strings"
)

type MSearch struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewMSearch() AbstractEndpoint {
	return MSearch{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c MSearch) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c MSearch) GetIndex() string {
	return c.index
}
func (c MSearch) GetURI() string {

	// index := c.GetIndex()

	uri := fmt.Sprintf("/es/_msearch")

	return uri
}
func (c MSearch) GetMethod() string {

	c.method = "GET"
	if c.body != nil {
		c.method = "POST"
	}
	return c.method
}

func (c MSearch) GetBody() string {
	if _, ok := c.body["mquery"]; ok {
		return c.body["mquery"].(string)
	}
	return ""
}
func (c MSearch) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c MSearch) GetParams() string {
	return ""
}
func (c MSearch) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c MSearch) GetOptions() string {

	return ""
}
func (c MSearch) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c MSearch) GetType() string {
	return c.docType
}
func (c MSearch) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c MSearch) GetID() string {
	return c.id

}
func (c MSearch) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c MSearch) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
