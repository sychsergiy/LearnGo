package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubClient struct {
	AuthCreds BasicAuthCreds
}

const GitHubRestApiUrl = "https://api.github.com"

const (
	GET  = "GET"
	POST = "POST"
)

func (ghc *GitHubClient) buildRequest(method, url string, creds BasicAuthCreds) (*http.Request, error) {
	// todo: add body
	req, err := http.NewRequest(method, GitHubRestApiUrl+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.username, creds.password)
	return req, nil
}

func (ghc *GitHubClient) sendRequest(method, url string, creds BasicAuthCreds) (*http.Response, error) {
	client := &http.Client{}
	request, err := ghc.buildRequest(method, url, creds)
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

func (ghc *GitHubClient) ListIssues() ([]issue, error) {
	resp, err := ghc.sendRequest(GET, "/issues", ghc.AuthCreds)
	if err != nil {
		return nil, err
	}
	issues := make([]issue, 0)
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return issues, nil
}
