package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type EditIssueRequestData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	State string `json:"state"`
}

func (c *Client) EditIssue(repo string, issueNumber int, data EditIssueRequestData) (*Issue, error) {
	url := strings.Join([]string{"repos", c.AuthCreds.Username, repo, "issues", strconv.Itoa(issueNumber)}, "/")
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := c.sendRequest(PATH, url, c.AuthCreds, body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fail to update issue, status code: %d", resp.StatusCode)
	}

	issue := &Issue{}
	if err := json.NewDecoder(resp.Body).Decode(issue); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return issue, nil
}
