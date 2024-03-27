package gozincsearch

import (
	"fmt"

	endpoints "github.com/xscan/gozincsearch/Endpoints"
)

type ClientBuild struct {
	host      string
	apiKey    string
	endpoints map[string]endpoints.AbstractEndpoint
}

func (c *ClientBuild) SetHost(host string) *ClientBuild {
	c.host = host
	return c
}
func (c *ClientBuild) SetApiKey(apiKey string) *ClientBuild {
	c.apiKey = apiKey
	return c
}

func (c *ClientBuild) Build() *Client {
	// 注册endpoints
	endpoints := map[string]endpoints.AbstractEndpoint{
		"Get":     endpoints.NewGet(),
		"Index":   endpoints.NewIndex(),
		"Search":  endpoints.NewSearch(),
		"Delete":  endpoints.NewDelete(),
		"MSearch": endpoints.NewMSearch(),
		"Info":    endpoints.NewInfo(),
	}
	c.endpoints = endpoints

	return &Client{ClientBuild: c}
}

func (c *ClientBuild) String() string {
	return fmt.Sprintf("host:%s apikey:%s", c.host, c.apiKey)
}
