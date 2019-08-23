package main

import (
	"io/ioutil"
	"net/http"
)

func ListIssues(creds AuthCredentials) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", restAPIURL+"/issues", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(creds.username, creds.password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	s := string(body[:])
	println(s)
	return nil
}
