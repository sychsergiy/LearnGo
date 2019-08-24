package github

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (c *Client) UnlockIssue(repo string, issueNumber int) error {
	owner := c.AuthCreds.Username
	url := strings.Join([]string{"repos", owner, repo, "issues", strconv.Itoa(issueNumber), "lock"}, "/")

	resp, reqErr := c.sendRequest(DELETE, url, c.AuthCreds, nil)
	if reqErr != nil {
		return reqErr
	}

	if resp.StatusCode != http.StatusNoContent {
		resp.Body.Close()
		return fmt.Errorf("HTTP request failed: %s", resp.Status)
	}
	return nil
}
