package commands

import (
	"encoding/json"
	"log"
	"overview/exercises/issues_crud/cli"
	"overview/exercises/issues_crud/github"
)

type CreateIssue struct {
	GithubClient github.Client
}

func (handler *CreateIssue) Execute() {
	args := cli.RetrieveArgs(2)
	repo, title := args[0], args[1]
	body := cli.RetrieveArgFromEditor()
	data := github.CreateIssueRequestData{Title: title, Body: body}
	issue, err := handler.GithubClient.CreateIssue(repo, data)
	if err != nil {
		log.Fatal(err)
	}
	if issue != nil {
		marshaled, _ := json.MarshalIndent(issue, "", "\t")
		log.Println(string(marshaled))
	}
}
