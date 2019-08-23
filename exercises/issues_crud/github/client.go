package github

import (
	"fmt"
	"net/http"
)

type Client struct {
	AuthCreds BasicAuthCreds
}

const RestApiUrl = "https://api.github.com"

const (
	GET  = "GET"
	POST = "POST"
)

func (c *Client) buildRequest(method, url string, creds BasicAuthCreds) (*http.Request, error) {
	// todo: add body
	req, err := http.NewRequest(method, RestApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.Username, creds.Password)
	return req, nil
}

func (c *Client) sendRequest(method, url string, creds BasicAuthCreds) (*http.Response, error) {
	client := &http.Client{}
	request, err := c.buildRequest(method, url, creds)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("HTTP request failed: %s", resp.Status)
	}
	return resp, nil
}
