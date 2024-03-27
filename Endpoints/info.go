package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Info struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewInfo() AbstractEndpoint {
	return Info{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c Info) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c Info) GetIndex() string {
	return c.index
}
func (c Info) GetURI() string {

	uri := fmt.Sprintf("/es/")

	return uri
}
func (c Info) GetMethod() string {
	c.method = "GET"
	return c.method
}

func (c Info) GetBody() string {
	b, _ := json.Marshal(c.body)
	return string(b)
}
func (c Info) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c Info) GetParams() string {
	return ""
}
func (c Info) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c Info) GetOptions() string {

	return ""
}
func (c Info) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c Info) GetType() string {
	return c.docType
}
func (c Info) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c Info) GetID() string {
	return c.id

}
func (c Info) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c Info) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
