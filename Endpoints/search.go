package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Search struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewSearch() AbstractEndpoint {
	return Search{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c Search) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c Search) GetIndex() string {
	return c.index
}
func (c Search) GetURI() string {

	index := c.GetIndex()

	uri := fmt.Sprintf("/es/%s/_search", index)

	return uri
}
func (c Search) GetMethod() string {

	c.method = "GET"
	if c.body != nil {
		c.method = "POST"
	}
	return c.method
}

func (c Search) GetBody() string {
	b, _ := json.Marshal(c.body)
	return string(b)
}
func (c Search) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c Search) GetParams() string {
	return ""
}
func (c Search) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c Search) GetOptions() string {

	return ""
}
func (c Search) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c Search) GetType() string {
	return c.docType
}
func (c Search) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c Search) GetID() string {
	return c.id

}
func (c Search) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c Search) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
