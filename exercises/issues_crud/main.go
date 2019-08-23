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

func main() {
	creds := getBasicAuthCredsFromEnv()
	githubClient := github.Client{AuthCreds: creds}

	issues, err := githubClient.ListIssues()

	if issues != nil {
		for _, issue := range issues {
			json, _ := json.MarshalIndent(issue, "", "\t")
			log.Println(string(json))
		}
	}
	if err != nil {
		log.Fatal(err)
	}

}
