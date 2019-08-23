package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type issue struct {
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

func ListIssues(creds BasicAuthCreds) (*[]issue, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", restAPIURL+"/issues", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(creds.username, creds.password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	issues := make([]issue, 0)
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		resp.Body.Close()
		return nil, err
	}
	return &issues, nil
}
