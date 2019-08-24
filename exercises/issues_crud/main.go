package main

import (
	"encoding/json"
	"log"
	"os"
	"overview/exercises/issues_crud/github"
)

func getBasicAuthCredsFromEnv() github.BasicAuthCreds {
	username := os.Getenv("GIT_HUB_USERNAME")
	password := os.Getenv("GIT_HUB_PASSWORD")
	return github.BasicAuthCreds{Username: username, Password: password}
}

func listIssues(githubClient github.Client) {
	issues, err := githubClient.ListIssues()

	if issues != nil {
		for _, issue := range issues {
			marshaled, _ := json.MarshalIndent(issue, "", "\t")
			log.Println(string(marshaled))
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}

func deleteIssue(githubClient github.Client) {
	err := githubClient.LockIssue("LearnGo", 1, "any reason")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	creds := getBasicAuthCredsFromEnv()
	githubClient := github.Client{AuthCreds: creds}
	deleteIssue(githubClient)
}
