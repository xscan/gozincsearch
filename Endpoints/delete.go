package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Delete struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewDelete() AbstractEndpoint {
	return Delete{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c Delete) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c Delete) GetIndex() string {
	return c.index
}
func (c Delete) GetURI() string {
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

	uri := fmt.Sprintf("/es/%s/%s/%s", index, docType, id)

	return uri
}
func (c Delete) GetMethod() string {
	c.method = "DELETE"
	return c.method
}

func (c Delete) GetBody() string {
	b, _ := json.Marshal(c.body)
	return string(b)
}
func (c Delete) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c Delete) GetParams() string {
	return ""
}
func (c Delete) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c Delete) GetOptions() string {

	return ""
}
func (c Delete) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c Delete) GetType() string {
	return c.docType
}
func (c Delete) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c Delete) GetID() string {
	return c.id

}
func (c Delete) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c Delete) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
