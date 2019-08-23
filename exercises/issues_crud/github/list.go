package github

import (
	"encoding/json"
	"time"
)

type label struct {
	Id          int
	NodeId      string
	URL         string
	Name        string
	Description string
	Color       string
	Default     bool
}
type user struct {
	Login string
	Id    int
	URL   string
	Type  string
}

type Issue struct {
	Id        int
	Number    int
	URL       string
	State     string
	Title     string
	Body      string
	User      user
	Labels    []label
	Assignee  user
	Assignees []user
	Locked    bool
	ClosedAt,
	CreatedAt,
	UpdatedAt time.Time
}

func (c *Client) ListIssues() ([]Issue, error) {
	resp, err := c.sendRequest(GET, "/issues", c.AuthCreds)
	if err != nil {
		return nil, err
	}
	issues := make([]Issue, 0)
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return issues, nil
}
