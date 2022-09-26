package clients

import "net/http"

type clientStruct struct{}

type ClientInterface interface {
	Get(*http.Request) (*http.Response, error)
}

var (
	ClientStruct ClientInterface = &clientStruct{}
)

func (ci *clientStruct) Get(r *http.Request) (*http.Response, error) {
	client := http.Client{}
	return client.Do(r)
}
