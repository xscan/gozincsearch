package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Get struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewGet() AbstractEndpoint {
	return Get{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c Get) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c Get) GetIndex() string {
	return c.index
}
func (c Get) GetURI() string {
	if len(c.id) == 0 {
		panic("id is required for Get")
	}
	if len(c.index) == 0 {
		panic("index is required for Get")
	}
	if len(c.docType) == 0 {
		panic("type is required for Get")
	}
	id := c.GetID()
	index := c.GetIndex()
	docType := c.GetType()

	// uri = "/index/type/id"

	uri := fmt.Sprintf("/api/%s/%s/%s", index, docType, id)

	return uri
}
func (c Get) GetMethod() string {
	c.method = "GET"
	return c.method
}

func (c Get) GetBody() string {
	b, _ := json.Marshal(c.body)
	return string(b)
}
func (c Get) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c Get) GetParams() string {
	return ""
}
func (c Get) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c Get) GetOptions() string {

	return ""
}
func (c Get) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c Get) GetType() string {
	return c.docType
}
func (c Get) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c Get) GetID() string {
	return c.id

}
func (c Get) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c Get) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
