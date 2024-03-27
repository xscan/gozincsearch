# gozincsearch
zincsearch的golang的封装库

#  暂时支持的接口
- [x] index
- [x] delete
- [x] get
- [x] info
- [x] search
- [x] msearch

# 安装
```
go get -u github.com/xscan/gozincsearch
```

# 基本使用例子
```golang
package main

import (
	"fmt"

	"github.com/xscan/gozincsearch"
)

var client *gozincsearch.Client

func init() {
	clientBuild := &gozincsearch.ClientBuild{}
	client = clientBuild.
		SetApiKey("YWRtaW46Q29tcGxleHBhc3MjMTIz").
		SetHost("localhost:4080").
		Build()
}
func main() {

	info := client.Info(nil)
	fmt.Println(info)
}

```
# 全部API使用例子

```golang
package gozincsearch_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/xscan/gozincsearch"
)

var client *gozincsearch.Client

func init() {
	clientBuild := &gozincsearch.ClientBuild{}
	client = clientBuild.
		SetApiKey("YWRtaW46Q29tcGxleHBhc3MjMTIz").
		SetHost("localhost:4080").
		Build()
}

func TestCreate(t *testing.T) {

	data1 := map[string]string{
		"id":    "testid",
		"type":  "_doc",
		"body":  `{"a":"12"}`,
		"index": "olympics",
	}
	res := client.Index(data1)
	fmt.Println(res)

}

func TestUpdate(t *testing.T) {

	data1 := map[string]string{
		"id":    "testid",
		"type":  "_doc",
		"body":  `{"a":"123333333"}`,
		"index": "olympics",
	}
	res := client.Index(data1)
	fmt.Println(res)

}

func TestGet(t *testing.T) {
	data := map[string]string{
		"id":    "testid",
		"type":  "_doc",
		"index": "olympics",
	}
	res := client.Get(data)

	fmt.Println(res)
}

func TestDelete(t *testing.T) {

	data3 := map[string]string{
		"id":    "testid",
		"type":  "_doc",
		"index": "olympics",
	}
	res3 := client.Delete(data3)
	fmt.Println(res3)
}

func TestSearch(t *testing.T) {

	// 搜索
	//
	search := `{
			"query": {
				"bool": {
					"must": [
						{ "match": { "City":  "paris" }},
						{ "match": { "Medal": "gold" }}
					],
					"should": [],
					"must_not": [],
					"filter": [
						{ "term":  { "Country": "ger" }},
						{ "range": { "@timestamp": { "gte": "2015-01-01", "format": "2006-01-02" }}}
					]
				}
			},
			"size": 10
		}
	
	`

	data2 := map[string]string{
		"type":  "_doc",
		"body":  search,
		"index": "olympics",
	}

	res2 := client.Search(data2)
	fmt.Println(res2)
}

func TestMSearch(t *testing.T) {

	msearch := `{{"index": "olympics"}
	{ "query": { "bool": { "must": [ { "match": { "City":  "paris"        }}, { "match": { "Medal": "gold" }} ], "filter": [ { "term":  { "Country": "ger" }}, { "range": { "@timestamp": { "gte": "2015-01-01", "format": "2006-01-02" }}} ] } } }
	{"index": "olympics"}
	{"query" : {"match_all" : {}}}}`

	msearchparams := map[string]interface{}{
		"mquery": msearch,
	}

	msearchoaramsstr, err := json.Marshal(msearchparams)

	if err != nil {
		panic(err)
	}

	data2 := map[string]string{
		"type":  "_doc",
		"body":  string(msearchoaramsstr),
		"index": "olympics",
	}
	res2 := client.MSearch(data2)
	fmt.Println(res2)

}
func TestInfo(t *testing.T) {

	data5 := map[string]string{}

	res5 := client.Info(data5)
	fmt.Println(res5)

}

```