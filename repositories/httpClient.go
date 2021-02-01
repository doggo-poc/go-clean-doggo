package repositories

import "net/http"

type HttpClient interface {
	Execute(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) HttpClient {
	return &httpClient{
		client: client,
	}
}

func (httpClient *httpClient) Execute(req *http.Request) (*http.Response, error) {
	return httpClient.client.Do(req)
}
