package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const restAPIURL = "https://api.github.com"

type SearchIssuesResult struct {
}

func SearchIssues() (*SearchIssuesResult, error) {
	listIssuesURL := restAPIURL + "/search/issues?q=test"
	resp, err := http.Get(listIssuesURL)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	s := string(body[:])
	println(s)
	result := SearchIssuesResult{}

	//if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	//	resp.Body.Close()
	//	return nil, err
	//}
	resp.Body.Close()
	return &result, nil
}
