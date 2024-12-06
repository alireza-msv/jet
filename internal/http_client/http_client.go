package httpclient

import (
	"io"
	"net/http"
	"net/url"

	"github.com/alireza-msv/jet/internal/utils"
)

// A simple wrapper for http.Client
type HttpClient struct {
	client  *http.Client
	baseURL string
}

func NewHttpClient(baseURL string) *HttpClient {
	c := &http.Client{}

	return &HttpClient{
		client:  c,
		baseURL: baseURL,
	}
}

// Calls http client Post method with application/json contentType
func (hc *HttpClient) PostJSON(path string, body io.Reader) (*http.Response, error) {
	reqURL, err := url.JoinPath(hc.baseURL, path)
	if err != nil {
		return nil, err
	}

	return hc.client.Post(reqURL, utils.HttpContentTypeJSONHeader, body)
}
