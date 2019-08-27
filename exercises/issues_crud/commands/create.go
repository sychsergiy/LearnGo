package commands

import (
	"encoding/json"
	"log"
	"overview/exercises/issues_crud/github"
)

type CreateIssue struct {
	GithubClient github.Client
}

func (handler *CreateIssue) Execute() {
	data := github.CreateIssueRequestData{Title: "Test title", Body: "Test Body"}
	issue, err := handler.GithubClient.CreateIssue("LearnGo", data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
