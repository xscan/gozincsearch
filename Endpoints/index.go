package endpoints

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Index struct {
	id      string
	index   string
	body    map[string]interface{}
	params  map[string]string
	options map[string]string
	method  string
	docType string
}

func NewIndex() AbstractEndpoint {
	return Index{
		body:    make(map[string]interface{}),
		options: make(map[string]string),
		params:  make(map[string]string),
	}
}

func (c Index) SetIndex(index string) AbstractEndpoint {
	c.index = index
	return c
}
func (c Index) GetIndex() string {
	return c.index
}
func (c Index) GetURI() string {
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

	uri := fmt.Sprintf("/api/%s/%s", index, docType)

	if c.id != "" {
		uri = fmt.Sprintf("/api/%s/%s/%s", index, docType, id)
	}

	return uri
}
func (c Index) GetMethod() string {
	// c.method = "GET"
	if c.id != "" {
		return "PUT"
	} else {
		return "POST"
	}

	// return c.method
}

func (c Index) GetBody() string {
	b, err := json.Marshal(c.body)
	if err != nil {
		return err.Error()
	}
	return string(b)

}
func (c Index) SetBody(body map[string]interface{}) AbstractEndpoint {
	c.body = body
	return c
}

func (c Index) GetParams() string {
	return ""
}
func (c Index) SetParams(params map[string]string) AbstractEndpoint {
	c.params = params
	return c
}
func (c Index) GetOptions() string {

	return ""
}
func (c Index) SetOptions(options map[string]string) AbstractEndpoint {

	return c
}

func (c Index) GetType() string {
	return c.docType
}
func (c Index) SetType(docType string) AbstractEndpoint {
	c.docType = docType
	return c
}

func (c Index) GetID() string {
	return c.id

}
func (c Index) SetID(id string) AbstractEndpoint {
	c.id = id
	return c
}
func (c Index) String() string {
	msg := strings.Builder{}
	msg.WriteString("url:" + c.GetURI())

	return msg.String()
}
