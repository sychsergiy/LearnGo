package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type DeleteRequestBody struct {
	Locked           bool   `json:"locked"`
	ActiveLockReason string `json:"active_lock_reason"`
}

func (c *Client) LockIssue(repo string, issueNumber int, reason string) error {
	owner := c.AuthCreds.Username
	url := strings.Join([]string{"repos", owner, repo, "issues", strconv.Itoa(issueNumber), "lock"}, "/")

	body, err := json.Marshal(DeleteRequestBody{true, reason})
	if err != nil {
		return err
	}

	resp, reqErr := c.sendRequest(PUT, url, c.AuthCreds, body)
	if reqErr != nil {
		return reqErr
	}

	if resp.StatusCode != http.StatusNoContent {
		resp.Body.Close()
		return fmt.Errorf("HTTP request failed: %s", resp.Status)
	}
	return nil
}
