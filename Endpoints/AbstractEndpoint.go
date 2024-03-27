package endpoints

type AbstractEndpoint interface {
	SetIndex(index string) AbstractEndpoint
	GetIndex() string
	GetURI() string
	GetBody() string
	SetBody(body map[string]interface{}) AbstractEndpoint
	GetParams() string
	SetParams(params map[string]string) AbstractEndpoint
	GetOptions() string
	SetOptions(options map[string]string) AbstractEndpoint
	SetType(docType string) AbstractEndpoint
	GetType() string
	GetID() string
	SetID(id string) AbstractEndpoint
	GetMethod() string
	String() string
}
