package github

import (
	"encoding/json"
	"strings"
)

type CreateIssueRequestData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (c *Client) CreateIssue(repo string, data CreateIssueRequestData) (*Issue, error) {
	url := strings.Join([]string{"repos", c.AuthCreds.Username, repo, "issues"}, "/")
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(POST, url, c.AuthCreds, body)
	if err != nil {
		return nil, err
	}

	createdIssue := &Issue{}
	if err := json.NewDecoder(resp.Body).Decode(createdIssue); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return createdIssue, nil
}
