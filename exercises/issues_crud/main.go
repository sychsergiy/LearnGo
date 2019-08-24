package main

import (
	"encoding/json"
	"log"
	"os"
	"overview/exercises/issues_crud/cli"
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

func lockIssue(githubClient github.Client) {
	err := githubClient.LockIssue("LearnGo", 1, "any reason")
	if err != nil {
		log.Fatal(err)
	}
}

func unlockIssue(githubClient github.Client) {
	err := githubClient.UnlockIssue("LearnGo", 1)
	if err != nil {
		log.Fatal(err)
	}
}

func createIssue(githubClient github.Client) {
	data := github.CreateIssueRequestData{Title: "Test title", Body: "Test Body"}
	issue, err := githubClient.CreateIssue("LearnGo", data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}

func editIssue(githubClient github.Client) {
	data := github.EditIssueRequestData{Title: "New title", Body: "edited Body", State: "closed"}
	issue, err := githubClient.EditIssue("LearnGo", 3, data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}

func main() {
	LIST := "list"
	client := github.Client{AuthCreds: getBasicAuthCredsFromEnv()}
	CLI := cli.New()
	CLI.RegisterHandler(LIST, ListIssues{GithubClient: client})

	CLI.Run()
}

//func main() {
//	creds := getBasicAuthCredsFromEnv()
//	githubClient := github.Client{AuthCreds: creds}
//lockIssue(githubClient)
//unlockIssue(githubClient)
//createIssue(githubClient)
//editIssue(githubClient)
//}
