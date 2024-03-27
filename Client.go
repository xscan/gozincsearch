package gozincsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	endpoints "github.com/xscan/gozincsearch/Endpoints"
)

type Client struct {
	ClientBuild *ClientBuild
}

func (c *Client) Create() {

}

func (cli *Client) Index(params map[string]string) string {
	id := cli.ExtractArgument(params, "id")
	index := cli.ExtractArgument(params, "index")
	docType := cli.ExtractArgument(params, "type")
	body := cli.ExtractArgument(params, "body")

	var bodyMap map[string]interface{}

	json.Unmarshal([]byte(body), &bodyMap)

	// fmt.Println(mapjson)

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["Index"]

	endpoint = endpoint.SetID(id).
		SetIndex(index).
		SetType(docType).
		SetParams(params).
		SetBody(bodyMap)
	fmt.Println(endpoint)

	return cli.performRequest(endpoint)
}

func (cli *Client) ExtractArgument(params map[string]string, arg string) string {

	if _, ok := params[arg]; ok {
		val := params[arg]
		delete(params, arg)
		return val
	} else {
		return ""
	}

}
func (cli *Client) Get(params map[string]string) string {
	id := cli.ExtractArgument(params, "id")
	index := cli.ExtractArgument(params, "index")
	docType := cli.ExtractArgument(params, "type")

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["Get"]
	endpoint = endpoint.SetID(id).
		SetIndex(index).
		SetType(docType).
		SetParams(params)

	return cli.performRequest(endpoint)
}

func (cli *Client) performRequest(endpoint endpoints.AbstractEndpoint) string {

	uri := fmt.Sprintf("http://%s%s", cli.ClientBuild.host, endpoint.GetURI())
	// fmt.Println(uri)
	req, err := http.NewRequest(endpoint.GetMethod(), uri, nil)
	if endpoint.GetBody() != "" {
		req, err = http.NewRequest(endpoint.GetMethod(), uri, bytes.NewBuffer([]byte(endpoint.GetBody())))
	}

	if err != nil {
		log.Println(err)
	}
	// Basic YWRtaW46Q29tcGxleHBhc3MjMTIz
	req.Header = map[string][]string{
		"Content-Type":  {"application/json"},
		"Accept":        {"application/json"},
		"Authorization": {"Basic " + cli.ClientBuild.apiKey},
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)

	return string(bodyBytes)
}

func (cli *Client) Search(params map[string]string) string {
	id := cli.ExtractArgument(params, "id")
	index := cli.ExtractArgument(params, "index")
	docType := cli.ExtractArgument(params, "type")
	body := cli.ExtractArgument(params, "body")

	var bodyMap map[string]interface{}

	json.Unmarshal([]byte(body), &bodyMap)

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["Search"]
	endpoint = endpoint.SetID(id).
		SetIndex(index).
		SetType(docType).
		SetBody(bodyMap).
		SetParams(params)

	return cli.performRequest(endpoint)
}

func (cli *Client) MSearch(params map[string]string) string {
	id := cli.ExtractArgument(params, "id")
	index := cli.ExtractArgument(params, "index")
	docType := cli.ExtractArgument(params, "type")
	body := cli.ExtractArgument(params, "body")

	var bodyMap map[string]interface{}

	json.Unmarshal([]byte(body), &bodyMap)

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["MSearch"]
	endpoint = endpoint.SetID(id).
		SetIndex(index).
		SetType(docType).
		SetBody(bodyMap).
		SetParams(params)

	return cli.performRequest(endpoint)
}

func (cli *Client) Delete(params map[string]string) string {
	id := cli.ExtractArgument(params, "id")
	index := cli.ExtractArgument(params, "index")
	docType := cli.ExtractArgument(params, "type")

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["Delete"]
	endpoint = endpoint.SetID(id).
		SetIndex(index).
		SetType(docType).
		SetParams(params)

	return cli.performRequest(endpoint)
}
func (cli *Client) Info(params map[string]string) string {

	var endpoint endpoints.AbstractEndpoint
	endpoint = cli.ClientBuild.endpoints["Info"]
	endpoint = endpoint.
		SetParams(params)

	return cli.performRequest(endpoint)
}
func (c *Client) Bluk() {

}
