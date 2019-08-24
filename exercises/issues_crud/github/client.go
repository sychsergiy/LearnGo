package github

import (
	"bytes"
	"net/http"
)

type Client struct {
	AuthCreds BasicAuthCreds
}

const RestApiUrl = "https://api.github.com/"

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	PATH   = "PATCH"
)

func (c *Client) buildRequest(method, url string, creds BasicAuthCreds, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, RestApiUrl+url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Username, creds.Password)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Client) sendRequest(method, url string, creds BasicAuthCreds, body []byte) (*http.Response, error) {
	client := &http.Client{}
	request, err := c.buildRequest(method, url, creds, body)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
